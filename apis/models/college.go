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
		OnboardingStatus string    `json:"onboarding_status"`
		IsActive         bool      `json:"is_active"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
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

		if err := rows.Scan(&c.Id, &c.Name, &c.Email, &c.Phone, &c.WebsiteUrl, &c.UniversityName, &c.CollegeType, &c.City, &c.State, &c.OnboardingStatus, &c.IsActive, &c.CreatedAt, &c.UpdatedAt); err != nil {
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
