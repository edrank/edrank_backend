package types

type (
	FeedBacksForIngestion []struct {
		QuestionId int `json:"question_id"`
		AnswerId   int `json:"answer_id"`
	}
)
