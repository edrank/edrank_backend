package models

import (
	"database/sql"
	"errors"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	CourseModel struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		Abbreviation    string `json:"abbreviation"`
		DurationInYears int    `json:"duration_in_years"`
	}
)

// TODO: have to optimize this
func GetAllCourses() ([]CourseModel, error) {
	database := db.GetDatabase()

	rows, err := database.Query("select * from ref_courses")

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var courses []CourseModel
	for rows.Next() {
		var c CourseModel

		if err := rows.Scan(&c.Id, &c.Name, &c.Abbreviation, &c.DurationInYears); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func GetCourseById(id int) (CourseModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from ref_courses where id = ?", id)
	if err == sql.ErrNoRows {
		return CourseModel{}, errors.New("cannot find course")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return CourseModel{}, err
	}

	var courses []CourseModel
	for rows.Next() {
		var t CourseModel

		if err := rows.Scan(&t.Id, &t.Name, &t.Abbreviation, &t.DurationInYears); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return CourseModel{}, err
		}
		courses = append(courses, t)
	}
	if len(courses) == 0 {
		return CourseModel{}, errors.New("Cannot find course")
	}
	return courses[0], nil
}
