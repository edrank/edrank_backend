package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/edrank/edrank_backend/apis/config"
	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/services"
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

		college, err := models.GetCollegeByField("id", ca.Cid)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
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
			Id                      int    `json:"id"`
			Cid                     int    `json:"cid"`
			Name                    string `json:"name"`
			Email                   string `json:"email"`
			CollegeOnBoardingStatus string `json:"onboarding_status"`
			IsActive                bool   `json:"is_active"`
		}{
			Id:                      ca.Id,
			Cid:                     ca.Cid,
			Name:                    ca.Name,
			Email:                   ca.Email,
			IsActive:                ca.IsActive,
			CollegeOnBoardingStatus: college.OnboardingStatus,
		}
	case utils.TenantMap["TEACHER"]:
		ca, err := models.GetTeacherByField("email", body.Email)

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
			Email:      ca.OfficialEmail,
			Cid:        ca.Cid,
		}
		user = struct {
			Id             int     `json:"id"`
			Cid            int     `json:"cid"`
			Name           string  `json:"name"`
			Email          string  `json:"email"`
			AlternateEmail string  `json:"alt_email"`
			IsActive       bool    `json:"is_active"`
			Department     string  `json:"department"`
			CourseId       int     `json:"course_id"`
			Designation    string  `json:"designation"`
			Score          float32 `json:"score"`
		}{
			Id:             ca.Id,
			Cid:            ca.Cid,
			Name:           ca.Name,
			Email:          ca.OfficialEmail,
			AlternateEmail: ca.AlternateEmail,
			IsActive:       ca.IsActive,
			Department:     ca.Department,
			Designation:    ca.Designation,
			Score:          ca.Score,
			CourseId:       ca.CourseId,
		}
	case utils.TenantMap["PARENT"]:
		p, err := models.GetParentByField("email", body.Email)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}

		if !checkPass(body.Password, p.Password) {
			utils.SendError(c, http.StatusUnauthorized, errors.New("Invalid Credentials"))
			return
		}
		tenant_id = p.Id
		cc = types.CustomClaims{
			TenantId:   p.Id,
			TenantType: tenant_type,
			IsActive:   p.IsActive,
			Email:      p.Email,
		}
		user = struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
			IsActive bool   `json:"is_active"`
		}{
			Id:       p.Id,
			Name:     p.Name,
			Email:    p.Email,
			Phone:    p.Phone,
			IsActive: p.IsActive,
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
	case utils.TenantMap["SUPER_ADMIN"]:
		p, err := models.GetRegulatorByField("email", body.Email)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}

		if !checkPass(body.Password, p.Password) {
			utils.SendError(c, http.StatusUnauthorized, errors.New("Invalid Credentials"))
			return
		}
		tenant_id = p.Id
		cc = types.CustomClaims{
			TenantId:   p.Id,
			TenantType: tenant_type,
			IsActive:   p.IsActive,
			Email:      p.Email,
		}
		user = struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
			IsActive bool   `json:"is_active"`
		}{
			Id:       p.Id,
			Name:     p.Name,
			Email:    p.Email,
			Phone:    p.Phone,
			IsActive: p.IsActive,
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
	case utils.TenantMap["PARENT"]:
		st, err := models.GetParentByField("id", tenant_id)

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

		_, err = models.UpdateParentByFields(updateFields, where)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}
	case utils.TenantMap["TEACHER"]:
		st, err := models.GetTeacherByField("id", tenant_id)

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

		_, err = models.UpdateTeacherByFields(updateFields, where)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}
	default:
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("%s change password not implemented yet", tenant_type)))
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

func GetMyCollegeForNonCollegeController(c *gin.Context) {
	var body struct {
		Cid int `json:"college_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	college, err := models.GetCollegeByField("id", body.Cid)

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
	case utils.TenantMap["TEACHER"]:
		t, err := models.GetTeacherByField("id", tenant_id)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}
		utils.SendResponse(c, "My Profile fetched!", map[string]any{
			"profile": t,
		})
	case utils.TenantMap["PARENT"]:
		p, err := models.GetParentByField("id", tenant_id)

		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}
		utils.SendResponse(c, "My Profile fetched!", map[string]any{
			"profile": p,
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

	college_id, exists := c.Get("CollegeId")

	if !exists {
		var body struct {
			Cid int `json:"cid"`
		}
		if err := c.BindJSON(&body); err != nil {
			utils.SendError(c, http.StatusBadRequest, err)
			return
		}
		college_id = body.Cid
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

	feedbackTypeToDriveTypeMap := map[string]string{
		"SC": "college",
		"HC": "college",
		"PC": "college",
		"ST": "teacher",
	}

	drives, err := models.GetCollegeFeedbackDrives(college_id.(int), feedbackTypeToDriveTypeMap[ff_type])

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	driveActive := false
	did := 0
	for _, dr := range drives {
		if dr.IsActive {
			driveActive = true
			did = dr.Id
			break
		}
	}

	if !driveActive {
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New("No active feedback drive going on"))
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
		"drive_id":  did,
	})
}

func SubmitFeedbackController(c *gin.Context) {
	tenant_type, exists := c.Get("TenantType")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	tenant_id, exists := c.Get("TenantId")

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

	switch fmt.Sprintf("%s|%s", tenant_type, ff_type) {
	case "STUDENT|ST":
		var body types.STSubmitFeedBackBody
		if err := c.BindJSON(&body); err != nil {
			utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
			return
		}

		if len(body.Feedbacks) == 0 {
			// no feedbacks
			utils.SendError(c, http.StatusBadRequest, errors.New("No feedbacks provided"))
			return
		}

		/*
			- every feedback is processed by the scoring engine to produce a score (0-100) (avg for now maybe)
			- then feedback will be inserted in the feedback table
			- then responses will be inserted in responses table (in a way that it can be used to re-evaluate the feedback and generate a score (just in case))
			- after feedback score generation...the score will again be fed to scoring engine (or can be done when first generated too) and processed along with existing score to generate final teacher/college score
			- then the updated score with be updated in scores table (and leaderboard gets updated)
		*/
		var responses []models.ResponsesModel
		var teacher_ingestion_data_map map[int]types.FeedBacksForIngestion = make(map[int]types.FeedBacksForIngestion)
		var teacher_text_feedback_map map[int]string = make(map[int]string)
		var fb_score float32
		var sa_score float32
		var err error

		for _, feedback := range body.Feedbacks {
			// for teacher_id, teacher_feedback := range feedback {
			for _, feedback_response := range feedback.Mcq {
				teacher_ingestion_data_map[feedback.TeacherID] = append(teacher_ingestion_data_map[feedback.TeacherID], struct {
					QuestionId int "json:\"question_id\""
					AnswerId   int "json:\"answer_id\""
				}{
					QuestionId: feedback_response.QuestionId,
					AnswerId:   feedback_response.AnswerId,
				})
			}
			teacher_text_feedback_map[feedback.TeacherID] = feedback.TextFeedback
		}

		for teacher_id, ingestion_data := range teacher_ingestion_data_map {
			fb_score, sa_score, err = services.GetFeedbackScore(ingestion_data, teacher_text_feedback_map[teacher_id], teacher_id)
			if err != nil {
				utils.PrintToConsole("Error processing feedback. Aborted", err.Error())
				utils.SendError(c, http.StatusInternalServerError, err)
				return
			}
			services.UpdateEntityScore("teacher", teacher_id, fb_score, sa_score)
			feedbackOpts := models.FeedbackModel{
				DriveId:       body.DriveId,
				TenantId:      tenant_id.(int),
				TenantType:    tenant_type.(string),
				TextFeedback:  teacher_text_feedback_map[teacher_id],
				FeedbackScore: fb_score, //score from scoring engine
				IsActive:      true,
				VictimId:      teacher_id,
				VictimType:    "TEACHER",
				SAScore:       sa_score,
			}
			// insert feedback to db to get ID

			feedback_id, err := models.CreateNewFeedback(feedbackOpts)

			if err != nil {
				utils.SendError(c, http.StatusInternalServerError, err)
				return
			}

			for _, _iData := range ingestion_data {
				responses = append(responses, models.ResponsesModel{
					FeedbackId: feedback_id,
					QuestionId: _iData.QuestionId,
					Answer:     _iData.AnswerId,
					IsActive:   true,
				})
			}
		}

		err = models.BulkCreateResponses(responses)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Feedback submitted!", map[string]any{})
	case "STUDENT|SC":
		var body types.SCSubmitFeedBackBody
		if err := c.BindJSON(&body); err != nil {
			utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
			return
		}

		college_id, exists := c.Get("CollegeId")

		if !exists {
			utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
			return
		}

		var responses []models.ResponsesModel
		var feedbacks_for_ingestion types.FeedBacksForIngestion
		for _, feedback_response := range body.Feedback.Mcq {
			feedbacks_for_ingestion = append(feedbacks_for_ingestion, struct {
				QuestionId int "json:\"question_id\""
				AnswerId   int "json:\"answer_id\""
			}{
				QuestionId: feedback_response.QuestionId,
				AnswerId:   feedback_response.AnswerId,
			})
		}

		fb_score, sa_score, err := services.GetFeedbackScore(feedbacks_for_ingestion, body.Feedback.TextFeedback, college_id.(int))

		if err != nil {
			utils.PrintToConsole("Error processing feedback. Aborted", err.Error())
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		services.UpdateEntityScore("college", college_id.(int), fb_score, sa_score)

		feedbackOpts := models.FeedbackModel{
			DriveId:       body.DriveId,
			TenantId:      tenant_id.(int),
			TenantType:    tenant_type.(string),
			TextFeedback:  body.Feedback.TextFeedback,
			FeedbackScore: fb_score, //score from scoring engine
			IsActive:      true,
			VictimId:      college_id.(int),
			VictimType:    "COLLEGE",
			SAScore:       sa_score,
		}
		// insert feedback to db to get ID

		feedback_id, err := models.CreateNewFeedback(feedbackOpts)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		for _, _iData := range feedbacks_for_ingestion {
			responses = append(responses, models.ResponsesModel{
				FeedbackId: feedback_id,
				QuestionId: _iData.QuestionId,
				Answer:     _iData.AnswerId,
				IsActive:   true,
			})
		}

		err = models.BulkCreateResponses(responses)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Feedback submitted!", map[string]any{})

	case "PARENT|PC":
		var body types.PCSubmitFeedBackBody
		if err := c.BindJSON(&body); err != nil {
			utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
			return
		}

		var responses []models.ResponsesModel
		var feedbacks_for_ingestion types.FeedBacksForIngestion
		for _, feedback_response := range body.Feedback.Mcq {
			feedbacks_for_ingestion = append(feedbacks_for_ingestion, struct {
				QuestionId int "json:\"question_id\""
				AnswerId   int "json:\"answer_id\""
			}{
				QuestionId: feedback_response.QuestionId,
				AnswerId:   feedback_response.AnswerId,
			})
		}

		fb_score, sa_score, err := services.GetFeedbackScore(feedbacks_for_ingestion, body.Feedback.TextFeedback, body.CollegeId)

		if err != nil {
			utils.PrintToConsole("Error processing feedback. Aborted", err.Error())
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		// services.UpdateEntityScore("college", body.CollegeId, fb_score, sa_score)

		feedbackOpts := models.FeedbackModel{
			DriveId:       body.DriveId,
			TenantId:      tenant_id.(int),
			TenantType:    tenant_type.(string),
			TextFeedback:  body.Feedback.TextFeedback,
			FeedbackScore: fb_score, //score from scoring engine
			IsActive:      true,
			VictimId:      body.CollegeId,
			VictimType:    "COLLEGE",
			SAScore:       sa_score,
		}
		// insert feedback to db to get ID

		feedback_id, err := models.CreateNewFeedback(feedbackOpts)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		for _, _iData := range feedbacks_for_ingestion {
			responses = append(responses, models.ResponsesModel{
				FeedbackId: feedback_id,
				QuestionId: _iData.QuestionId,
				Answer:     _iData.AnswerId,
				IsActive:   true,
			})
		}

		err = models.BulkCreateResponses(responses)

		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, err)
			return
		}

		utils.SendResponse(c, "Feedback submitted!", map[string]any{})

	case "HEI|HC":
		utils.SendResponse(c, "Feedback submission not implemented yet", map[string]any{})
		return
	default:
		utils.SendError(c, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("No such feedback form type of %s for tenant %s", ff_type, tenant_type)))
		return
	}
}

func GetFeedbackTeachersController(c *gin.Context) {
	college_id, exists := c.Get("CollegeId")

	if !exists {
		utils.SendError(c, http.StatusInternalServerError, errors.New("Cannot validate context"))
		return
	}

	links, err := models.GetTeacherLinksByField("college_id", college_id)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	var ids []int = make([]int, len(links))

	for i, link := range links {
		ids[i] = link.Id
	}

	var body struct {
		CourseId int `json:"course_id"`
	}

	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	teachers, err := models.GetTeachersByTeacherIds(ids, body.CourseId)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Teachers", map[string]any{
		"teachers": teachers,
	})
}

func GetMyCollegesRankController(c *gin.Context) {
	var body types.Top3TeachersBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	if utils.Find([]string{"COLLEGE", "REGIONAL", "STATE", "NATIONAL"}, body.RequestType) == -1 {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid Request Type"))
		return
	}

	rank, err := models.GetRankOfCollegeByType(body.RequestType, body.City, body.State, body.Cid)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "My college rank", map[string]any{
		"rank": rank,
	})

}

func GetCourse(c *gin.Context) {
	course_id := c.Param("id")

	if course_id == "" {
		utils.SendError(c, http.StatusBadRequest, errors.New("Invalid course id"))
		return
	}

	cid, err := strconv.Atoi(course_id)

	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	course, err := models.GetCourseById(cid)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Course", map[string]any{
		"course": course,
	})
}
