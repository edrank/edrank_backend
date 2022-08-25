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
	TeacherScoresModel struct {
		Id        int       `json:"id"`
		TeacherId int       `json:"teacher_id"`
		Score     float32   `json:"score"`
		FbCount   int       `json:"fb_count"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CollegeScoresModel struct {
		Id        int       `json:"id"`
		CollegeId int       `json:"college_id"`
		Score     float32   `json:"score"`
		FbCount   int       `json:"fb_count"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ResponsesModel struct {
		Id         int       `json:"id"`
		FeedbackId int       `json:"feedback_id"`
		QuestionId int       `json:"question_id"`
		Answer     int       `json:"answer"`
		IsActive   bool      `json:"is_active"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	FeedbackDrivesModel struct {
		Id        int       `json:"id"`
		CollegeId int       `json:"cid"`
		Type      string    `json:"type"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	FeedbackModel struct {
		Id            int       `json:"id"`
		DriveId       int       `json:"drive_id"`
		TenantId      int       `json:"tenant_id"`
		TenantType    string    `json:"tenant_type"`
		VictimId      int       `json:"victim_id"`
		VictimType    string    `json:"victim_type"`
		TextFeedback  string    `json:"text_feedback"`
		FeedbackScore float32   `json:"feedback_score"`
		SAScore       float32   `json:"sa_score"`
		IsActive      bool      `json:"is_active"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	}
)

func CreateNewFeedback(feedback FeedbackModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into feedbacks (drive_id, tenant_id, tenant_type, victim_id, victim_type, text_feedback, feedback_score, sa_score, is_active) values (?,?,?,?,?,?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(feedback.DriveId, feedback.TenantId, feedback.TenantType, feedback.VictimId, feedback.VictimType, feedback.TextFeedback, feedback.FeedbackScore, feedback.SAScore, feedback.IsActive)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	id, err := resp.LastInsertId()

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	return int(id), nil
}

func CreateNewFeedbackDrive(feedback FeedbackDrivesModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into feedback_drives (cid, type, is_active) values (?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(feedback.CollegeId, feedback.Type, feedback.IsActive)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	id, err := resp.LastInsertId()

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	return int(id), nil
}

func BulkCreateResponses(responses []ResponsesModel) error {
	database := db.GetDatabase()
	query := "insert into responses (feedback_id, question_id, answer, is_active) values (?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return err
	}

	for _, response := range responses {
		_, err := stmt.Exec(response.FeedbackId, response.QuestionId, response.Answer, response.IsActive)

		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return err
		}

	}

	return nil
}

func GetFeedbackDriveByField(fieldName string, fieldValue any) (FeedbackDrivesModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from feedback_drives where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return FeedbackDrivesModel{}, errors.New("Cannot find drive")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return FeedbackDrivesModel{}, err
	}

	var drives []FeedbackDrivesModel
	for rows.Next() {
		var drive FeedbackDrivesModel

		if err := rows.Scan(&drive.Id, &drive.CollegeId, &drive.IsActive, &drive.CreatedAt, &drive.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return FeedbackDrivesModel{}, err
		}
		drives = append(drives, drive)
	}
	if len(drives) == 0 {
		return FeedbackDrivesModel{}, errors.New("Cannot find drive")
	}
	return drives[0], nil
}

func GetTextFeedbacksOfTeacher(tid int) ([]struct {
	Feedback string `json:"text_feedback"`
	Score    string `json:"sa_score"`
}, error) {
	database := db.GetDatabase()

	rows, err := database.Query("select text_feedback, sa_score from feedbacks where victim_id = ?", tid)

	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find fb")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var fbs []struct {
		Feedback string `json:"text_feedback"`
		Score    string `json:"sa_score"`
	}

	for rows.Next() {
		var fb struct {
			Feedback string `json:"text_feedback"`
			Score    string `json:"sa_score"`
		}

		if err := rows.Scan(&fb.Feedback, &fb.Score); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		fbs = append(fbs, fb)
	}
	return fbs, nil
}

func GetScoreByTypeAndId(score_type string, id int) (float32, int, error) {
	database := db.GetDatabase()
	var query string
	switch score_type {
	case "teacher":
		query = "select score, fb_count from teacher_scores where teacher_id = ?"
		rows, err := database.Query(query, id)
		if err == sql.ErrNoRows {
			return -1, -1, errors.New("cannot find score")
		}
		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return -1, -1, err
		}

		var scores []TeacherScoresModel
		for rows.Next() {
			var t TeacherScoresModel

			if err := rows.Scan(&t.Score, &t.FbCount); err != nil {
				utils.PrintToConsole(err.Error(), "red")
				return -1, -1, err
			}
			scores = append(scores, t)
		}
		if len(scores) == 0 {
			return -1, -1, errors.New("Cannot find score")
		}
		return scores[0].Score, scores[0].FbCount, nil
	case "college":
		query = "select score, fb_count from college_scores where college_id = ?"
		rows, err := database.Query(query, id)
		if err == sql.ErrNoRows {
			return -1, -1, errors.New("cannot find score")
		}
		if err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return -1, -1, err
		}

		var scores []CollegeScoresModel
		for rows.Next() {
			var t CollegeScoresModel

			if err := rows.Scan(&t.Score, &t.FbCount); err != nil {
				utils.PrintToConsole(err.Error(), "red")
				return -1, -1, err
			}
			scores = append(scores, t)
		}
		if len(scores) == 0 {
			return -1, -1, errors.New("Cannot find score")
		}
		return scores[0].Score, scores[0].FbCount, nil
	default:
		return 0, -1, errors.New("Invalid score type")
	}
}

func UpdateScoreByType(score_type string, fieldValues map[string]any, whereValues map[string]any) (string, error) {
	database := db.GetDatabase()
	var query string
	switch score_type {
	case "teacher":
		query = "update teacher_scores set "
	case "college":
		query = "update college_scores set "
	}
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

func GetCollegeFeedbackDrives(cid int, drive_type string) ([]FeedbackDrivesModel, error) {
	database := db.GetDatabase()

	rows, err := database.Query("select * from feedback_drives where cid = ? AND type = ? ORDER BY created_at DESC", cid, drive_type)

	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find fb drive")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var fbs []FeedbackDrivesModel

	for rows.Next() {
		var fb FeedbackDrivesModel

		if err := rows.Scan(&fb.Id, &fb.CollegeId, &fb.Type, &fb.IsActive, &fb.CreatedAt, &fb.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		fbs = append(fbs, fb)
	}
	return fbs, nil
}

func UpdateFeedbackDriveByType(fieldValues map[string]any, whereValues map[string]any) (string, error) {
	database := db.GetDatabase()
	var query string = "update feedback_drives set "
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

func GetFeedbackForValidation(tenant_type string, tenant_id int, victim_id int, victim_type string, drive_id int) (bool, error) {
	database := db.GetDatabase()

	rows, err := database.Query("select id from feedbacks where tenant_id = ? AND tenant_type = ? AND victim_id = ? AND victim_type = ? AND drive_id = ?;", tenant_id, tenant_type, victim_id, victim_type, drive_id)

	if err == sql.ErrNoRows {
		return false, errors.New("Cannot find fbs")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return false, err
	}

	var fbs []FeedbackModel

	for rows.Next() {
		var fb FeedbackModel

		if err := rows.Scan(&fb.Id); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return false, err
		}
		fbs = append(fbs, fb)
	}

	return len(fbs) > 0, nil
}

func GetQuestionById(qid int) (QuestionsModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select * from questions where id = ?", qid)
	if err == sql.ErrNoRows {
		return QuestionsModel{}, errors.New("cannot find question")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return QuestionsModel{}, err
	}

	var questions []QuestionsModel
	for rows.Next() {
		var t QuestionsModel

		if err := rows.Scan(&t.Id, &t.Title, &t.Option1, &t.Option2, &t.Option3, &t.Option4, &t.Option5, &t.Type, &t.IsActive, &t.CreatedAt, &t.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return QuestionsModel{}, err
		}
		questions = append(questions, t)
	}
	if len(questions) == 0 {
		return QuestionsModel{}, errors.New("Cannot find question")
	}
	return questions[0], nil
}

func GetFeedbacksForGraph(tid int) ([]FeedbackModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query("select id from feedbacks where victim_id = ?", tid)
	if err == sql.ErrNoRows {
		return nil, errors.New("cannot find feedback")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var feedbacks []FeedbackModel
	for rows.Next() {
		var t FeedbackModel

		if err := rows.Scan(&t.Id); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		feedbacks = append(feedbacks, t)
	}
	if len(feedbacks) == 0 {
		return nil, errors.New("Cannot find feedback")
	}
	return feedbacks, nil
}

func GetFeedbacksOfGType(id int, gType string) ([]FeedbackModel, error) {
	database := db.GetDatabase()

	rows, err := database.Query("select * from feedbacks where victim_id = ? AND victim_type = ?;", id, gType)

	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find fbs")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var fbs []FeedbackModel

	for rows.Next() {
		var feedback FeedbackModel

		if err := rows.Scan(&feedback.Id, &feedback.DriveId, &feedback.TenantId, &feedback.TenantType, &feedback.VictimId, &feedback.VictimType, &feedback.TextFeedback, &feedback.FeedbackScore, &feedback.SAScore, &feedback.IsActive, &feedback.CreatedAt, &feedback.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		fbs = append(fbs, feedback)
	}

	return fbs, nil
}