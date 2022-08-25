package models

import (
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
		query := "insert into drive_captures (victim_id, vicim_type, rank, drive_id, is_active) values (?,?,?,?,?);"

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
