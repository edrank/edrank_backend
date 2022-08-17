package controllers

import (
	"errors"
	"net/http"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func TopNTeachersController(c *gin.Context) {
	var body types.Top3TeachersBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	if utils.Find([]string{"COLLEGE", "REGIONAL", "STATE", "NATIONAL"}, body.RequestType) == -1 {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Request Type"))
		return
	}

	if body.N == -1 {
		body.N = utils.ONE_MILLION
	}
	top3, err := models.GetTopNTeachersByType(body)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Top 3 Teachers", map[string]any{
		"teachers": top3,
	})
}

func GetMyTextFeedbacksController(c *gin.Context) {
	tenant_id, exists := c.Get("TenantId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	fbs, err := models.GetTextFeedbacksOfTeacher(tenant_id.(int))

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Your feeedbacks", map[string]any{
		"feedbacks": fbs,
	})
}

func GetMyRankController(c *gin.Context) {
	tenant_id, exists := c.Get("TenantId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	rank_type := c.Param("rank_type")

	if rank_type == "" {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Rank Type"))
		return
	}

	if utils.Find([]string{"COLLEGE", "REGIONAL", "STATE", "NATIONAL"}, rank_type) == -1 {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Rank Type"))
		return
	}

	college, err := models.GetCollegeByField("id", college_id.(int))

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	top, err := models.GetTopNTeachersByType(types.Top3TeachersBody{
		RequestType: rank_type,
		Cid:         college_id.(int),
		N:           utils.ONE_MILLION,
		City:        college.City,
		State:       college.State,
	})
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	var rank int
	for _, t := range top {
		if t.Id == tenant_id.(int) {
			rank = t.Rank
			break
		}
	}

	utils.SendResponse(c, "My Rank", map[string]any{
		"rank": rank,
	})
}
