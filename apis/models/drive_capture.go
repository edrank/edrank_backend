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
	DriveCaptureModel struct {
		Id         int       `json:"id"`
		VictimId   int       `json:"victim_id"`
		VictimType string    `json:"vicim_type"`
		Rank       int       `json:"rank"`
		DriveId    int       `json:"drive_id"`
		IsActive   bool      `json:"is_active"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func BulkCreateDriveCapture(dcs []DriveCaptureModel) (string, error) {
	database := db.GetDatabase()
	for _, dc := range dcs {
		query := "insert into drive_captures (victim_id, victim_type, rank, drive_id, is_active) values (?,?,?,?,?);"

		stmt, err := database.Prepare(query)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return "", err
		}
		_, err = stmt.Exec(dc.VictimId, dc.VictimType, dc.Rank, dc.DriveId, dc.IsActive)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return "", err
		}
	}
	return "Inserted", nil
}

func GetDriveCaptureByField(fieldName string, fieldValue any) ([]DriveCaptureModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from drive_captures where %s = ? and victim_type = 'teacher';", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find capture")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var colleges []DriveCaptureModel
	for rows.Next() {
		var c DriveCaptureModel

		if err := rows.Scan(&c.Id,&c.VictimId,&c.VictimType,&c.Rank,&c.DriveId, &c.IsActive, &c.CreatedAt, &c.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		colleges = append(colleges, c)
	}
	if len(colleges) == 0 {
		return nil, errors.New("Cannot find capture")
	}
	return colleges, nil
}

func GetFeedbacksForReport(teacher_id int) ([]FeedbackModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from feedbacks where victim_id = ? and victim_type = 'TEACHER' ORDER BY sa_score DESC;", teacher_id)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find feedback")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var colleges []FeedbackModel
	for rows.Next() {
		var c FeedbackModel

		if err := rows.Scan(&c.Id,&c.DriveId,&c.TenantId,&c.TenantType,&c.VictimId,&c.VictimType,&c.TextFeedback,&c.FeedbackScore,&c.SAScore,&c.IsActive,&c.CreatedAt,&c.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		colleges = append(colleges, c)
	}
	if len(colleges) == 0 {
		return nil, errors.New("Cannot find feedback")
	}
	return colleges, nil
}

func GetQuestionAnswerCountForReport(teacher_id int) {

}