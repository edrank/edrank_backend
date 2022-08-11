package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	TeacherScoresModel struct {
		Id        int       `json:"id"`
		TeacherId int       `json:"teacher_id"`
		Score     float32   `json:"score"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CollegeScoresModel struct {
		Id        int       `json:"id"`
		CollegeId int       `json:"college_id"`
		Score     float32   `json:"score"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ResponsesModel struct {
		Id         int       `json:"id"`
		FeedbackId int       `json:"feedback_id"`
		QuestionId int       `json:"question_id"`
		Answer     int       `json:"answer"`
		IsActive   bool      `json:"is_active"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	FeedbackDrivesModel struct {
		Id        int       `json:"id"`
		CollegeId int       `json:"cid"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	FeedbackModel struct {
		Id            int       `json:"id"`
		TenantId      int       `json:"tenant_id"`
		TenantType    string    `json:"tenant_type"`
		VictimId      int       `json:"victim_id"`
		TextFeedback  string    `json:"text_feedback"`
		FeedbackScore float32   `json:"feedback_score"`
		SAScore float32   `json:"sa_score"`
		IsActive      bool      `json:"is_active"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	}
)

func CreateNewFeedback(feedback FeedbackModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into feedbacks (tenant_id, tenant_type, victim_id, text_feedback, feedback_score, is_active) values (?,?,?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(feedback.TenantId, feedback.TenantType, feedback.VictimId, feedback.TextFeedback, feedback.FeedbackScore, feedback.IsActive)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	id, err := resp.LastInsertId()

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	return int(id), nil
}

func BulkCreateResponses(responses []ResponsesModel) error {
	database := db.GetDatabase()
	query := "insert into responses (feedback_id, question_id, answer, is_active) values (?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return err
	}

	for _, response := range responses {
		_, err := stmt.Exec(response.FeedbackId, response.QuestionId, response.Answer, response.IsActive)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return err
		}

	}

	return nil
}

func GetFeedbackDriveByField(fieldName string, fieldValue any) (FeedbackDrivesModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from feedback_drives where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return FeedbackDrivesModel{}, errors.New("Cannot find drive")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return FeedbackDrivesModel{}, err
	}

	var drives []FeedbackDrivesModel
	for rows.Next() {
		var drive FeedbackDrivesModel

		if err := rows.Scan(&drive.Id, &drive.CollegeId, &drive.IsActive, &drive.CreatedAt, &drive.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return FeedbackDrivesModel{}, err
		}
		drives = append(drives, drive)
	}
	if len(drives) == 0 {
		return FeedbackDrivesModel{}, errors.New("Cannot find drive")
	}
	return drives[0], nil
}
