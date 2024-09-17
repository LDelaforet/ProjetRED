package main

import (
	RED "RED/Internals"
)

// TODO, currently main.go is a test for the Internals functions

func main() {
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "Test",
	})
}
