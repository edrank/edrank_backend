package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/edrank/edrank_backend/db"
	"github.com/edrank/edrank_backend/utils"
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

var database *sql.DB

func GetAllCollegeAdminsOfCollege(cid int) ([]CollegeAdminModel, error) {
	database = db.GetDatabase()
	rows, err := database.Query("select * from college_admin where cid = ?", cid)

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

func GetAllCollegeAdminByField(fieldName string, fieldValue any) (CollegeAdminModel, error) {
	database = db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from college_admin where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return CollegeAdminModel{}, errors.New("Cannot find college admin")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return CollegeAdminModel{}, errors.New("Something went wrong!")
	}

	var college_admins []CollegeAdminModel
	for rows.Next() {
		var ca CollegeAdminModel

		if err := rows.Scan(&ca.Id, &ca.Cid, &ca.Name, &ca.Email, &ca.IsActive, &ca.Password, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return CollegeAdminModel{}, errors.New("Something went wrong!")
		}
		college_admins = append(college_admins, ca)
	}
	if len(college_admins) == 0 {
		return CollegeAdminModel{}, errors.New("Cannot find college admin")
	}
	return college_admins[0], nil
}

func GetAllCollegeAdminsByField(fieldName string, fieldValue any) ([]CollegeAdminModel, error) {
	database = db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from college_admin where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find college admin")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, errors.New("Something went wrong!")
	}

	var college_admins []CollegeAdminModel
	for rows.Next() {
		var ca CollegeAdminModel

		if err := rows.Scan(&ca.Id, &ca.Cid, &ca.Name, &ca.Email, &ca.IsActive, &ca.Password, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, errors.New("Something went wrong!")
		}
		college_admins = append(college_admins, ca)
	}
	if len(college_admins) == 0 {
		return nil, errors.New("Cannot find college admin")
	}
	return college_admins, nil
}


func UpdateCollegeAdminById(id int, college_admin CollegeAdminModel) CollegeAdminModel {
	database = db.GetDatabase()
	return CollegeAdminModel{}
}
