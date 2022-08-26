package models

import (
	"time"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	QuestionsModel struct {
		Id        int       `json:"id"`
		Title     string    `json:"title"`
		Option1   string    `json:"option_1"`
		Option2   string    `json:"option_2"`
		Option3   string    `json:"option_3"`
		Option4   string    `json:"option_4"`
		Option5   string    `json:"option_5"`
		Type      string    `json:"type"`
		IsActive  string    `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func GetAllQuestionsByType(ff_type string) ([]QuestionsModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from questions where type = ?", ff_type)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var questions []QuestionsModel
	for rows.Next() {
		var q QuestionsModel

		if err := rows.Scan(&q.Id, &q.Title, &q.Option1, &q.Option2, &q.Option3, &q.Option4, &q.Option5, &q.Type, &q.IsActive, &q.CreatedAt, &q.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

// func
