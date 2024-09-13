package main

import (
	RED "RED/Internals"

	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Initialise les pointeurs et lis les strings de menu.json
	goblin := RED.Enemy{
		Type:    "Goblin",
		PvMax:   100,
		Pv:      10,
		Damage:  20,
		Defence: 20,
	}

	RED.BattleInit(goblin)

	RED.PointersInit()
	RED.ReadMenuStrings()
	//DisplayMainMenu()
}

func DisplayMainMenu() {
	var height int
	_, height = RED.SizeTest()

	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("gameName"),
	})

	RED.NewLine(5)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: " + RED.GetLineById("newGame"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("loadGame"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "2: " + RED.GetLineById("options"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "3: " + RED.GetLineById("quit"),
	})
	fmt.Println(strings.Repeat("\n", height-19))
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
	var height int
	_, height = RED.SizeTest()
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
		TextToPrint: "0: " + RED.GetLineById("comfirm"), // ptn tu fais chier a avoir hardcodé le nombre, faudrait faire un compteur qui incrémente ou une connerie comme ca
	}) //TG Leo, cordialement Yolan
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("return"),
	})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("newGameWarning"),
		FgColor:     color.FgRed,
	})
	fmt.Println(strings.Repeat("\n", height-19))

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
		TextToPrint: "0: " + RED.GetLineById("changeLanguage"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("return"),
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
			TextToPrint: "0: " + RED.GetLineById("languageFrench"), // pareil faudrait incrémenter le digit au début, ca doit pas etre bien compliqué
		})
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "1: " + RED.GetLineById("languageEnglish"),
		})
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "2: " + RED.GetLineById("return"),
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
	RED.PlayerPointer.Name = "Player"
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
		TextToPrint: "0: " + RED.GetLineById("nameSelection"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("currentName") + RED.PlayerPointer.Name,
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("classSelection"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("currentClass") + "None",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "2: " + RED.GetLineById("return"),
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "3: " + RED.GetLineById("finish"),
	})
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
	RED.DisplayLine()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "| 0: " + RED.GetLineById("inventory") + " | 1: " + RED.GetLineById("map") + " | 2: " + RED.GetLineById("Paramètres"),
	})
	RED.DisplayLine()
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: " + RED.GetLineById("nameSelection"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("currentName") + RED.PlayerPointer.Name,
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("classSelection"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("currentClass") + "None",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "2: " + RED.GetLineById("return"),
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "3: " + RED.GetLineById("finish"),
	})
	RED.NewLine(1)
	RED.DisplayLine()
}
