package greeting

import (
	"github.com/fatih/color"
)

var greeting string = color.BlueString("Golang") + " for " + color.RedString("Brave") + "!"

func Hello() string {
	return greeting
}
