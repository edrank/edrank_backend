package controllers

import (
	"errors"
	"net/http"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func Top3TeachersController(c *gin.Context) {
	var body types.Top3TeachersBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	if utils.Find([]string{"COLLEGE", "REGIONAL", "STATE", "NATIONAL"}, body.RequestType) == -1 {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Request Type"))
		return
	}

	top3, err := models.GetTop3TeachersByType(body)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Top 3 Teachers", map[string]any{
		"teachers": top3,
	})
}
