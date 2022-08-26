package controllers

import (
	"errors"
	"net/http"

	"github.com/edrank/edrank_backend/apis/models"
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

	tenant_id, exists := c.Get("TenantId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	tenant_type, exists := c.Get("TenantType")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	
	}

	if utils.Find([]string{"college", "teacher"}, body.ComplaintFor) == -1 {
		utils.SendError(c, http.StatusBadRequest, errors.New("invalid 'complaint for'"))
		return
	}

	opts := models.GCInputsModel{
		TenantId:     tenant_id.(int),
		TenantType:   tenant_type.(string),
		ComplaintFor: body.ComplaintFor,
		Subject:      body.Subject,
		IsCC:         body.IsCC,
		CCResponse:   body.CCResponse,
		ProofFileId:  body.ProofFileId,
		Description:  body.Description,
		VictimId:     0,
		VictimType:   "",
		IsActive:     true,
	}
	id, err := models.CreateNewGCInput(opts)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Grievance Submitted. Will be reviewed shortly", map[string]any{
		"id": id,
	})

}
