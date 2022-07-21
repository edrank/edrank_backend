package types

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	TenantId   string `json:"tenant_id"`
	TenantType string `json:"tenant_type"`
	IsActive string `json:"is_active"`
	Email string `json:"email"`
}

type AuthCustomClaims struct {
	jwt.StandardClaims
	CustomClaims
}