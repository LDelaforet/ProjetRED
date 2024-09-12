package main

import (
	RED "RED/Internals"
)

// TODO, currently main.go is a test for the Internals functions

func main() {
	//RED.DisplayText("test")
	RED.DisplayTitle("NOM DU JEU")
	RED.DisplayText("", false, 0)
	RED.DisplayText("0: Nouvelle Partie", false, 0)
	RED.DisplayText("1: Continuer", false, 0)
	RED.DisplayText("2: Option", false, 0)
	RED.DisplayText("3: Quitter", false, 0)
}
