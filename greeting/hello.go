package greeting

import (
	"github.com/fatih/color"
)

var greeting string = color.BlueString("Golang") + " for the " + color.RedString("Brave") + " and " + color.GreenString("Skillful") + "!"

func Hello() string {
	return greeting
}
