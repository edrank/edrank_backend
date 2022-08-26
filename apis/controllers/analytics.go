package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

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

	graphType := c.Param("type")
	var gid int
	if graphType == "" || (utils.Find([]string{"college", "teacher"}, graphType) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Graph Type"))
		return
	}

	switch graphType {
	case "teacher":
		if body.TeacherId == 0 {
			utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Teacher id"))
			return
		}
		gid = body.TeacherId

	case "college":
		if body.CollegeId == 0 {
			utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid College id"))
			return
		}
		gid = body.CollegeId
	}

	feedbacks, err := models.GetFeedbacksOfGType(gid, graphType)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	var sCountMap map[string]int = map[string]int{
		"Positive": 0,
		"Negative": 0,
		"Neutral":  0,
	}
	for _, fb := range feedbacks {
		sCountMap[utils.GetSentimentByScore(fb.SAScore)]++
	}

	utils.SendResponse(c, "Sentimental Analysis Graph", map[string]any{
		"sa_graph": sCountMap,
	})
}

func GetProgessGraph(c *gin.Context) {
	teacher_id := c.Param("tid")

	if teacher_id == "" {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Teacher Id"))
		return
	}

	captures, err := models.GetDriveCaptureByField("victim_id", teacher_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	var progress map[int]int = make(map[int]int)
	var start map[int]time.Time = make(map[int]time.Time)
	for _, capture := range captures {
		progress[capture.DriveId] = capture.Rank
		start[capture.DriveId] = capture.CreatedAt
	}

	var progressArr []struct {
		DriveId   int       `json:"drive_id"`
		StartDate time.Time `json:"started_at"`
		Rank      int       `json:"rank"`
	}
	for k, v := range progress {
		progressArr = append(progressArr, struct {
			DriveId   int       "json:\"drive_id\""
			StartDate time.Time "json:\"started_at\""
			Rank      int       "json:\"rank\""
		}{
			DriveId:   v,
			Rank:      k,
			StartDate: start[k],
		})
	}

	utils.SendResponse(c, "Progress", map[string]any{
		"progress": progressArr,
	})
}
