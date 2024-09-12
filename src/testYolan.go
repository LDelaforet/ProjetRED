package main

import (
	RED "RED/Internals"
)

func main() {
	DisplayMainMenu()
}

func DisplayMainMenu() {
	RED.DisplayTitle("NOM DU JEU")
	RED.DisplayText("", false, 0)
	RED.DisplayText("0: Nouvelle Partie", false, 0)
	RED.DisplayText("1: Continuer", false, 0)
	RED.DisplayText("2: Option", false, 0)
	RED.DisplayText("3: Quitter", false, 0)
	RED.NewLine(1)
	RED.DisplayLine()
}
