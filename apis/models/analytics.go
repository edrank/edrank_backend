package models

import (
	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

func GetResponsesOfQuestionByTeacher(qid int, tid int) ([]ResponsesModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from responses where question_id = ? and victim_id = ?", qid, tid)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var responses []ResponsesModel
	for rows.Next() {
		var r ResponsesModel

		if err := rows.Scan(&r.Id, &r.FeedbackId, &r.QuestionId, &r.Answer, &r.IsActive, &r.CreatedAt, &r.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		responses = append(responses, r)
	}
	return responses, nil
}
