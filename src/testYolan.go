package main

import (
	RED "RED/Internals"
)

func main() {
	DisplayMainMenu()
}

func DisplayMainMenu() {
	RED.ClearScreen()
	RED.DisplayText("\n ██████╗  █████╗ ███╗   ███╗███████╗    ███╗   ██╗ █████╗ ███╗   ███╗███████╗\n██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ████╗  ██║██╔══██╗████╗ ████║██╔════╝\n██║  ███╗███████║██╔████╔██║█████╗      ██╔██╗ ██║███████║██╔████╔██║█████╗\n██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║╚██╗██║██╔══██║██║╚██╔╝██║██╔══╝\n╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗\n ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝", false, 0)
	RED.NewLine(1)
	RED.DisplayText("0: Nouvelle Partie", false, 0)
	RED.DisplayText("1: Charger la partie", false, 0)
	RED.DisplayText("2: Paramètres", false, 0)
	RED.DisplayText("3: Quitter", false, 0)
	RED.NewLine(1)
	RED.DisplayLine()
	input := RED.GetInput()
	RED.DisplayText(input, false, 0)
	if input == "0" {
		DisplayNewGameMenu()
	}
	if input == "2" {
		DisplayOptionMenu()
	}
	if input == "3" {
		RED.ClearScreen()
	}
}

func DisplayOptionMenu() {
	RED.ClearScreen()
	RED.DisplayTitle("PARAMETRES")
	RED.NewLine(2)
	RED.DisplayText("0: Changer de langue", false, 0)
	RED.DisplayText("1: Retour", false, 0)
	RED.NewLine(2)
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "1" {
		DisplayMainMenu()
	}
}

func DisplayNewGameMenu() {
	RED.ClearScreen()
	RED.DisplayTitle("NOUVELLE PARTIE")
	RED.NewLine(1)
	RED.DisplayText("Souhaites tu commencer une nouvelle partie?", true, 0)
	RED.NewLine(1)
	RED.DisplayText("0: Comfirmer", true, 0)
	RED.DisplayText("1: Retour", true, 0)
	RED.NewLine(1)
	RED.DisplayText("/!\\ Commencer une nouvelle partie supprimera la précédente /!\\", true, 0)
	RED.NewLine(1)
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "1" {
		DisplayMainMenu()
	}
}
