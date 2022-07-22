package models

import (
	"database/sql"

	"github.com/edrank/edrank_backend/db"
	"github.com/edrank/edrank_backend/utils"
)

type (
	CollegeAdminModel struct {
		Id        int
		Cid       int
		Name      string
		Email     string
		Password  string
		IsActive  bool
		CreatedAt string
		UpdatedAt string
	}
)

var database *sql.DB = db.DB

func (collgeAdmin *CollegeAdminModel) GetAllCollegeAdminsOfCollege(cid int) []CollegeAdminModel {
	rows, err := database.Query("select * from college_admin where cid = ?", cid)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil
	}

	var college_admins []CollegeAdminModel
	for rows.Next() {
		var ca CollegeAdminModel

		if err := rows.Scan(&ca.Id, &ca.Cid, &ca.Name, &ca.Email, &ca.Password, &ca.CreatedAt, &ca.UpdatedAt, &ca.IsActive); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil
		}
		college_admins = append(college_admins, ca)
	}
	return college_admins
}

func (collgeAdmin *CollegeAdminModel) GetAllCollegeAdminsById(id int) CollegeAdminModel {
	return CollegeAdminModel{}
}

func (collgeAdmin *CollegeAdminModel) UpdateCollegeAdminById(id int, college_admin CollegeAdminModel) CollegeAdminModel {
	return CollegeAdminModel{}
}
