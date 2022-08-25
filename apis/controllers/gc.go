package controllers

import (
	"net/http"

	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func SubmitGCFormController(c *gin.Context) {
	var body types.GCSubmitBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}
	
	
}