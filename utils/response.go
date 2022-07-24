package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getError(e error) string {
	if e != nil {
		return e.Error()
	}
	return ""
}

func SendError(c *gin.Context, code int, err error) {
	switch {
	case code >= 400 && code < 500:
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": "Bad Request",
			"error":   getError(err),
		})
	case code >= 500 && code < 600:
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": "Something went wrong!",
			"error":   getError(err),
		})
	}
}

func SendValidationError(c *gin.Context, errorFieldsList []string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"data":         nil,
		"message":      "Bad Request",
		"error":        "Invalid Input",
		"error_fields": errorFieldsList,
	})
}

func SendResponse(c *gin.Context, msg string, data map[string]any) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    data,
	})
}

// sends an unauthorized http response
func SendUnauthorized(c *gin.Context, message string, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
		"error":   err.Error(),
		"data":    nil,
	})
	c.Abort()
}
