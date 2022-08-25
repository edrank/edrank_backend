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
	CollegeAdminModel struct {
		Id        int       `json:"id"`
		Cid       int       `json:"cid"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		IsActive  bool      `json:"is_active"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func GetAllCollegeAdminsOfCollege(cid int, limit int, offset int) ([]CollegeAdminModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from college_admin where cid = ? limit ? offset ?", cid, limit, offset)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var college_admins []CollegeAdminModel
	for rows.Next() {
		var ca CollegeAdminModel

		if err := rows.Scan(&ca.Id, &ca.Cid, &ca.Name, &ca.Email, &ca.IsActive, &ca.Password, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		college_admins = append(college_admins, ca)
	}
	return college_admins, nil
}

func GetCollegeAdminByField(fieldName string, fieldValue any) (CollegeAdminModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from college_admin where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return CollegeAdminModel{}, errors.New("Cannot find college admin")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return CollegeAdminModel{}, err
	}

	var college_admins []CollegeAdminModel
	for rows.Next() {
		var ca CollegeAdminModel

		if err := rows.Scan(&ca.Id, &ca.Cid, &ca.Name, &ca.Email, &ca.IsActive, &ca.Password, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return CollegeAdminModel{}, err
		}
		college_admins = append(college_admins, ca)
	}
	if len(college_admins) == 0 {
		return CollegeAdminModel{}, errors.New("Cannot find college admin")
	}
	return college_admins[0], nil
}

func GetCollegeAdminsByField(fieldName string, fieldValue any) ([]CollegeAdminModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from college_admin where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find college admin")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var college_admins []CollegeAdminModel
	for rows.Next() {
		var ca CollegeAdminModel

		if err := rows.Scan(&ca.Id, &ca.Cid, &ca.Name, &ca.Email, &ca.IsActive, &ca.Password, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		college_admins = append(college_admins, ca)
	}
	if len(college_admins) == 0 {
		return nil, errors.New("Cannot find college admin")
	}
	return college_admins, nil
}

func UpdateCollegeAdminByFields(fieldValues map[string]any, whereValues map[string]any) (string, error) {
	database := db.GetDatabase()
	var query string = "update college_admin set "
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

func CreateCollegeAdmin(college_admin CollegeAdminModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into college_admin (cid, name, email, is_active, password) values (?,?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(college_admin.Cid, college_admin.Name, college_admin.Email, college_admin.IsActive, college_admin.Password)

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
