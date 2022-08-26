package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/services"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
)

var StudentSheetColumns []string = []string{"Name", "Email", "Phone", "Year", "Batch", "Section", "Enrollment Number", "Date of Birth", "Fathers Name", "Mothers Name", "Guardian Email", "Guardian Phone"}
var TeacherSheetColumns []string = []string{"Name", "Official Email", "Alternate Email", "Department", "Course", "Designation"}

const studentsSheet string = "Students"
const teachersSheet string = "Teachers"

func GetOnboardingSheet(c *gin.Context) {
	var body types.GetOnboardingSheetBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// insert college in db
	college := models.CollegeModel{
		Name:             body.College.Name,
		Email:            body.College.Email,
		Phone:            body.College.Phone,
		WebsiteUrl:       body.College.Website,
		UniversityName:   body.College.UniversityName,
		City:             body.College.City,
		State:            body.College.State,
		CollegeType:      body.College.CollegeType,
		OnboardingStatus: "ON_GOING",
		IsActive:         true,
	}

	cid, err := models.CreateCollege(college)

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// create new xlsx file
	f := excelize.NewFile()

	// create sheets for all the courses
	for _, course := range body.Courses {
		f.NewSheet(course)
	}
	// create 1 sheet for teachers
	f.NewSheet(teachersSheet)

	sheetList := f.GetSheetList()

	// set students sheet headers
	for _, sheet := range sheetList {
		for index, col := range StudentSheetColumns {
			f.SetCellValue(sheet, fmt.Sprint("A%d", index+1), col)
		}
	}

	for index, col := range TeacherSheetColumns {
		f.SetCellValue(teachersSheet, fmt.Sprint("A%d", index+1), col)
	}

	filename := "COLLEGE_ONBOARDING_" + strconv.Itoa(cid) + ".xlsx"
	fileKey := "tmp/" + filename
	location := "edrank/" + filename

	// Save the spreadsheet by the given path.
	if err := f.SaveAs(fileKey); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	fileUrl, err := services.UploadToS3(c, "dir", fileKey, filename)

	file_registry := models.FileRegistryModel{
		Cid:      cid,
		Location: location,
		Url:      fileUrl,
		Type:     "COLLEGE_ONBOARDING",
	}

	frId, err := models.CreateFileRegistry(file_registry)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, errors.New("File registry creation failed : "+err.Error()))
		return
	}

	// remove the downloaded file
	os.Remove(fileKey)

	utils.SendResponse(c, "College created.", map[string]any{
		"college_id":        cid,
		"onboarding_status": "ON_GOING",
		"file_registry_id":  frId,
	})
}

// get file
// store it locally
// validate
// manipulate
// insert in db
// update onboarding status
// delete the file
func OnBoardCollegeController(c *gin.Context) {
	var body types.OnBoardCollegeBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	fr, err := models.GetFileRegistryByField("id", body.FileRegistryId)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	err = services.DownloadFromS3(c, fr.Location)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	f, err := excelize.OpenFile("tmp/" + fr.Location)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	sheetList := f.GetSheetList()

	if len(sheetList) != 2 {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Sheets Count. 2 Sheets required"))
		return
	}

	// var studentSheetValues map[string]string = map[string]string{
	// 	"Name":              "name",
	// 	"Email":             "email",
	// 	"Phone":             "phone",
	// 	"Year":              "year",
	// 	"Batch":             "batch",
	// 	"Section":           "section",
	// 	"Enrollment Number": "enrollment_number",
	// 	"Date of Birth":     "dob",
	// 	"Fathers Name":      "fathers_name",
	// 	"Mothers Name":      "mothers_name",
	// 	"Guardian Email":    "guardian_email",
	// 	"Guardian Phone":    "guardian_phone",
	// }

	// var teacherSheetValues map[string]string = map[string]string{
	// 	"Name":            "name",
	// 	"Official Email":  "email",
	// 	"Alternate Email": "alt_email",
	// 	"Department":      "department",
	// 	"Course":          "course_id",
	// 	"Designation":     "designation",
	// }

	for _, course := range body.Courses {
		studentRows, err := f.GetRows(course)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Column Sequence"))
			return
		}
		students := []models.StudentModel{}
		for index, row := range studentRows {
			if !utils.AreEqualArray(row, StudentSheetColumns) {
				utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Column Sequence"))
				return
			}
			if index != 0 {
				student := models.StudentModel{}
				students = append(students, student)
			}
		}
		models.CreateBulkStudents(students)
	}

	// create parents accounts
	// create teachers accounts

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// remove the downloaded file
	os.Remove("tmp/" + fr.Location)
}

func CreateNewCollgeAdminController(c *gin.Context) {
	var body types.CreateNewCollegeAdminBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// check if college admin with similar email exists
	ca, err := models.GetCollegeAdminByField("email", body.Email)

	if err != nil && err.Error() != "Cannot find college admin" {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}
	if ca.Email == body.Email {
		utils.SendResponse(c, "College Admin Already Exists with this email", map[string]any{})
		return
	}

	college_id := c.GetInt("CollegeId")
	name := body.Name

	if len(strings.Split(name, " ")) > 0 {
		name = strings.Split(body.Name, " ")[0]
	}
	password := "admin@" + name
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	college_admin := models.CollegeAdminModel{
		Cid:      college_id,
		Name:     body.Name,
		Email:    body.Email,
		IsActive: true,
		Password: string(bytes),
	}

	caId, err := models.CreateCollegeAdmin(college_admin)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// TODO: send mail to new college admin

	utils.SendResponse(c, "College Admin Created!", map[string]any{
		"college_admin_id": caId,
	})
}

// paginated
func GetEntitiesOfMyCollegeController(c *gin.Context) {
	cid, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	entity := c.Param("entity")

	if entity == "" || (utils.Find([]string{"teachers", "parents", "students", "college_admins"}, entity) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Entity"))
		return
	}

	size, exists := c.GetQuery("size")

	if !exists {
		size = "10"
	}

	page, exists := c.GetQuery("page")

	if !exists {
		page = "1"
	}

	// pagination opts
	limit, offset, err := utils.GetPaginationOpts(size, page)

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	switch entity {
	case "students":
		students, err := models.GetAllStudentsOfMyCollege(cid.(int), limit, offset)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Students of your College!", map[string]any{
			"students": students,
		})
	case "parents":
		parents, err := models.GetParentsOfCollege(cid.(int), limit, offset)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Parents of your College!", map[string]any{
			"parents": parents,
		})

	case "teachers":
		teachers, err := models.GetAllTeachersOfMyCollege(cid.(int), limit, offset)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Teachers of your College!", map[string]any{
			"teachers": teachers,
		})
	case "college_admins":
		college_admins, err := models.GetAllCollegeAdminsOfCollege(cid.(int), limit, offset)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "College Admins of your College!", map[string]any{
			"college_admins": college_admins,
		})
	}
}

func ToggleFeedbackDriveController(c *gin.Context) {
	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	var body struct {
		Type    string `json:"drive_type"`
		Toggle  string `json:"toggle"`
		DriveId int    `json:"drive_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	if body.Type == "" || (utils.Find([]string{"college", "teacher"}, body.Type) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Drive Type"))
		return
	}

	if body.Toggle == "" || (utils.Find([]string{"enable", "disable"}, body.Toggle) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Toggle Type"))
		return
	}

	drives, err := models.GetCollegeFeedbackDrives(college_id.(int), body.Type)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	switch body.Toggle {
	case "enable":

		if len(drives) == 0 {
			fb_drive := models.FeedbackDrivesModel{
				CollegeId: college_id.(int),
				Type:      body.Type,
				IsActive:  true,
			}

			drive_id, err := models.CreateNewFeedbackDrive(fb_drive)

			if err != nil {
				utils.SendError(c, http.StatusInternalServerError, err)
				return
			}

			utils.SendResponse(c, "Drive Enabled!", map[string]any{
				"drive_id": drive_id,
			})
			return
		}

		var lastDrive models.FeedbackDrivesModel = drives[0]

		for _, drive := range drives {
			if drive.IsActive {
				lastDrive = drive
				break
			}
		}

		if lastDrive.IsActive {
			utils.SendError(c, http.StatusUnprocessableEntity, errors.New("A feedback drive is already in progress"))
			return
		}

		lastDriveDiff := lastDrive.CreatedAt.Sub(time.Now())

		if lastDrive.Type == "teacher" && lastDriveDiff.Hours() > -2190 {
			utils.SendError(c, http.StatusUnprocessableEntity, errors.New("Difference between two drives should be at least 3 months"))
			return
		}

		if lastDrive.Type == "college" && lastDriveDiff.Hours() > -2190 {
			utils.SendError(c, http.StatusUnprocessableEntity, errors.New("Difference between two drives should be at least 1 Year"))
			return
		}

		// create new drive
		fb_drive := models.FeedbackDrivesModel{
			CollegeId: college_id.(int),
			Type:      body.Type,
			IsActive:  true,
		}

		drive_id, err := models.CreateNewFeedbackDrive(fb_drive)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Drive Enabled!", map[string]any{
			"drive_id": drive_id,
		})
		return

	case "disable":
		var found bool = false
		for _, drive := range drives {
			if drive.Id == body.DriveId {
				found = true
				break
			}
		}

		if !found {
			utils.SendError(c, http.StatusNotFound, errors.New("Drive not Found"))
			return
		}

		updateFields := map[string]any{
			"is_active": false,
		}

		where := map[string]any{
			"id": body.DriveId,
		}

		_, err := models.UpdateFeedbackDriveByType(updateFields, where)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		if body.Type == "teacher" {
			err = captureFeedback(college_id.(int), body.DriveId)
		}

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Drive Disabled!", map[string]any{})
		return
	}
}

func GetDashboardMetricsController(c *gin.Context) {
	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	st_count, err := models.GetStudentsCountByCollegeId(college_id.(int))

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		st_count = 0
	}

	t_count, err := models.GetTeachersCountByCollegeId(college_id.(int))

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		t_count = 0
	}

	d_count, err := models.GetDrivesCountByCollegeId(college_id.(int))

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		d_count = 0
	}

	clg_fb_count, err := models.GetCollegeFeedbackCountByCollegeId(college_id.(int))

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		d_count = 0
	}

	utils.SendResponse(c, "College Metrics", map[string]any{
		"teachers":          t_count,
		"students":          st_count,
		"drives":            d_count,
		"college_feedbacks": clg_fb_count,
	})
}

func captureFeedback(cid int, drive_id int) error {
	topTeachers, err := models.GetTopNTeachersByType(types.Top3TeachersBody{
		N:           utils.ONE_MILLION,
		Cid:         cid,
		RequestType: "COLLEGE",
	})

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return err
	}
	var insertCaptures []models.DriveCaptureModel
	for _, teacher := range topTeachers {
		capture_opts := models.DriveCaptureModel{
			VictimId:   teacher.Id,
			VictimType: "teacher",
			DriveId:    drive_id,
			Rank:       teacher.Rank,
			IsActive:   true,
		}
		insertCaptures = append(insertCaptures, capture_opts)
	}

	// bulk insert captures
	_, err = models.BulkCreateDriveCapture(insertCaptures)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return err
	}

	return nil
}

func GetCollegeFeedbackDrivesController(c *gin.Context) {
	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	var totalDrives []models.FeedbackDrivesModel
	drives, err := models.GetCollegeFeedbackDrives(college_id.(int), "college")

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	totalDrives = append(totalDrives, drives...)

	drives, err = models.GetCollegeFeedbackDrives(college_id.(int), "teacher")

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	totalDrives = append(totalDrives, drives...)

	utils.SendResponse(c, "Feedback Drives", map[string]any{
		"drives": drives,
	})

}

func GetTeacherLinkingDataController(c *gin.Context) {
	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	students, err := models.GetStudentsByField("cid", college_id)

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	teachers, err := models.GetAllTeachersOfMyCollege(college_id.(int), utils.ONE_MILLION, 0)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	courses, err := models.GetAllCourses()

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	utils.SendResponse(c, "Linking Data", map[string]any{
		"teachers": teachers,
		"students": students,
		"courses":  courses,
	})
}
