package main

import "fmt"

type Color string

var (
	greenBg      = Color(string([]byte{27, 91, 57, 55, 59, 52, 50, 109}))
	whiteBg      = Color(string([]byte{27, 91, 57, 48, 59, 52, 55, 109}))
	yellowBg     = Color(string([]byte{27, 91, 57, 48, 59, 52, 51, 109}))
	redBg        = Color(string([]byte{27, 91, 57, 55, 59, 52, 49, 109}))
	blueBg       = Color(string([]byte{27, 91, 57, 55, 59, 52, 52, 109}))
	magentaBg    = Color(string([]byte{27, 91, 57, 55, 59, 52, 53, 109}))
	cyanBg       = Color(string([]byte{27, 91, 57, 55, 59, 52, 54, 109}))
	green        = Color(string([]byte{27, 91, 51, 50, 109}))
	white        = Color(string([]byte{27, 91, 51, 55, 109}))
	yellow       = Color(string([]byte{27, 91, 51, 51, 109}))
	red          = Color(string([]byte{27, 91, 51, 49, 109}))
	blue         = Color(string([]byte{27, 91, 51, 52, 109}))
	magenta      = Color(string([]byte{27, 91, 51, 53, 109}))
	cyan         = Color(string([]byte{27, 91, 51, 54, 109}))
	reset        = Color(string([]byte{27, 91, 48, 109}))
	disableColor = false
)

func show(str string, color Color) {
	fmt.Println(color, str, reset)
}
