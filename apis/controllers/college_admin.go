package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/services"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// get file
// store it locally
// validate
// manipulate
// insert in db
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

	const studentsSheet string = "Students"
	const teachersSheet string = "Teachers"

	var studentSheetValues map[string]string = map[string]string{
		"Name": "name",
		"Email":"email",
		"Phone":"phone",
		"Course":"course_id",
		"Year":"year",
		"Batch":"batch",
		"Section": "section",
		"Enrollment Number":"enrollment_number",
		"Date of Birth":"dob",
		"Fathers Name":"fathers_name",
		"Mothers Name":"mothers_name",
		"Guardian Email": "guardian_email",
		"Guardian Phone":"guardian_phone",
	}

	var teacherSheetValues map[string]string = map[string]string{
		"Name": "name",
		"Official Email":"email",
		"Alternate Email":"alt_email",
		"Department":"department",
		"Course":"course_id",
		"Designation":"designation",
	}

	studentRows, err := f.GetRows(studentsSheet)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(studentRows)

	headers := studentRows[0]
	data := studentRows[1:]

	studentsData = make(map[string]any)

	for header := range headers {

	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	// remove the downloaded file
	os.Remove("tmp/"+fr.Location)
}
