package middlewares

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/edrank/edrank_backend/apis/config"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
)

var signingMethod = jwt.SigningMethodHS256
var secretKey = config.TOKEN_SECRET

func JWTMiddleware() gin.HandlerFunc {
	const BEARER_SCHEMA = "Bearer"

	return func(c *gin.Context) {
		// get jwt stored in cookie
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			utils.SendUnauthorized(c, "Please login again.", errors.New("Invalid Token"))
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		// parse jwt token
		user, err := ParseToken(tokenString)
		if err != nil {
			utils.SendUnauthorized(c, "Please Login again. Bad Token", err)
			return
		}

		if !user.IsActive {
			utils.SendResponse(c, "Account not active", map[string]any{})
			return
		}

		// store user data in this context
		c.Set("TenantId", user.TenantId)
		c.Set("TenantType", user.TenantType)

		if utils.Find([]string{utils.TenantMap["STUDENT"], utils.TenantMap["TEACHER"], utils.TenantMap["COLLEGE_ADMIN"]}, c.GetString("TenantType")) != -1 {
			c.Set("CollegeId", user.Cid)
		}
		c.Next()
	}
}

// validate token and extract claims
func ParseToken(tokenString string) (*types.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&types.AuthCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if signingMethod.Alg() != t.Method.Alg() {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return secretKey, nil
		},
	)

	// extract claims
	claims, ok := token.Claims.(*types.AuthCustomClaims)
	if err == nil && ok && token.Valid {
		user := types.CustomClaims{Email: claims.Email, TenantId: claims.TenantId, TenantType: claims.TenantType, IsActive: claims.IsActive, Cid: claims.Cid}
		return &user, nil
	}

	return nil, errors.New("Invalid Token, parse error")
}
