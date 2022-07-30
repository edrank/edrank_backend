package models

import "time"

type (
	TeacherModel struct {
		Id             int       `json:"id"`
		Cid            int       `json:"cid"`
		Name           string    `json:"name"`
		OfficialEmail  string    `json:"email"`
		AlternateEmail string    `json:"alt_email"`
		Department     string    `json:"department"`
		CourseId       int       `json:"course_id"`
		Designation    string    `json:"designation"`
		Password       string    `json:"password"`
		IsActive       bool      `json:"is_active"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)
