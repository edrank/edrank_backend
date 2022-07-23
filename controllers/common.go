package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/edrank/edrank_backend/config"
	"github.com/edrank/edrank_backend/models"
	"github.com/edrank/edrank_backend/types"
	"github.com/edrank/edrank_backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var signingMethod = jwt.SigningMethodHS256
var secretKey = config.TOKEN_SECRET

func LoginController(c *gin.Context) {
	var body types.LoginBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	tenant_type := c.GetHeader("x-edrank-tenant-type")
	var tenant_id int
	if utils.Find(utils.ValidTentantTypes[:], tenant_type) == -1 {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Tenant Type"))
		return
	}
	var cc types.CustomClaims
	var user any
	switch tenant_type {
	case utils.TenantMap["COLLEGE_ADMIN"]:
		var ca models.CollegeAdminModel

		ca, err := models.GetAllCollegeAdminByField("email", body.Email)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}

		if !ca.IsActive {
			utils.SendError(c, http.StatusUnprocessableEntity, errors.New("Account is not active"))
			return
		}

		if !checkPass(body.Password, ca.Password) {
			utils.SendError(c, http.StatusUnauthorized, errors.New("Invalid Credentials"))
			return
		}
		tenant_id = ca.Id
		cc = types.CustomClaims{
			TenantId:   ca.Id,
			TenantType: tenant_type,
			IsActive:   ca.IsActive,
			Email:      ca.Email,
		}
		user = struct {
			Id       int    `json:"id"`
			Cid      int    `json:"cid"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			IsActive bool   `json:"is_active"`
		}{
			Id:       ca.Id,
			Cid:      ca.Cid,
			Name:     ca.Name,
			Email:    ca.Email,
			IsActive: ca.IsActive,
		}
	default:
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("%s login not implemented yet", tenant_type)))
		return
	}

	token, err := GenerateTokenString(cc)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, map[string]any{
		"tenant_type":  tenant_type,
		"tenant_id":    tenant_id,
		"access_token": token,
		"user":         user,
	})
	return
}

func ForgetPasswordController(c *gin.Context) {

}

func ChangePasswordController(c *gin.Context) {
	var body types.ChangePasswordTypes
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	tenant_type := c.GetString("TenantType")
	tenant_id := c.GetInt("TenantId")

	switch tenant_type {
	case utils.TenantMap["COLLEGE_ADMIN"]:
		var ca models.CollegeAdminModel

		ca, err := models.GetAllCollegeAdminByField("id", tenant_id)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}

		if !ca.IsActive {
			utils.SendError(c, http.StatusUnprocessableEntity, errors.New("Account is not active"))
			return
		}

		if !checkPass(body.OldPassword, ca.Password) {
			utils.SendError(c, http.StatusUnauthorized, errors.New("Old password doesn't match"))
			return
		}

		var hashedPassword []byte
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(body.NewPassword), 14)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		var updateFields = map[string]any{
			"password": string(hashedPassword),
		}

		var where = map[string]any{
			"id": tenant_id,
		}

		_, err = models.UpdateCollegeAdminByFields(updateFields, where)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, errors.New("Something went wrong"))
			return
		}

	default:
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("%s login not implemented yet", tenant_type)))
		return
	}

	utils.SendResponse(c, map[string]any{
		"message":     "Password changed successfully!",
		"tenant_type": tenant_type,
	})
}

// generate jwt using data provided as payload
func GenerateTokenString(customClaims types.CustomClaims) (string, error) {
	claim := types.AuthCustomClaims{jwt.StandardClaims{}, customClaims}
	token := jwt.NewWithClaims(signingMethod, claim)

	// sign the token using secret key
	return token.SignedString(secretKey)
}

func checkPass(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
