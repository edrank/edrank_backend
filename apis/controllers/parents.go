package controllers

import (
	"errors"
	"net/http"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func GetStudentsOfParent(c *gin.Context) {
	tenant_id, exists := c.Get("TenantId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	students, err := models.GetStudentsByField("parent_id", tenant_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	var sts []struct {
		Cid  int    `json:"cid"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	}

	for _, st := range students {
		sts = append(sts, struct {
			Cid  int    `json:"cid"`
			Name string `json:"name"`
			Id   int    `json:"id"`
		}{
			Cid:  st.Cid,
			Name: st.Name,
			Id:   st.Id,
		})
	}
	utils.SendResponse(c, "My Students", map[string]any{
		"students": sts,
	})
}
