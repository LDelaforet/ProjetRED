package main

import (
	RED "RED/Internals"

	"github.com/fatih/color"
)

func main() {
	// Initialise les pointeurs et lis les strings de menu.json
	RED.PointersInit()
	RED.ReadMenuStrings()

	DisplayMainMenu()
}

func DisplayMainMenu() {
	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("gameName"),
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
	RED.NewLine(8)
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
		TextToPrint: RED.GetLineById("gameName"),
	})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("newGameQuestion"),
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: Comfirmer", // ptn tu fais chier a avoir hardcodé le nombre, faudrait faire un compteur qui incrémente ou une connerie comme ca
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: Retour",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("newGameWarning"),
		FgColor:     color.FgRed,
	})
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "0" {
		DisplayCharacterCustomizationPanel()
	}
	if input == "1" {
		DisplayMainMenu()
	}
}

func DisplayOptionMenu() {
	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("gameName"),
	})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{ // Pourquoi il est en majuscules?
		TextToPrint: RED.GetLineById("settingsTitle"),
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: Selection de la langue",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: Retour",
	})
	RED.NewLine(10)
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "0" {
		DisplayLanguageMenu()
		if input == "1" {
			DisplayMainMenu()
		}
	}
}

func DisplayLanguageMenu() {
	for {
		RED.ClearScreen()
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("gameName"),
		})
		RED.NewLine(2)
		RED.DisplayText(RED.DisplayTextOptions{ // Pourquoi il est en majuscules?
			TextToPrint: RED.GetLineById("languageTitle"),
		})
		RED.NewLine(1)
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "0: Français", // pareil faudrait incrémenter le digit au début, ca doit pas etre bien compliqué
		})
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "1: Anglais",
		})
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "2: Retour",
		})
		RED.NewLine(10)
		RED.DisplayLine()
		input := RED.GetInput()
		if input == "0" {
			*RED.IsGameInFrenchPointer = true
			RED.ReadItemList()
			RED.ReadMenuStrings()
		}
		if input == "1" {
			*RED.IsGameInFrenchPointer = false
			RED.ReadItemList()
			RED.ReadMenuStrings()
		}
		if input == "2" {
			break
		}
	}
	DisplayMainMenu()
}

func DisplayCharacterCustomizationPanel() {
	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("gameName"),
	})
	RED.DisplayLine()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("characterMakerTitle"),
	})
	RED.DisplayLine()
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: Choisit un nom",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: Selecteur de classes",
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "2: Retour",
	})
	RED.NewLine(1)
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "0" {
		DisplayMainMenu()
	}
	if input == "1" {
		DisplayMainMenu()
	}
	if input == "2" {
		DisplayMainMenu()
	}
}
