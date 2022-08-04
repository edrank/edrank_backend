package models

import (
	"database/sql"
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
		Id          int     `json:"id"`
		Name        string  `json:"name"`
		Score       float32 `json:"score"`
		Rank        int     `json:"rank"`
		CollegeName string  `json:"college_name"`
	}
)

func GetTop3TeachersByType(params types.Top3TeachersBody) ([]Top3TeachersResponse, error) {
	var query string
	var rows *sql.Rows
	var err error
	database := db.GetDatabase()

	switch params.RequestType {
	case "COLLEGE":
		query = "select teachers.id, teachers.name, teachers.score, colleges.name as college_name from teachers, colleges where cid = ? AND teachers.is_active = 1 ORDER BY score DESC LIMIT 3;"
		rows, err = database.Query(query, params.Cid)
	case "STATE":
		query = "SELECT teachers.id, teachers.name, teachers.score, colleges.name as college_name FROM `teachers` inner join `colleges` on colleges.id = teachers.cid AND colleges.state = ? ORDER BY score DESC LIMIT 3;"
		rows, err = database.Query(query, params.State)
	case "REGIONAL":
		query = "SELECT teachers.id, teachers.name, teachers.score, colleges.name as college_name FROM `teachers` inner join `colleges` on colleges.id = teachers.cid AND colleges.city = ? ORDER BY score DESC LIMIT 3;"
		rows, err = database.Query(query, params.City)
	case "NATIONAL":
		query = "select teachers.id, teachers.name, teachers.score, colleges.name as college_name from teachers, colleges where teachers.is_active = 1 ORDER BY score DESC LIMIT 3;"
		rows, err = database.Query(query)
	}

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var teachers []Top3TeachersResponse
	var rank int = 1
	for rows.Next() {
		var teacher Top3TeachersResponse

		if err := rows.Scan(&teacher.Id, &teacher.Name, &teacher.Score, &teacher.CollegeName); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		teacher.Rank = rank
		rank++
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func GetAllTeachersOfMyCollege(cid int, limit int, offset int) ([]TeacherModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from teachers where cid = ? limit ? offset ?", cid, limit, offset)
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var teachers []TeacherModel
	for rows.Next() {
		var teacher TeacherModel

		if err := rows.Scan(&teacher.Id, &teacher.Cid, &teacher.Name, &teacher.OfficialEmail, &teacher.AlternateEmail, &teacher.Department, &teacher.CourseId, &teacher.Designation, &teacher.Score, &teacher.Password, &teacher.IsActive, &teacher.CreatedAt, &teacher.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}
