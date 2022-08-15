package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	CollegeModel struct {
		Id               int       `json:"id"`
		Name             string    `json:"name"`
		Email            string    `json:"email"`
		Phone            string    `json:"phone"`
		WebsiteUrl       string    `json:"website_url"`
		UniversityName   string    `json:"university_name"`
		CollegeType      string    `json:"college_type"`
		City             string    `json:"city"`
		State            string    `json:"state"`
		Score            float32   `json:"score"`
		OnboardingStatus string    `json:"onboarding_status"`
		IsActive         bool      `json:"is_active"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	}
	Top3CollegesResponse struct {
		Id    int     `json:"id"`
		Score float32 `json:"score"`
		Name  string  `json:"name"`
		City  string  `json:"city"`
		State string  `json:"state"`
		Rank  int     `json:"rank"`
	}
)

func GetCollegeByField(fieldName string, fieldValue any) (CollegeModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from colleges where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return CollegeModel{}, errors.New("Cannot find college")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return CollegeModel{}, err
	}

	var colleges []CollegeModel
	for rows.Next() {
		var c CollegeModel

		if err := rows.Scan(&c.Id, &c.Name, &c.Email, &c.Phone, &c.WebsiteUrl, &c.UniversityName, &c.CollegeType, &c.City, &c.State, &c.Score, &c.OnboardingStatus, &c.IsActive, &c.CreatedAt, &c.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return CollegeModel{}, err
		}
		colleges = append(colleges, c)
	}
	if len(colleges) == 0 {
		return CollegeModel{}, errors.New("Cannot find college")
	}
	return colleges[0], nil
}

func CreateCollege(college CollegeModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into college values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(&college.Name, &college.Email, &college.Phone, &college.WebsiteUrl, &college.UniversityName, &college.City, &college.State, &college.Score, &college.CollegeType, &college.OnboardingStatus, &college.IsActive)

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

func GetRankOfCollegeByType(rank_type string, city string, state string, cid int) (int, error) {
	var query string
	var rows *sql.Rows
	var err error
	database := db.GetDatabase()

	switch rank_type {
	case "NATIONAL":
		query = "select college_id from college_scores order by score DESC;"
		rows, err = database.Query(query)
	case "STATE":
		query = "select college_id from college_scores inner join `colleges` on colleges.id = college_scores.college_id AND colleges.state = ? order by college_scores.score DESC;"
		rows, err = database.Query(query, state)
	case "REGIONAL":
		query = "select college_id from college_scores inner join `colleges` on colleges.id = college_scores.college_id AND colleges.city = ? order by college_scores.score DESC;"
		rows, err = database.Query(query, city)
	}

	if err == sql.ErrNoRows {
		return -1, errors.New("Cannot find colleges")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	var college_scores []CollegeScoresModel
	for rows.Next() {
		var cs CollegeScoresModel

		if err := rows.Scan(&cs.CollegeId); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return -1, err
		}
		college_scores = append(college_scores, cs)
	}
	if len(college_scores) == 0 {
		return -1, errors.New("Cannot find college score")
	}

	for i, college_sc := range college_scores {
		if college_sc.CollegeId == cid {
			return i + 1, nil
		}
	}
	return -1, nil
}

func GetTopNCollegesByType(params types.Top3TeachersBody) ([]Top3CollegesResponse, error) {
	var query string
	var rows *sql.Rows
	var err error
	database := db.GetDatabase()

	switch params.RequestType {
	case "STATE":
		query = "select college_scores.id as id, college_scores.score as score, colleges.name as name, colleges.state as state, colleges.city as city from college_scores join colleges on colleges.id = college_scores.college_id AND colleges.state = ? ORDER BY college_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.State, params.N)
	case "REGIONAL":
		query = "select college_scores.id as id, college_scores.score as score, colleges.name as name, colleges.state as state, colleges.city as city from college_scores join colleges on colleges.id = college_scores.college_id AND colleges.city = ? ORDER BY college_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.City, params.N)
	case "NATIONAL":
		query = "select college_scores.id as id, college_scores.score as score, colleges.name as name, colleges.state as state, colleges.city as city from college_scores join colleges on colleges.id = college_scores.college_id ORDER BY college_scores.score DESC LIMIT ?;"
		rows, err = database.Query(query, params.N)
	}

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var colleges []Top3CollegesResponse
	var rank int = 1
	for rows.Next() {
		var college Top3CollegesResponse

		if err := rows.Scan(&college.Id, &college.Score, &college.Name, &college.State, &college.City); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		college.Rank = rank
		rank++
		colleges = append(colleges, college)
	}

	return colleges, nil
}

func GetDrivesCountByCollegeId(cid int) (int, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select count(*) as d_count from feedback_drives where cid = ?;", cid)
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

func GetCollegeFeedbackCountByCollegeId(cid int) (int, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select count(*) as clg_fb_count from feedbacks where victim_id = ? AND victim_type = 'COLLEGE';", cid)
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