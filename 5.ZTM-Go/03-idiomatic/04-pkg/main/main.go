package main

import (
	"ztm_course/03-idiomatic/04-pkg/display"
	"ztm_course/03-idiomatic/04-pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("an exciting message")
}
