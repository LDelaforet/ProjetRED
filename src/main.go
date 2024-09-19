package main

import (
	RED "RED/Internals"
	"fmt"
	"strconv"

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
	RED.ChestsInit()
	RED.EnnemisInit()
	*RED.CurrentMapIdPointer = 0
	*RED.CurrentTileIdPointer = 1
	*RED.CurrentMapPointer = RED.GetMapById(*RED.CurrentMapIdPointer)
	//DisplayShop()
	//DisplayMainMenu()
	MapNavigation()
}

func DisplayMainMenu() {

	RED.ClearScreen()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("gameName"),
	})

	RED.NewLine(3)
	RED.BoxStrings([]string{"0: " + RED.GetLineById("newGame"), "1: " + RED.GetLineById("loadGame"), "2: " + RED.GetLineById("training"), "3: " + RED.GetLineById("options"), "4: " + RED.GetLineById("quit")})

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
		RED.PlayerPointer.Name = "Player"
		RED.PlayerPointer.Class = 0
		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 20
		RED.PlayerPointer.Pv = 20
		RED.PlayerPointer.Defence = 3
		RED.PlayerPointer.Heal = 5
		RED.PlayerPointer.Money = 0
		RED.PlayerPointer.Level = 1

		goblin := RED.Enemy{
			Type:    "Training Dummy",
			PvMax:   100,
			Pv:      100,
			Damage:  1,
			Defence: 0,
		}
		fmt.Println(goblin)
		RED.BattleInit(goblin)
	} else if input == "3" {
		DisplayOptionMenu()
	} else if input == "4" {
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
	fmt.Printf("Choix: ")
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
		MapNavigation()

		goblin := RED.Enemy{
			Type:    "Goblin",
			PvMax:   15,
			Pv:      10,
			Damage:  3,
			Defence: 2,
		}
		fmt.Println(goblin)
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
	RED.BoxStrings([]string{RED.GetLineById("classSelection")})
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

		fmt.Print("Choisit la classe que tu veux analyser (0 - 3): ")
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
			RED.NewLine(1)
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: RED.GetLineById("class0Description"),
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
			RED.NewLine(1)
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: RED.GetLineById("class1Description"),
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
			RED.NewLine(1)
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: RED.GetLineById("class2Description"),
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
			RED.NewLine(1)
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: RED.GetLineById("class3Description"),
			})
		} else {
			ClassSelection()
		}
		RED.NewLine(2)
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

		*RED.DiscoveredPointer = append(*RED.DiscoveredPointer, *RED.CurrentTileIdPointer)
		CurrentTile := RED.GetMapTileById(*RED.CurrentTileIdPointer)

		ExecuteEvent()
		RED.GetInput()
		RED.ClearScreen()
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

		options[fmt.Sprintf("%d", optionCount)] = "nextmap"
		fmt.Printf("%d: Prochaine map\n", optionCount)
		optionCount++

		options[fmt.Sprintf("%d", optionCount)] = "playerProfile"
		fmt.Printf("%d: Profil joueur\n", optionCount)
		optionCount++

		/*options[fmt.Sprintf("%d", optionCount)] = "executeEvent"
		fmt.Printf("%d: Executer l'evenement\n", optionCount)
		optionCount++*/

		fmt.Print("Choix: ")
		input := RED.GetInput()
		if action, exists := options[input]; exists {
			switch action {
			case "8":
				*RED.CurrentTileIdPointer = CurrentTile.ToUpID
			case "5":
				*RED.CurrentTileIdPointer = CurrentTile.ToDownID
			case "4":
				*RED.CurrentTileIdPointer = CurrentTile.ToLeftID
			case "6":
				*RED.CurrentTileIdPointer = CurrentTile.ToRightID
			case "executeEvent":
				//ExecuteEvent()
				RED.GetInput()
			case "nextmap":
				*RED.CurrentMapIdPointer++
				*RED.CurrentTileIdPointer = 1
				*RED.DiscoveredPointer = []int{}
			case "playerProfile":
				DisplayInfo()
			}
		} else {
			fmt.Println("Option invalide.")
		}
	}
}

func ExecuteEvent() {
	CurrentTile := RED.GetMapTileById(*RED.CurrentTileIdPointer)
	eventID := CurrentTile.EventType

	switch eventID {
	case 0:
		// Start
	case 1:
		// End
	case 2:
		// Ennemi
		RED.BattleInit(RED.Ennemis[*RED.CurrentMapIdPointer])

	case 3:
		// Tuto
	case 4:
		// RED.Chests est égal a map[1:[{{8 Armure de plaques TODO 35} 1} {{11 Amulette de Tartempion TODO 100} 1} {{5 Epée en fer TODO 30} 1}] 2:[] 3:[] 4:[] 5:[]]
		fmt.Println("ID de la map", *RED.CurrentMapIdPointer)

		// Vérification de l'existence de la clé dans RED.Chests
		if items, exists := RED.Chests[*RED.CurrentMapIdPointer]; exists {
			if len(items) == 0 {
				fmt.Println("Aucun item dans le coffre pour cette carte.")
			} else {
				for _, item := range items {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, item)
				}
			}
		} else {
			fmt.Println("Aucun coffre trouvé pour cette carte.")
		}
	case 5:
		// Forge
	case 6:
		// Shop
		DisplayShop()
	case 7:
		// Doorlocked niveau 4
	case 8:
		// Story 1
	case 9:
		// Story 2
	case 10:
		// Story 3
	case 11:
		// Story 4
	case 12:
		// Story 5
	}

	// 0  Start
	// 1  End
	// 2  Ennemi
	// 3  Tuto
	// 4  Coffre
	// 5  Forge
	// 6  Shop
	// 7  Doorlocked niveau 4
	// 8  Story 1
	// 9  Story 2
	// 10  Story 3
	// 11 Story 4
	// 12 Story 5
}

func DisplayInfo() {
	RED.ClearScreen()
	RED.NewLine(2)
	RED.DisplayLine()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "PROFIL DE " + string(RED.PlayerPointer.Name),
	})
	RED.DisplayLine()
	RED.NewLine(1)
	RED.BoxStrings([]string{
		RED.GetLineById("damage") + ": 5",
		RED.GetLineById("hp") + ": 20",
		RED.GetLineById("hpMax") + ": 25",
		RED.GetLineById("defense") + ": 3",
		RED.GetLineById("heal") + ": 5",
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "0: " + RED.GetLineById("openInventory"),
	})
	RED.NewLine(1)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "1: " + RED.GetLineById("return"),
	})
	RED.NewLine(3)
	RED.DisplayLine()
	fmt.Printf("Choix: ")
	input := RED.GetInput()
	if input == "0" {
		RED.AccessInventory("playerProfile")
	} else if input == "1" {
		MapNavigation()
	} else {
		DisplayInfo()
	}
}

func DisplayShop() {
	RED.ClearScreen()
	RED.DisplayLine()
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "  " + RED.GetLineById("shopTitle"),
	})
	RED.DisplayLine()
	RED.NewLine(1)
	RED.BoxStrings([]string{"Que veux-tu acheter?"})
	RED.NewLine(2)

	if RED.PlayerPointer.Money >= RED.GetItemById(0).Price {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "0: " + RED.GetItemById(0).Name + ": " + strconv.Itoa(RED.GetItemById(0).Price) + "$\n  - " + RED.GetItemById(0).Description,
		})
	} else {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "0: " + RED.GetItemById(0).Name + ": " + strconv.Itoa(RED.GetItemById(0).Price) + "$\n  - " + RED.GetItemById(0).Description,
			FgColor:     color.FgRed,
		})
	}
	RED.NewLine(1)
	if RED.PlayerPointer.Money >= RED.GetItemById(1).Price {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "1: " + RED.GetItemById(1).Name + ": " + strconv.Itoa(RED.GetItemById(1).Price) + "$\n  - " + RED.GetItemById(1).Description,
		})
	} else {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "1: " + RED.GetItemById(1).Name + ": " + strconv.Itoa(RED.GetItemById(1).Price) + "$\n  - " + RED.GetItemById(1).Description,
			FgColor:     color.FgRed,
		})
	}
	RED.NewLine(1)
	if RED.PlayerPointer.Money >= RED.GetItemById(2).Price {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "2: " + RED.GetItemById(2).Name + ": " + strconv.Itoa(RED.GetItemById(2).Price) + "$\n  - " + RED.GetItemById(2).Description,
		})
	} else {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "2: " + RED.GetItemById(2).Name + ": " + strconv.Itoa(RED.GetItemById(2).Price) + "$\n  - " + RED.GetItemById(2).Description,
			FgColor:     color.FgRed,
		})
	}

	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "3: " + RED.GetLineById("quit"),
	})

	RED.NewLine(1)
	RED.DisplayLine()
	fmt.Print("Choix: ")
	input := RED.GetInput()
	if input == "0" {
		if RED.PlayerPointer.Money >= RED.GetItemById(0).Price {
			RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
				Item:     RED.GetItemById(0),
				Quantity: 1,
			})
		}
		DisplayShop()
	} else if input == "1" {
		if RED.PlayerPointer.Money >= RED.GetItemById(1).Price {
			RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
				Item:     RED.GetItemById(1),
				Quantity: 1,
			})
		}
		DisplayShop()
	} else if input == "2" {
		if RED.PlayerPointer.Money >= RED.GetItemById(2).Price {
			RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
				Item:     RED.GetItemById(2),
				Quantity: 1,
			})
		}
		DisplayShop()
	}
	if input == "3" {
		MapNavigation()
	} else {
		DisplayShop()
	}
}
