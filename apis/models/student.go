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
	StudentModel struct {
		Id               int       `json:"id"`
		ParentId         int       `json:"parent_id"`
		Cid              int       `json:"cid"`
		Name             string    `json:"name"`
		Email            string    `json:"email"`
		Phone            string    `json:"phone"`
		CourseId         int       `json:"course_id"`
		Year             int       `json:"year"`
		Batch            string    `json:"batch"`
		Password         string    `json:"password"`
		EnrollmentNumber string    `json:"enrollment"`
		Dob              time.Time `json:"dob"`
		FathersName      string    `json:"fathers_name"`
		MotherName       string    `json:"mother_name"`
		GuardianEmail    string    `json:"guardian_email"`
		GuardianPhone    string    `json:"guardian_phone"`
		IsActive         bool      `json:"is_active"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	}
)

func CreateBulkStudents(students []StudentModel) error {
	database := db.GetDatabase()
	query := "insert into students (parent_id,cid,name,email,phone,course_id,year,batch,password,enrollment,dob,fathers_name,mother_name,guardian_email,guardian_phone,is_active) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return err
	}

	for _, student := range students {
		_, err := stmt.Exec(student.ParentId, student.Cid, student.Name, student.Email, student.Phone, student.CourseId, student.Year, student.Batch, student.Password, student.EnrollmentNumber, student.Dob, student.FathersName, student.MotherName, student.GuardianEmail, student.GuardianPhone, student.IsActive)
		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return err
		}
	}
	return nil
}

func GetStudentByField(fieldName string, fieldValue any) (StudentModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from students where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return StudentModel{}, errors.New("Cannot find student")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return StudentModel{}, err
	}

	var students []StudentModel
	for rows.Next() {
		var st StudentModel
		if err := rows.Scan(&st.Id, &st.ParentId, &st.Cid, &st.Name, &st.Email, &st.Phone, &st.CourseId, &st.Year, &st.Batch, &st.Password, &st.EnrollmentNumber, &st.Dob, &st.FathersName, &st.MotherName, &st.GuardianEmail, &st.GuardianPhone, &st.IsActive, &st.CreatedAt, &st.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return StudentModel{}, err
		}
		students = append(students, st)
	}
	if len(students) == 0 {
		return StudentModel{}, errors.New("Cannot find student")
	}
	return students[0], nil
}

func GetStudentsByField(fieldName string, fieldValue any) ([]StudentModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from students where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find student")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var students []StudentModel
	for rows.Next() {
		var st StudentModel
		if err := rows.Scan(&st.Id, &st.ParentId, &st.Cid, &st.Name, &st.Email, &st.Phone, &st.CourseId, &st.Year, &st.Batch, &st.Password, &st.EnrollmentNumber, &st.Dob, &st.FathersName, &st.MotherName, &st.GuardianEmail, &st.GuardianPhone, &st.IsActive, &st.CreatedAt, &st.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		students = append(students, st)
	}
	if len(students) == 0 {
		return nil, errors.New("Cannot find student")
	}
	return students, nil
}

func UpdateStudentByFields(fieldValues map[string]any, whereValues map[string]any) (string, error) {
	database := db.GetDatabase()
	var query string = "update students set "
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

func GetAllStudentsOfMyCollege(cid int, limit int, offset int) ([]StudentModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from students where cid = ? limit ? offset ?", cid, limit, offset)
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var students []StudentModel
	for rows.Next() {
		var st StudentModel

		if err := rows.Scan(&st.Id, &st.ParentId, &st.Cid, &st.Name, &st.Email, &st.Phone, &st.CourseId, &st.Year, &st.Batch, &st.Password, &st.EnrollmentNumber, &st.Dob, &st.FathersName, &st.MotherName, &st.GuardianEmail, &st.GuardianPhone, &st.IsActive, &st.CreatedAt, &st.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		students = append(students, st)
	}
	return students, nil
}

func GetStudentsCountByCollegeId(cid int) (int, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select count(*) as st_count from students where cid = ?;", cid)
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
