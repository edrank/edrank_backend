package utils

import (
	"strconv"

	"github.com/fatih/color"
)

func PrintToConsole(msg string, reportType string) {
	switch reportType {
	case "error":
		color.Red(msg)
	case "info":
		color.Blue(msg)
	case "log":
		color.Cyan(msg)
	case "success":
		color.Green(msg)
	}
}

func GetVersion() string {
	return "v1"
}

func Find(arr []string, val string) int {
	for i, s := range arr {
		if s == val {
			return i
		}
	}
	return -1
}

func AreEqualArray(a []string, b []string) bool {
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func GetPaginationOpts(size string, page string) (int, int, error) {
	limit, err := strconv.Atoi(size)

	if err != nil {
		return ONE_MILLION, 0, err
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return ONE_MILLION, 0, err
	}

	return limit, (offset - 1) * limit, nil
}

func GetSentimentByScore(score float32) string {
	if score >= 0 && score < 35 {
		return "Negative"
	} else if score >= 35 && score < 66 {
		return "Neutral"
	} else if score >= 66 {
		return "Positive"
	} else {
		return "Neutral"
	}
}
