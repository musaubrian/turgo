package util

import "github.com/fatih/color"

var (
	bold    = color.New(color.Bold)
	success = color.New(color.Bold, color.FgGreen)
)

func HeaderMsg(msg string) {
	bold.Println("\n// ", msg)
}

func Success(msg string) {
	success.Printf("\n%s\n", msg)
}
