package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

func GetTopNTeachersByType(params types.Top3TeachersBody) ([]Top3TeachersResponse, error) {
	var query string
	var rows *sql.Rows
	var err error
	database := db.GetDatabase()

	switch params.RequestType {
	case "COLLEGE":
		query = "select teachers.id as id, teachers.name as name,teacher_scores.score as score, colleges.name as college_name from teacher_scores join teachers on teachers.id = teacher_scores.teacher_id inner join colleges on colleges.id = teachers.cid AND colleges.id = ? ORDER BY teacher_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.Cid, params.N)
	case "STATE":
		query = "select teachers.id as id, teachers.name as name,teacher_scores.score as score, colleges.name as college_name from teacher_scores join teachers on teachers.id = teacher_scores.teacher_id inner join colleges on colleges.id = teachers.cid AND colleges.state = ? ORDER BY teacher_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.State, params.N)
	case "REGIONAL":
		query = "select teachers.id as id, teachers.name as name,teacher_scores.score as score, colleges.name as college_name from teacher_scores join teachers on teachers.id = teacher_scores.teacher_id inner join colleges on colleges.id = teachers.cid AND colleges.city = ? ORDER BY teacher_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.City, params.N)
	case "NATIONAL":
		query = "select teachers.id as id, teachers.name as name,teacher_scores.score as score, colleges.name as college_name from teacher_scores join teachers on teachers.id = teacher_scores.teacher_id inner join colleges on colleges.id = teachers.cid ORDER BY teacher_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.N)
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

func UpdateTeacherByFields(fieldValues map[string]any, whereValues map[string]any) (string, error) {
	database := db.GetDatabase()
	var query string = "update teachers set "
	var values []any
	for field, value := range fieldValues {
		query += fmt.Sprintf("%s = ?, ", field)
		values = append(values, value)
	}
	query = query[:len(query)-2] + " where "

	for field, value := range whereValues {
		query += fmt.Sprintf("%s = ?, ", field)
		values = append(values, value)
	}
	query = query[:len(query)-2] + ";"

	result, err := database.Exec(query, values...)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return "", err
	}

	_, err = result.RowsAffected()

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return "", err
	}

	return "Fields Updated", nil
}

func GetTeacherByField(fieldName string, fieldValue any) (TeacherModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from teachers where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return TeacherModel{}, errors.New("cannot find teacher")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return TeacherModel{}, err
	}

	var teachers []TeacherModel
	for rows.Next() {
		var t TeacherModel

		if err := rows.Scan(&t.Id, &t.Cid, &t.Name, &t.OfficialEmail, &t.AlternateEmail, &t.Department, &t.CourseId, &t.Designation, &t.Score, &t.Password, &t.IsActive, &t.CreatedAt, &t.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return TeacherModel{}, err
		}
		teachers = append(teachers, t)
	}
	if len(teachers) == 0 {
		return TeacherModel{}, errors.New("Cannot find teacher")
	}
	return teachers[0], nil
}

func GetTeachersByTeacherIds(ids []int, course_id int) ([]TeacherModel, error) {
	database := db.GetDatabase()

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	stmt := `SELECT * from teachers where id in (?` + strings.Repeat(",?", len(args)-1) + `)` + fmt.Sprintf(" AND course_id = %d", course_id)
	rows, err := database.Query(stmt, args...)
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


func GetTeachersCountByCollegeId(cid int) (int, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select count(*) as t_count from teachers where cid = ?;", cid)
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	var count int
	for rows.Next() {

		if err := rows.Scan(&count); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return -1, err
		}
	}
	return count, nil
}