package main

import (
	RED "RED/Internals"
	"fmt"

	"github.com/fatih/color"
)

// TODO: Une fonction TitleScreen (force a toi Yolan)
// TODO: Un systeme de shop (force a toi Leo)

func main() {
	TitleScreen()
}

func TitleScreen() {
	RED.PointersInit()
	RED.ReadItemList()
	RED.ReadMapLists()
	RED.ReadMenuStrings()
	*RED.CurrentMapIdPointer = 0
	*RED.CurrentTileIdPointer = 1
	*RED.CurrentMapPointer = RED.GetMapById(*RED.CurrentMapIdPointer)
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

	if input == "0" {
		DisplayNewGameMenu()
	} else if input == "2" {
		DisplayOptionMenu()
	} else if input == "3" {
		RED.ClearScreen()
	} else {
		DisplayMainMenu()
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
		RED.PlayerPointer.Name = "Player"
		RED.PlayerPointer.Class = 0
		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 20
		RED.PlayerPointer.Pv = 20
		RED.PlayerPointer.Defence = 3
		RED.PlayerPointer.Heal = 5
		RED.PlayerPointer.Money = 0
		RED.PlayerPointer.Level = 1
		DisplayCharacterCustomizationPanel()
	} else if input == "1" {
		DisplayMainMenu()
	} else {
		DisplayNewGameMenu()
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
	} else if input == "1" {
		DisplayMainMenu()
	} else {
		DisplayOptionMenu()
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
		} else if input == "1" {
			*RED.IsGameInFrenchPointer = false
			RED.ReadItemList()
			RED.ReadMenuStrings()
		} else if input == "2" {
			break
		} else {
			DisplayLanguageMenu()
		}
	}
	DisplayMainMenu()
}

func DisplayCharacterCustomizationPanel() {
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

	if RED.PlayerPointer.Class == 0 {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("currentClass") + RED.GetLineById("class0"),
		})
	} else if RED.PlayerPointer.Class == 1 {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("currentClass") + RED.GetLineById("class1"),
		})
	} else if RED.PlayerPointer.Class == 2 {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("currentClass") + RED.GetLineById("class2"),
		})
	} else if RED.PlayerPointer.Class == 3 {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("currentClass") + RED.GetLineById("class3"),
		})
	}

	RED.NewLine(1)
	RED.BoxStrings([]string{"2: " + RED.GetLineById("return"), "3: " + RED.GetLineById("finish")})
	RED.NewLine(1)
	RED.DisplayLine()
	fmt.Print("Choix: ")
	input := RED.GetInput()
	if input == "0" {
		fmt.Print("Entre un nom: ")
		name := RED.GetInput()
		RED.PlayerPointer.Name = name
		DisplayCharacterCustomizationPanel()
	} else if input == "1" {
		ClassSelection()
	} else if input == "2" {
		DisplayMainMenu()
	} else if input == "3" {
		//MapNavigation()

		goblin := RED.Enemy{
			Type:    "Goblin",
			PvMax:   15,
			Pv:      10,
			Damage:  3,
			Defence: 2,
		}

		RED.BattleInit(goblin)
	} else {
		DisplayCharacterCustomizationPanel()
	}
}

func ClassSelection() {
	RED.ClearScreen()
	RED.DisplayLine()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("characterMakerTitle"),
	})
	RED.DisplayLine()
	RED.NewLine(1)
	RED.BoxStrings([]string{"0: " + RED.GetLineById("classSelection")})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: " + RED.GetLineById("class0"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("class1"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "2: " + RED.GetLineById("class2"),
	})
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "3: " + RED.GetLineById("class3"),
	})
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "4: " + RED.GetLineById("classInfo"),
	})

	RED.NewLine(3)
	RED.DisplayLine()
	fmt.Print("Choix: ")
	input := RED.GetInput()
	if input == "0" {
		RED.PlayerPointer.Class = 0

		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 25
		RED.PlayerPointer.Pv = 20
		RED.PlayerPointer.Defence = 3
		RED.PlayerPointer.Heal = 3

		DisplayCharacterCustomizationPanel()
	} else if input == "1" {
		RED.PlayerPointer.Class = 1

		RED.PlayerPointer.Damage = 10
		RED.PlayerPointer.PvMax = 20
		RED.PlayerPointer.Pv = 15
		RED.PlayerPointer.Defence = 3
		RED.PlayerPointer.Heal = 2

		DisplayCharacterCustomizationPanel()
	} else if input == "2" {
		RED.PlayerPointer.Class = 2

		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 25
		RED.PlayerPointer.Pv = 25
		RED.PlayerPointer.Defence = 7
		RED.PlayerPointer.Heal = 1

		DisplayCharacterCustomizationPanel()
	} else if input == "3" {
		RED.PlayerPointer.Class = 3

		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 30
		RED.PlayerPointer.Pv = 20
		RED.PlayerPointer.Defence = 1
		RED.PlayerPointer.Heal = 5

		DisplayCharacterCustomizationPanel()
	} else if input == "4" {
		// Afficher les infos de la classe

		fmt.Print("Choisit la classe que tu veux analyser: ")
		classInfo := RED.GetInput()
		RED.ClearScreen()
		RED.DisplayLine()
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("characterMakerTitle"),
		})
		RED.DisplayLine()
		if classInfo == "0" {
			RED.NewLine(1)
			RED.BoxStrings([]string{RED.GetLineById("class0")})
			RED.BoxStrings([]string{
				RED.GetLineById("damage") + ": 5",
				RED.GetLineById("hp") + ": 20",
				RED.GetLineById("hpMax") + ": 25",
				RED.GetLineById("defense") + ": 3",
				RED.GetLineById("heal") + ": 5",
			})
		} else if classInfo == "1" {
			RED.NewLine(1)
			RED.BoxStrings([]string{RED.GetLineById("class1")})
			RED.BoxStrings([]string{
				RED.GetLineById("damage") + ": 10",
				RED.GetLineById("hp") + ": 15",
				RED.GetLineById("hpMax") + ": 20",
				RED.GetLineById("defense") + ": 3",
				RED.GetLineById("heal") + ": 7",
			})
		} else if classInfo == "2" {
			RED.NewLine(1)
			RED.BoxStrings([]string{RED.GetLineById("class2")})
			RED.BoxStrings([]string{
				RED.GetLineById("damage") + ": 5",
				RED.GetLineById("hp") + ": 25",
				RED.GetLineById("hpMax") + ": 25",
				RED.GetLineById("defense") + ": 7",
				RED.GetLineById("heal") + ": 2",
			})
		} else if classInfo == "3" {
			RED.NewLine(1)
			RED.BoxStrings([]string{RED.GetLineById("class3")})
			RED.BoxStrings([]string{
				RED.GetLineById("damage") + ": 5",
				RED.GetLineById("hp") + ": 20",
				RED.GetLineById("hpMax") + ": 50",
				RED.GetLineById("defense") + ": 1",
				RED.GetLineById("heal") + ": 10",
			})
		} else {
			ClassSelection()
		}
		RED.NewLine(5)
		RED.DisplayLine()
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "Appuyez sur une touche pour continuer...",
		})
		_ = RED.GetInput()
		ClassSelection()

	} else {
		ClassSelection()
	}
}

func MapNavigation() {
	// Script principal: Navigation dans la carte
	for {
		RED.ClearScreen()
		*RED.DiscoveredPointer = append(*RED.DiscoveredPointer, *RED.CurrentTileIdPointer)
		CurrentTile := RED.GetMapTileById(*RED.CurrentTileIdPointer)

		options := make(map[string]string)
		optionCount := 1

		RED.LiteMapDisplay()

		if CurrentTile.ToRightID != 0 {
			options[fmt.Sprintf("6")] = "6"
			fmt.Printf("6: Aller à droite\n")
		}
		if CurrentTile.ToDownID != 0 {
			options[fmt.Sprintf("5")] = "5"
			fmt.Printf("5: Aller en bas\n")
		}
		if CurrentTile.ToLeftID != 0 {
			options[fmt.Sprintf("4")] = "4"
			fmt.Printf("4: Aller à gauche\n")
		}
		if CurrentTile.ToUpID != 0 {
			options[fmt.Sprintf("8")] = "8"
			fmt.Printf("8: Aller en haut\n")
		}

		if optionCount == 8 || optionCount == 4 || optionCount == 5 || optionCount == 6 {
			optionCount++
		}
		options[fmt.Sprintf("%d", optionCount)] = "map"
		fmt.Printf("%d: Carte\n", optionCount)
		optionCount++

		if optionCount == 8 || optionCount == 4 || optionCount == 5 || optionCount == 6 {
			optionCount++
		}
		options[fmt.Sprintf("%d", optionCount)] = "nextmap"
		fmt.Printf("%d: Prochaine map\n", optionCount)
		optionCount++

		fmt.Print("Choix: ")
		input := RED.GetInput()
		if action, exists := options[input]; exists {
			fmt.Println("Vous avez choisi:", action)
			switch action {
			case "8":
				*RED.CurrentTileIdPointer = CurrentTile.ToUpID
			case "5":
				*RED.CurrentTileIdPointer = CurrentTile.ToDownID
			case "4":
				*RED.CurrentTileIdPointer = CurrentTile.ToLeftID
			case "6":
				*RED.CurrentTileIdPointer = CurrentTile.ToRightID
			case "map":
				RED.MapDisplay()
			case "nextmap":
				*RED.CurrentMapIdPointer++
				*RED.CurrentTileIdPointer = 1
				*RED.DiscoveredPointer = []int{}
			}
		} else {
			fmt.Println("Option invalide.")
		}
	}
}
