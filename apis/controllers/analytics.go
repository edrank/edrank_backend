package controllers

import (
	"errors"
	"net/http"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func KBCGraphController(c *gin.Context) {
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

	var body types.KBCGraphBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	question, err := models.GetQuestionById(body.QuestionId)

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// feedback, err := models.GetFeedbackByForGraph("", 1)

	switch tenant_type {
	case "TEACHER":
		responses, err := models.GetResponsesOfQuestionByTeacher(1, tenant_id.(int))

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		var qCountMap map[int]int = map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
		}
		for _, response := range responses {
			qCountMap[response.Answer]++
		}

		var graphRes []struct {
			OptionName    string
			ResponseCount int
		} = []struct {
			OptionName    string
			ResponseCount int
		}{
			{
				OptionName:    question.Option1,
				ResponseCount: qCountMap[1],
			},
			{
				OptionName:    question.Option2,
				ResponseCount: qCountMap[2],
			},
			{
				OptionName:    question.Option3,
				ResponseCount: qCountMap[3],
			},
			{
				OptionName:    question.Option4,
				ResponseCount: qCountMap[4],
			},
			{
				OptionName:    question.Option5,
				ResponseCount: qCountMap[5],
			},
		}

		utils.SendResponse(c, "Graph Fetched!", map[string]any{
			"graph_data": graphRes,
		})

	case "COLLEGE_ADMIN":
	}
}
