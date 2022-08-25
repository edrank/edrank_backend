package models

import (
	"fmt"
	"strings"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

func GetResponsesOfQuestionByTeacher(qid int, f_ids []int) ([]ResponsesModel, error) {
	database := db.GetDatabase()
	args := make([]interface{}, len(f_ids)+1)
	args = append(args, qid)
	for i, id := range f_ids {
		args[i] = id
	}
	str := "select * from responses where question_id = ? and feedback_id in (?" + strings.Repeat(",?", len(args)-1) + `)`

	fmt.Println(str, len(args))
	rows, err := database.Query(str, args...)

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
