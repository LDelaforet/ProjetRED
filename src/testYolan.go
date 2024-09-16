package main

import (
	RED "RED/Internals"
	"fmt"

	"github.com/fatih/color"
)

// Remove the duplicate declaration of the main function
func main() {
	RED.PointersInit()
	RED.ReadMenuStrings()

	DisplayMainMenu()
}

func DisplayMainMenu() {

	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("gameName"),
	})

	RED.NewLine(3)
	RED.BoxStrings([]string{"0: " + RED.GetLineById("newGame"), "1: " + RED.GetLineById("loadGame"), "2: " + RED.GetLineById("options"), "3: " + RED.GetLineById("quit")})

	RED.NewLine(3)
	RED.DisplayLine()
	fmt.Print("Choix: ")
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
	RED.BoxStrings([]string{"0: " + RED.GetLineById("comfirm"), "1: " + RED.GetLineById("return")})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("newGameWarning"),
		FgColor:     color.FgRed,
	})
	RED.NewLine(2)
	RED.DisplayLine()
	fmt.Print("Choix: ")
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
	RED.BoxStrings([]string{"0: " + RED.GetLineById("changeLanguage"), "1: " + RED.GetLineById("return")})
	RED.NewLine(4)
	RED.DisplayLine()
	fmt.Print("Choix: ")
	input := RED.GetInput()
	if input == "0" {
		DisplayLanguageMenu()
	}
	if input == "1" {
		DisplayMainMenu()
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
			// J'en conçois
			TextToPrint: RED.GetLineById("languageTitle"),
		})
		RED.NewLine(1)
		RED.BoxStrings([]string{"0: " + RED.GetLineById("languageFrench"), "1: " + RED.GetLineById("languageEnglish")})
		RED.BoxStrings([]string{"2: " + RED.GetLineById("return")})
		RED.NewLine(1)
		RED.DisplayLine()
		fmt.Print("Choix: ")
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
	RED.PlayerPointer.Name = "Player"
	RED.ClearScreen()
	RED.DisplayLine()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("characterMakerTitle"),
	})
	RED.DisplayLine()
	RED.NewLine(1)
	RED.BoxStrings([]string{"0: " + RED.GetLineById("nameSelection")})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("currentName") + RED.PlayerPointer.Name,
	})
	RED.NewLine(1)
	RED.BoxStrings([]string{"1: " + RED.GetLineById("classSelection")})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("currentClass") + "None",
	})
	RED.NewLine(1)
	RED.BoxStrings([]string{"2: " + RED.GetLineById("return"), "3: " + RED.GetLineById("finish")})
	RED.NewLine(1)
	RED.DisplayLine()
	input := RED.GetInput()
	if input == "0" {

	}
	if input == "2" {
		DisplayMainMenu()
	}
	if input == "3" {
		DisplayGameUI()
	}
}

func DisplayGameUI() {
	RED.ClearScreen()

	goblin := RED.Enemy{
		Type:    "Goblin",
		PvMax:   15,
		Pv:      15,
		Damage:  10,
		Defence: 3,
	}

	RED.BattleInit(goblin)
}
