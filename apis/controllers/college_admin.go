package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/services"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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
