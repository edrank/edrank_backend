package services

import (
	"fmt"

	"github.com/edrank/edrank_backend/apis/models"
	"github.com/edrank/edrank_backend/apis/types"
	"github.com/edrank/edrank_backend/apis/utils"
)

var (
	AnswerToWeightageMap map[int]float32 = map[int]float32{
		1: 50.00,
		2: 25.00,
		3: 10,
		4: -25.00,
		5: -50.00,
	}
	SA_WEIGHTAGE float32 = 12.50
	// client       *http.Client = &http.Client{}
)

func GetSAScore(text string) float32 {
	// bodyJsonStr := []byte(fmt.Sprintf(`{ "text_to_analyze": "%s" }`, text))
	// req, err := http.NewRequest("POST", utils.ML_ENGINE_URL+"/analyze", bytes.NewBuffer(bodyJsonStr))

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return 0.0
	// }

	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Content-Type", "application/json")

	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return 0.0
	// }

	// defer resp.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return 0.0
	// }

	// var responseObject struct {
	// 	Score float32 `json:"score"`
	// }

	// err = json.Unmarshal(bodyBytes, &responseObject)

	// if err != nil {
	// 	return 0.0
	// }
	// return responseObject.Score

	sentiment := VaderUp(text)
	mainCompound := sentiment.Compound
	if mainCompound < 0 {
		mainCompound = -mainCompound
		mainCompound = 1 - mainCompound
	} else {
		mainCompound += 1
	}

	fmt.Println("sentiment + - $ !", sentiment.Positive, sentiment.Negative, sentiment.Neutral, sentiment.Compound)
	
	return ((float32(mainCompound) * 100) / 200 * 100)
}

func GetFeedbackScore(feedbacks types.FeedBacksForIngestion, text_feedback string, teacher_id int) (float32, float32, error) {
	var finalScore float32
	var sa_score float32 = GetSAScore(text_feedback) // api call to ml_engine to get sa score
	for _, feedback := range feedbacks {
		finalScore += AnswerToWeightageMap[feedback.AnswerId]
	}
	divider := float32(len(feedbacks) * 50)
	if finalScore <= 0 {
		finalScore = 0
	}
	return (finalScore / divider) * 100, sa_score, nil
}

func UpdateEntityScore(entity_type string, id int, fb_score float32, sa_score float32) bool {
	sc, fb_count, err := models.GetScoreByTypeAndId(entity_type, id)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return false
	}

	old_score := sc

	old_sum := (float32(fb_count) * old_score)

	new_sum := old_sum + fb_score

	new_score := (new_sum) / (float32(fb_count + 1))

	left_over := 100 - new_score

	left_score_percent := (left_over * sa_score) / 100

	new_score += left_score_percent

	// update tables
	fieldValues := map[string]any{
		"score":    new_score,
		"fb_count": fb_count + 1,
	}

	where := map[string]any{
		"id": id,
	}

	_, err = models.UpdateScoreByType(entity_type, fieldValues, where)

	return err == nil
}
