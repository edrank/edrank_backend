package controllers

import (
	"fmt"

	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func OnboardFromJSController(c *gin.Context) {
	fmt.Println(c.Request.Body)
	utils.SendResponse(c, "Hello", map[string]any{
		"data": "edrank",
	})
}
