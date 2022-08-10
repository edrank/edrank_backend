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
	TeacherLinksModel struct {
		Id        int       `json:"id"`
		TeacherId int       `json:"teacher_id"`
		CollegeId int       `json:"college_id"`
		CourseId  int       `json:"course_id"`
		Year      int       `json:"year"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func GetTeacherLinksByField(fieldName string, fieldValue any) ([]TeacherLinksModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from teacher_links where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find links")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var links []TeacherLinksModel
	for rows.Next() {
		var link TeacherLinksModel

		if err := rows.Scan(&link.Id, &link.TeacherId, &link.CollegeId, &link.CourseId, &link.Year, &link.IsActive, &link.CreatedAt, &link.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		links = append(links, link)
	}
	if len(links) == 0 {
		return nil, errors.New("Cannot find links")
	}
	return links, nil
}
