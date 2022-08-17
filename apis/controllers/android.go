package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/services"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func AndroidSubmitFeedbackController(c *gin.Context) {
	tenant_type, exists := c.Get("TenantType")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	tenant_id, exists := c.Get("TenantId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	ff_type := c.Param("type")

	if ff_type == "" || (utils.Find(utils.ValidFeedbackFormTypes[:], ff_type) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	if tenant_type == utils.TenantMap["PARENT"] && ff_type != "PC" {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	if tenant_type == utils.TenantMap["STUDENT"] && (utils.Find([]string{"ST", "SC"}, ff_type) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	if tenant_type == utils.TenantMap["HEI"] && ff_type != "HC" {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	var body types.AndroidSubmitFeedbackBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	if len(body.Feedback.Mcq) == 0 {
		// no feedbacks
		utils.SendError(c, http.StatusBadRequest, errors.New("No feedbacks provided"))
		return
	}

	var t_type string
	var t_id int
	var err error
	switch ff_type {
	case "ST":
		if body.TeacherID == 0 {
			utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Teacher Id"))
			return
		}
		t_type = "teacher"
		t_id = body.TeacherID
	case "SC", "PC":
		if body.CollegeId == 0 {
			utils.SendError(c, http.StatusBadRequest, errors.New("Invalid College Id"))
			return
		}
		t_type = "college"
		t_id = body.CollegeId
	}

	feedbackGiven, err := models.GetFeedbackForValidation(tenant_type.(string), tenant_id.(int), t_id, t_type, body.DriveId)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	if feedbackGiven {
		utils.SendError(c, http.StatusBadRequest, errors.New("Feedback already provided in this drive"))
		return
	}

	var responses []models.ResponsesModel
	var feedbacks_for_ingestion types.FeedBacksForIngestion

	for _, feedback_response := range body.Feedback.Mcq {
		feedbacks_for_ingestion = append(feedbacks_for_ingestion, struct {
			QuestionId int "json:\"question_id\""
			AnswerId   int "json:\"answer_id\""
		}{
			QuestionId: feedback_response.QuestionId,
			AnswerId:   feedback_response.AnswerId,
		})
	}

	fb_score, sa_score, err := services.GetFeedbackScore(feedbacks_for_ingestion, body.Feedback.TextFeedback, t_id)

	if err != nil {
		utils.PrintToConsole("Error processing feedback. Aborted", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	services.UpdateEntityScore(t_type, t_id, fb_score, sa_score)

	feedbackOpts := models.FeedbackModel{
		DriveId:       body.DriveId,
		TenantId:      tenant_id.(int),
		TenantType:    tenant_type.(string),
		TextFeedback:  body.Feedback.TextFeedback,
		FeedbackScore: fb_score, //score from scoring engine
		IsActive:      true,
		VictimId:      t_id,
		VictimType:    strings.ToUpper(t_type),
		SAScore:       sa_score,
	}
	// insert feedback to db to get ID

	feedback_id, err := models.CreateNewFeedback(feedbackOpts)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	for _, _iData := range feedbacks_for_ingestion {
		responses = append(responses, models.ResponsesModel{
			FeedbackId: feedback_id,
			QuestionId: _iData.QuestionId,
			Answer:     _iData.AnswerId,
			IsActive:   true,
		})
	}

	err = models.BulkCreateResponses(responses)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Feedback submitted!", map[string]any{})

}
