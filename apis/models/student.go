package models

import "time"

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

func CreateBulkStudents(students []StudentModel) {

}
