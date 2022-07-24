package types

import "github.com/dgrijalva/jwt-go"


type (
	CustomClaims struct {
		TenantId   int `json:"tenant_id"`
		TenantType string `json:"tenant_type"`
		IsActive bool `json:"is_active"`
		Email string `json:"email"`
		Cid int `json:"cid"`
	}
	AuthCustomClaims struct {
		jwt.StandardClaims
		CustomClaims
	}
	LoginBody struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
)