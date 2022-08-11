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
	ParentModel struct {
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

func GetParentByField(fieldName string, fieldValue any) (ParentModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from parents where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return ParentModel{}, errors.New("cannot find parent")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return ParentModel{}, err
	}

	var parents []ParentModel
	for rows.Next() {
		var p ParentModel

		if err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Phone, &p.Password, &p.IsActive, &p.CreatedAt, &p.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return ParentModel{}, err
		}
		parents = append(parents, p)
	}
	if len(parents) == 0 {
		return ParentModel{}, errors.New("Cannot find parent")
	}
	return parents[0], nil
}

func UpdateParentByFields(fieldValues map[string]any, whereValues map[string]any) (string, error) {
	database := db.GetDatabase()
	var query string = "update parents set "
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

func GetParentsOfCollege(cid int, limit int, offset int) ([]ParentModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select parents.id,parents.name,parents.email,parents.phone,parents.password,parents.is_active,parents.created_at,parents.updated_at from parents left join students on students.parent_id = parents.id AND students.cid = ? group by parents.id limit ? offset ?", cid, limit, offset)
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var parents []ParentModel
	for rows.Next() {
		var parent ParentModel

		if err := rows.Scan(&parent.Id, &parent.Name, &parent.Email, &parent.Phone, &parent.Password, &parent.IsActive, &parent.CreatedAt, &parent.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		parents = append(parents, parent)
	}
	return parents, nil
}