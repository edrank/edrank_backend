package controllers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/edrank/edrank_backend/config"
	"github.com/edrank/edrank_backend/types"
	"github.com/gin-gonic/gin"
)

var signingMethod = jwt.SigningMethodHS256
var secretKey = config.TOKEN_SECRET

func LoginController (c *gin.Context) {
	a, err := GenerateTokenString(types.CustomClaims{
		TenantId: "1",
		TenantType: "STUDENT",
		IsActive: "true",
		Email: "hello@gmail.com",
	})

	if err != nil {
		c.JSON(200, gin.H{
			"data": err.Error(),
		})	
		return
	}
	c.JSON(200, gin.H{
		"data": a,
	})
}

// generate jwt using data provided as payload
func GenerateTokenString(customClaims types.CustomClaims) (string, error) {
	fmt.Println(customClaims)
	claim := types.AuthCustomClaims{jwt.StandardClaims{},customClaims}
	fmt.Println(claim)
	token := jwt.NewWithClaims(signingMethod, claim)

	// sign the token using secret key
	return token.SignedString(secretKey)
}
