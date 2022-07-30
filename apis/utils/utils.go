package utils

import (
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