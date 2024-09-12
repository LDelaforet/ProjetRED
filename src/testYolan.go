package main

import (
	RED "RED/Internals"

	"github.com/fatih/color"
)

func main() {
	DisplayMainMenu()
}

func DisplayMainMenu() {
	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "\n ██████╗  █████╗ ███╗   ███╗███████╗    ███╗   ██╗ █████╗ ███╗   ███╗███████╗\n██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ████╗  ██║██╔══██╗████╗ ████║██╔════╝\n██║  ███╗███████║██╔████╔██║█████╗      ██╔██╗ ██║███████║██╔████╔██║█████╗\n██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║╚██╗██║██╔══██║██║╚██╔╝██║██╔══╝\n╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗\n ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝",
	})

	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: Nouvelle Partie",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: Charger la partie",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "2: Paramètres",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "3: Quitter",
	})
	RED.NewLine(2)
	RED.DisplayLine()
	input := RED.GetInput()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: input,
	})

	//Crée une nouvelle partie
	if input == "0" {
		DisplayNewGameMenu()
	}
	//Affiche les paramètres
	if input == "2" {
		DisplayOptionMenu()
	}
	//Quitte le jeu
	if input == "3" {
		RED.ClearScreen()
	}
}

func DisplayNewGameMenu() {
	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "\n ██████╗  █████╗ ███╗   ███╗███████╗    ███╗   ██╗ █████╗ ███╗   ███╗███████╗\n██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ████╗  ██║██╔══██╗████╗ ████║██╔════╝\n██║  ███╗███████║██╔████╔██║█████╗      ██╔██╗ ██║███████║██╔████╔██║█████╗\n██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║╚██╗██║██╔══██║██║╚██╔╝██║██╔══╝\n╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗\n ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝",
	})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "Souhaites tu commencer une nouvelle partie?",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: Comfirmer",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: Retour",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "/!\\ Commencer une nouvelle partie supprimera la précédente /!\\",
		FgColor:     color.FgRed,
	})
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "1" {
		DisplayMainMenu()
	}
}

func DisplayOptionMenu() {
	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "\n ██████╗  █████╗ ███╗   ███╗███████╗    ███╗   ██╗ █████╗ ███╗   ███╗███████╗\n██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ████╗  ██║██╔══██╗████╗ ████║██╔════╝\n██║  ███╗███████║██╔████╔██║█████╗      ██╔██╗ ██║███████║██╔████╔██║█████╗\n██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║╚██╗██║██╔══██║██║╚██╔╝██║██╔══╝\n╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗\n ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝",
	})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "PARAMETRES",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: Changer de langue",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: Retour",
	})
	RED.NewLine(2)
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "1" {
		DisplayMainMenu()
	}
}
