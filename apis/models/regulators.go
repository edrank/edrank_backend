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
	RegulatorModel struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Password  string    `json:"password"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func GetRegulatorByField(fieldName string, fieldValue any) (RegulatorModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from regulators where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return RegulatorModel{}, errors.New("cannot find regulator")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return RegulatorModel{}, err
	}

	var parents []RegulatorModel
	for rows.Next() {
		var p RegulatorModel

		if err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Phone, &p.Password, &p.IsActive, &p.CreatedAt, &p.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return RegulatorModel{}, err
		}
		parents = append(parents, p)
	}
	if len(parents) == 0 {
		return RegulatorModel{}, errors.New("Cannot find regulator")
	}
	return parents[0], nil
}
