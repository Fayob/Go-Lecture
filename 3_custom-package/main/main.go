package main

import (
	"lecture/3_custom-package/display"
	"lecture/3_custom-package/msg"
	"lecture/3_custom-package/comment"
)

func main()  {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("An exciting message")
	comment.Comment()
}