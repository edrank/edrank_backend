package models

import (
	"time"
)

type (
	GCInputsModel struct {
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
