package services

import (
	"github.com/jonreiter/govader"
)

func VaderUp(str string) (govader.Sentiment) {
	analyzer := govader.NewSentimentIntensityAnalyzer()
	sentiment := analyzer.PolarityScores(str)

	return sentiment
}
