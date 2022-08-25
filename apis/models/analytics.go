package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

func GetResponsesOfQuestionByTeacher(qid int, f_ids []int) ([]ResponsesModel, error) {
	database := db.GetDatabase()
	args := make([]interface{}, len(f_ids))
	// args = append(args, qid)
	for _, id := range f_ids {
		args = append(args, id)
	}
	str := "select * from responses where question_id = " + strconv.Itoa(qid) + " and feedback_id in (?" + strings.Repeat(",?", len(args)-1) + `)`

	fmt.Println(str, len(args), args)
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
