package middlewares

import (
	"errors"
	"net/http"

	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

func VerifyTenants(tenants []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.Find(tenants, c.GetString("TenantType")) == -1 {
			utils.SendError(c, http.StatusUnauthorized, errors.New("You are not authorized to use this API"))
			c.Abort()
		}
		c.Next()
	}
}
