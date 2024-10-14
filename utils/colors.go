package utils

import (
	"fmt"
	"log"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintColor(color string, message string) {
	fmt.Println(color + message + Reset)
}

func PrintError(message string) {
	PrintColor(Red, message)
}

func PrintSuccess(message string) {
	PrintColor(Green, message)
}

func PrintWarning(message string) {
	PrintColor(Yellow, message)
}

func LogColor(color string, message string) {
	log.Println(color + message + Reset)
}

func LogError(message string) {
	LogColor(Red, message)
}

func LogSuccess(message string) {
	LogColor(Green, message)
}

func LogWarning(message string) {
	LogColor(Yellow, message)
}
