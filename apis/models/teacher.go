package models

import (
	"time"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	TeacherModel struct {
		Id             int       `json:"id"`
		Cid            int       `json:"cid"`
		Name           string    `json:"name"`
		OfficialEmail  string    `json:"email"`
		AlternateEmail string    `json:"alt_email"`
		Department     string    `json:"department"`
		CourseId       int       `json:"course_id"`
		Designation    string    `json:"designation"`
		Score          float32   `json:"score"`
		Password       string    `json:"password"`
		IsActive       bool      `json:"is_active"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	Top3TeachersResponse struct {
		Id    int     `json:"id"`
		Name  string  `json:"name"`
		Score float32 `json:"score"`
		Rank  int     `json:"rank"`
	}
)

func GetTop3TeachersByType(params types.Top3TeachersBody) ([]Top3TeachersResponse, error) {
	var query string
	switch params.RequestType {
	case "COLLEGE":
		query = "select id, name, score from teachers where cid = ? AND is_active = 1 ORDER BY score DESC LIMIT 3"
	}
	database := db.GetDatabase()

	rows, err := database.Query(query, params.Cid)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var teachers []Top3TeachersResponse
	var rank int = 1
	for rows.Next() {
		var teacher Top3TeachersResponse

		if err := rows.Scan(&teacher.Id, &teacher.Name, &teacher.Score); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		teacher.Rank = rank
		rank++;
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}
