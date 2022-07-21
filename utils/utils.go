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