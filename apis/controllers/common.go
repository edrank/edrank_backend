package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/edrank/edrank_backend/apis/config"
	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var signingMethod = jwt.SigningMethodHS256
var secretKey = config.TOKEN_SECRET

func LoginController(c *gin.Context) {
	var body types.LoginBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	// bytes, err := bcrypt.GenerateFromPassword([]byte(body.Email), 14)
	// utils.SendResponse(c, "", map[string]any{
	// 	"a": string(bytes),
	// })
	// return

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
		ca, err := models.GetCollegeAdminByField("email", body.Email)

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
			Cid:        ca.Cid,
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
	case utils.TenantMap["STUDENT"]:
		st, err := models.GetStudentByField("email", body.Email)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}

		if !st.IsActive {
			utils.SendError(c, http.StatusUnprocessableEntity, errors.New("Account is not active"))
			return
		}

		if !checkPass(body.Password, st.Password) {
			utils.SendError(c, http.StatusUnauthorized, errors.New("Invalid Credentials"))
			return
		}
		tenant_id = st.Id
		cc = types.CustomClaims{
			TenantId:   st.Id,
			TenantType: tenant_type,
			IsActive:   st.IsActive,
			Email:      st.Email,
			Cid:        st.Cid,
		}
		user = struct {
			Id               int       `json:"id"`
			ParentId         int       `json:"parent_id"`
			Cid              int       `json:"cid"`
			Name             string    `json:"name"`
			Email            string    `json:"email"`
			Phone            string    `json:"phone"`
			CourseId         int       `json:"course_id"`
			Year             int       `json:"year"`
			Batch            string    `json:"batch"`
			Password         string    `json:"password"`
			EnrollmentNumber string    `json:"enrollment"`
			Dob              time.Time `json:"dob"`
			FathersName      string    `json:"fathers_name"`
			MotherName       string    `json:"mother_name"`
			GuardianEmail    string    `json:"guardian_email"`
			GuardianPhone    string    `json:"guardian_phone"`
			IsActive         bool      `json:"is_active"`
		}{
			Id:               st.Id,
			Cid:              st.Cid,
			ParentId:         st.ParentId,
			Name:             st.Name,
			Email:            st.Email,
			Phone:            st.Phone,
			CourseId:         st.CourseId,
			Year:             st.Year,
			Batch:            st.Batch,
			Password:         st.Password,
			EnrollmentNumber: st.EnrollmentNumber,
			Dob:              st.Dob,
			FathersName:      st.FathersName,
			MotherName:       st.MotherName,
			GuardianEmail:    st.GuardianEmail,
			GuardianPhone:    st.GuardianPhone,
			IsActive:         st.IsActive,
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

	utils.SendResponse(c, "Logged In", map[string]any{
		"tenant_type":  tenant_type,
		"tenant_id":    tenant_id,
		"access_token": token,
		"user":         user,
	})
}

func ForgetPasswordController(c *gin.Context) {

}

func ChangePasswordController(c *gin.Context) {
	var body types.ChangePasswordBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	tenant_type := c.GetString("TenantType")
	tenant_id := c.GetInt("TenantId")

	switch tenant_type {
	case utils.TenantMap["COLLEGE_ADMIN"]:
		var ca models.CollegeAdminModel

		ca, err := models.GetCollegeAdminByField("id", tenant_id)

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
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

	default:
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("%s login not implemented yet", tenant_type)))
		return
	}

	utils.SendResponse(c, "Password changed successfully!", map[string]any{
		"tenant_type": tenant_type,
	})
}

func GetCollegeController(c *gin.Context) {
	var college models.CollegeModel
	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New("You are not linked to any college"))
		return
	}

	college, err := models.GetCollegeByField("id", college_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	utils.SendResponse(c, "Fetched College", map[string]any{
		"college": college,
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
