package controllers

import (
	"errors"
	"fmt"
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

	fmt.Println(body)

	question, err := models.GetQuestionById(body.QuestionId)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	if tenant_type == "COLLEGE_ADMIN" {
		if body.TeacherId == 0 {
			utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Teacher Id"))
			return
		}
		tenant_id = body.TeacherId
	}

	feedbacks, err := models.GetFeedbacksForGraph(tenant_id.(int))

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	var f_ids []int

	for _, fb := range feedbacks {
		f_ids = append(f_ids, fb.Id)
	}

	responses, err := models.GetResponsesOfQuestionByTeacher(body.QuestionId, f_ids)

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
		fmt.Println(response.Answer, response.Id, response.QuestionId)
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
}

func GetSAGraphController(c *gin.Context) {
	var body types.SAGraphBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	
}