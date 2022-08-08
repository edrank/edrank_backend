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

	if body.NewPassword == body.OldPassword {
		utils.SendError(c, http.StatusBadRequest, errors.New("New password cannot be same as old password"))
		return
	}

	switch tenant_type {
	case utils.TenantMap["COLLEGE_ADMIN"]:
		var ca models.CollegeAdminModel

		ca, err := models.GetCollegeAdminByField("id", tenant_id)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
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

	case utils.TenantMap["STUDENT"]:
		st, err := models.GetStudentByField("id", tenant_id)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}

		if !checkPass(body.OldPassword, st.Password) {
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

		_, err = models.UpdateStudentByFields(updateFields, where)

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

func GetMyProfile(c *gin.Context) {
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
		utils.SendResponse(c, "My Profile fetched!", map[string]any{
			"profile": ca,
		})
	case utils.TenantMap["STUDENT"]:
		st, err := models.GetStudentByField("id", tenant_id)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}
		utils.SendResponse(c, "My Profile fetched!", map[string]any{
			"profile": st,
		})
	default:
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("my profile for %s is not implemented yet", tenant_type)))
		return
	}
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

func GetFeedbackQuestionsController(c *gin.Context) {
	tenant_type, exists := c.Get("TenantType")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	ff_type := c.Param("type")

	if ff_type == "" || (utils.Find(utils.ValidFeedbackFormTypes[:], ff_type) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	if tenant_type == utils.TenantMap["PARENT"] && ff_type != "PC" {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	if tenant_type == utils.TenantMap["STUDENT"] && (utils.Find([]string{"ST", "SC"}, ff_type) == -1) {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	if tenant_type == utils.TenantMap["HEI"] && ff_type != "HC" {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Invalid Feedback Form Type"))
		return
	}

	questions, err := models.GetAllQuestionsByType(ff_type)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Feedback Questions", map[string]any{
		"questions": questions,
		"type":      ff_type,
	})
}
