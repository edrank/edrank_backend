package controllers

import (
	"net/http"
	"strconv"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func GetReportDataController(c *gin.Context) {
	teacher_id := 2

	tid := "2"

	teacher, err := models.GetTeacherByField("id", teacher_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	course, err := models.GetCourseById(teacher.CourseId)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	college, err := models.GetCollegeByField("id", teacher.Cid)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// assessment period

	captures, err := models.GetDriveCaptureByField("victim_id", teacher_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	var q1r string = ""
	var q2r string = ""
	for _, cap := range captures {
		if q1r != "" && q2r != "" {
			break
		}
		if cap.DriveId == 1 {
			q1r = strconv.Itoa(cap.Rank)
		}
		if cap.DriveId == 2 {
			q2r = strconv.Itoa(cap.Rank)
		}
	}

	feedbacks, err := models.GetFeedbacksForReport(teacher_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	var top3Positive []string

	var q1tfb int = 0
	var q2tfb int = 0

	var q1pc int = 0
	var q1nc int = 0
	var q1ngc int = 0

	var q2pc int = 0
	var q2nc int = 0
	var q2ngc int = 0
	for i, fb := range feedbacks {
		if i <= 2 {
			top3Positive = append(top3Positive, fb.TextFeedback)
		}
		if fb.DriveId == 1 {
			q1tfb++
			switch utils.GetSentimentByScore(fb.SAScore) {
			case "Positive":
				q1pc++
			case "Negative":
				q1ngc++
			case "Neutral":
				q1nc++
			}
		} else if fb.DriveId == 2 {
			q2tfb++
			switch utils.GetSentimentByScore(fb.SAScore) {
			case "Positive":
				q2pc++
			case "Negative":
				q2ngc++
			case "Neutral":
				q2nc++
			}
		}
	}
	if q1r == "" {
		q1r = "7"
	}
	if q2r == "" {
		q2r = "9"
	}
	responseData := map[string]any{
		"teacher": map[string]string{
			"t_id":        tid,
			"name":        teacher.Name,
			"designation": teacher.Designation,
			"course":      course.Name,
			"email":       teacher.OfficialEmail,
			"contact":     teacher.AlternateEmail,
		},
		"college": map[string]string{
			"name":    college.Name,
			"address": college.City + ", " + college.State,
		},
		"assessment_period": map[string]string{
			"from": "Jan, 2022",
			"to":   "July, 2022",
		},
		"quater_1_rank": q1r,
		"quater_2_rank": q2r,
		"q1_feedback": map[string]int{
			"total_feedbacks": q1tfb,
			"positive":        q1pc,
			"negative":        q1ngc,
			"neutral":         q1nc,
		},
		"q2_feedback": map[string]int{
			"total_feedbacks": q2tfb,
			"positive":        q2pc,
			"negative":        q2ngc,
			"neutral":         q2nc,
		},
		"top_feedbacks": map[string][]string{
			"positive": top3Positive,
		},
	}

	utils.SendResponse(c, "Report Data", map[string]any{
		"data": responseData,
	})
}
