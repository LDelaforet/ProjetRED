package main

import (
	RED "RED/Internals"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var PastEvent []int = []int{}

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
	RED.ReadCraftingRecipes()
	RED.ChestsInit()
	RED.EnnemisInit()
	*RED.CurrentMapIdPointer = 0
	*RED.CurrentTileIdPointer = 1
	*RED.CurrentMapPointer = RED.GetMapById(*RED.CurrentMapIdPointer)
	//DisplayShop()
	DisplayMainMenu()
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
	fmt.Print(RED.GetLineById("choice") + ": ")
	input := RED.GetInput()

	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: input,
	})

	if input == "0" {
		DisplayNewGameMenu()
	} else if input == "1" {
		RED.LoadFromJSON()
		MapNavigation("mainMenu")
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

		RED.AddItemToInventory(0, 1)
		RED.AddItemToInventory(1, 1)
		RED.AddItemToInventory(2, 1)

		RED.BattleInit(RED.Ennemis[5])
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
	RED.NewLine(3)
	RED.DisplayLine()
	fmt.Print(RED.GetLineById("choice") + ": ")
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
	RED.NewLine(5)
	RED.DisplayLine()
	fmt.Print(RED.GetLineById("choice") + ": ")
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
		RED.NewLine(2)
		RED.DisplayLine()
		fmt.Print(RED.GetLineById("choice") + ": ")
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
	RED.NewLine(2)
	RED.DisplayLine()
	fmt.Print(RED.GetLineById("choice") + ": ")
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
		if RED.PlayerPointer.Class == 3 {
			RED.AddItemToInventory(0, 3)
		}
		MapNavigation("characterCustomization")
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
	fmt.Print(RED.GetLineById("choice") + ": ")
	input := RED.GetInput()
	if input == "0" {
		RED.PlayerPointer.Class = 0

		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 25
		RED.PlayerPointer.Pv = 20
		RED.PlayerPointer.Defence = 2
		RED.PlayerPointer.Heal = 3
		RED.PlayerPointer.InventorySize = 10

		DisplayCharacterCustomizationPanel()
	} else if input == "1" {
		RED.PlayerPointer.Class = 1

		RED.PlayerPointer.Damage = 10
		RED.PlayerPointer.PvMax = 20
		RED.PlayerPointer.Pv = 15
		RED.PlayerPointer.Defence = 3
		RED.PlayerPointer.Heal = 2
		RED.PlayerPointer.InventorySize = 5

		DisplayCharacterCustomizationPanel()
	} else if input == "2" {
		RED.PlayerPointer.Class = 2

		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 25
		RED.PlayerPointer.Pv = 25
		RED.PlayerPointer.Defence = 7
		RED.PlayerPointer.Heal = 1
		RED.PlayerPointer.InventorySize = 15

		DisplayCharacterCustomizationPanel()
	} else if input == "3" {
		RED.PlayerPointer.Class = 3

		RED.PlayerPointer.Damage = 5
		RED.PlayerPointer.PvMax = 30
		RED.PlayerPointer.Pv = 20
		RED.PlayerPointer.Defence = 1
		RED.PlayerPointer.Heal = 5
		RED.PlayerPointer.InventorySize = 10

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
				RED.GetLineById("defense") + ": 2",
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

func MapNavigation(previousMenu string) {
	// Script principal: Navigation dans la carte
	for {
		*RED.DiscoveredPointer = append(*RED.DiscoveredPointer, *RED.CurrentTileIdPointer)
		CurrentTile := RED.GetMapTileById(*RED.CurrentTileIdPointer)
		ExecuteEvent()

		RED.ClearScreen()
		RED.DisplayLine()
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: RED.GetLineById("map"),
		})
		RED.DisplayLine()
		RED.NewLine(1)
		options := make(map[string]string)
		optionCount := 1

		RED.LiteMapDisplay()
		RED.DisplayLine()
		if CurrentTile.ToRightID != 0 {
			options[fmt.Sprintf("6")] = "6"
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: "6: " + RED.GetLineById("goRight"),
			})
		}
		if CurrentTile.ToDownID != 0 {
			options[fmt.Sprintf("5")] = "5"
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: "5: " + RED.GetLineById("goDown"),
			})
		}
		if CurrentTile.ToLeftID != 0 {
			options[fmt.Sprintf("4")] = "4"
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: "4: " + RED.GetLineById("goLeft"),
			})
		}
		if CurrentTile.ToUpID != 0 {
			options[fmt.Sprintf("8")] = "8"
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: "8: " + RED.GetLineById("goUp"),
			})
		}
		RED.DisplayLine()
		RED.NewLine(1)
		if optionCount == 8 || optionCount == 4 || optionCount == 5 || optionCount == 6 {
			optionCount++
		}

		/*options[fmt.Sprintf("%d", optionCount)] = "nextmap"
		fmt.Printf("%d: Prochaine map\n", optionCount)
		optionCount++*/

		options[fmt.Sprintf("%d", optionCount)] = "playerProfile"
		fmt.Printf("%d: "+RED.GetLineById("playerProfile")+"\n", optionCount)
		optionCount++

		options[fmt.Sprintf("%d", optionCount)] = "saveGame"
		fmt.Printf("%d: "+RED.GetLineById("saveGame")+"\n", optionCount)
		optionCount++
		/*options[fmt.Sprintf("%d", optionCount)] = "executeEvent"
		fmt.Printf("%d: Executer l'evenement\n", optionCount)
		optionCount++*/
		RED.NewLine(1)
		RED.DisplayLine()
		fmt.Print(RED.GetLineById("choice") + ": ")
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
				//RED.GetInput()
			case "nextmap":
				*RED.CurrentMapIdPointer++
				*RED.CurrentTileIdPointer = 1
				*RED.DiscoveredPointer = []int{}
			case "playerProfile":
				DisplayInfo()
			case "saveGame":
				RED.SaveToJSON()
				RED.ClearScreen()
			}
		} else {
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: RED.GetLineById("invalidOption"),
				FgColor:     color.FgRed,
			})
		}
	}
}

func ExecuteEvent() {
	//Check si CurrentTile est dans pastEvent
	if RED.Contains(PastEvent, *RED.CurrentTileIdPointer) {
		//fmt.Println("Tu as déjà visité cette case.")
		return
	} else {
		//fmt.Println("Tu n'as pas encore visité cette case.")
		PastEvent = append(PastEvent, *RED.CurrentTileIdPointer)
	}
	//fmt.Println("PastEvent", PastEvent)
	//RED.GetInput()
	CurrentTile := RED.GetMapTileById(*RED.CurrentTileIdPointer)
	eventID := CurrentTile.EventType
	//fmt.Println("eventID", eventID)
	switch eventID {
	case 0:
		// Start
	case 1:
		// End
		if *RED.CurrentMapIdPointer < 4 {
			*RED.CurrentMapIdPointer++
			*RED.CurrentTileIdPointer = 1
			*RED.DiscoveredPointer = []int{}
			PastEvent = []int{}
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: "Tu as terminé la carte, tu passes à la suivante.",
			})
			RED.GetInput()
			MapNavigation("end")
		} else {
			RED.DisplayText(RED.DisplayTextOptions{
				TextToPrint: "Tu as terminé le jeu.",
			})
			RED.GetInput()
			DisplayMainMenu()
		}
	case 2:
		// Ennemi
		RED.BattleInit(RED.Ennemis[*RED.CurrentMapIdPointer])

	case 3:
		// Tuto
	case 4:
		// RED.Chests est égal a map[1:[{{8 Armure de plaques TODO 35} 1} {{11 Amulette de Tartempion TODO 100} 1} {{5 Epée en fer TODO 30} 1}] 2:[] 3:[] 4:[] 5:[]]
		//fmt.Println("ID de la map", *RED.CurrentMapIdPointer)

		// Vérification de l'existence de la clé dans RED.Chests
		if items, exists := RED.Chests[*RED.CurrentMapIdPointer]; exists {
			if len(items) == 0 {
				fmt.Println("Aucun item dans le coffre pour cette carte.")
			} else {
				for _, item := range items {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, item)
					RED.NewLine(1)
					RED.DisplayText(RED.DisplayTextOptions{
						TextToPrint: "Tu trouves " + item.Item.Name + " dans le coffre.",
					})
				}
			}
		} else {
			fmt.Println("Aucun coffre trouvé pour cette carte.")
		}
		RED.GetInput()
	case 5:
		// Forge
		RED.CraftingMenu()
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
		RED.GetLineById("damage") + ": " + strconv.Itoa(int(RED.PlayerPointer.Damage)),
		RED.GetLineById("hp") + ": " + strconv.Itoa(int(RED.PlayerPointer.Pv)),
		RED.GetLineById("hpMax") + ": " + strconv.Itoa(int(RED.PlayerPointer.PvMax)),
		RED.GetLineById("defense") + ": " + strconv.Itoa(int(RED.PlayerPointer.Defence)),
		RED.GetLineById("heal") + ": " + strconv.Itoa(int(RED.PlayerPointer.Heal)),
		RED.GetLineById("money") + ": " + strconv.Itoa(int(RED.PlayerPointer.Money)) + "$",
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
	fmt.Print(RED.GetLineById("choice") + ": ")
	input := RED.GetInput()
	if input == "0" {
		RED.AccessInventory("playerProfile")
	} else if input == "1" {
		MapNavigation("playerProfile")
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
	RED.BoxStrings([]string{"Que veux-tu acheter? | Argent: " + strconv.Itoa(RED.PlayerPointer.Money) + "$"})
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
	RED.NewLine(1)
	if RED.PlayerPointer.Money >= RED.GetItemById(10).Price {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "3: " + RED.GetItemById(10).Name + ": " + strconv.Itoa(RED.GetItemById(10).Price) + "$\n  - " + RED.GetItemById(10).Description,
		})
	} else {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "3: " + RED.GetItemById(10).Name + ": " + strconv.Itoa(RED.GetItemById(10).Price) + "$\n  - " + RED.GetItemById(10).Description,
			FgColor:     color.FgRed,
		})
	}
	RED.NewLine(1)
	if RED.PlayerPointer.Money >= RED.GetItemById(11).Price {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "4: " + RED.GetItemById(11).Name + ": " + strconv.Itoa(RED.GetItemById(11).Price) + "$\n  - " + RED.GetItemById(11).Description,
		})
	} else {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "4: " + RED.GetItemById(11).Name + ": " + strconv.Itoa(RED.GetItemById(11).Price) + "$\n  - " + RED.GetItemById(11).Description,
			FgColor:     color.FgRed,
		})
	}
	RED.NewLine(1)
	if RED.PlayerPointer.Money >= RED.GetItemById(99).Price {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "5: " + RED.GetItemById(99).Name + ": " + strconv.Itoa(RED.GetItemById(99).Price) + "$\n  - " + RED.GetItemById(99).Description,
		})
	} else {
		RED.DisplayText(RED.DisplayTextOptions{
			TextToPrint: "5: " + RED.GetItemById(99).Name + ": " + strconv.Itoa(RED.GetItemById(99).Price) + "$\n  - " + RED.GetItemById(99).Description,
			FgColor:     color.FgRed,
		})
	}
	RED.NewLine(2)
	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: "6: " + RED.GetLineById("quit"),
	})

	RED.NewLine(1)
	RED.DisplayLine()
	fmt.Print(RED.GetLineById("choice") + ": ")
	input := RED.GetInput()
	if len(RED.PlayerPointer.Inventory) <= 10 {
		if input == "0" {
			if RED.PlayerPointer.Money >= RED.GetItemById(0).Price && RED.PlayerPointer.Inventory[0].Quantity < byte(RED.PlayerPointer.InventorySize) {
				//Check si l'item existe déjà dans l'inventaire
				itemExists := false
				for _, slot := range RED.PlayerPointer.Inventory {
					if slot.Item == RED.GetItemById(0) {
						itemExists = true
						break
					}
				}
				if itemExists {
					RED.PlayerPointer.Inventory[0].Quantity++
				} else {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
						Item:     RED.GetItemById(0),
						Quantity: 1,
					})
				}
				RED.PlayerPointer.Money -= RED.GetItemById(0).Price
			}
			DisplayShop()
		} else if input == "1" {
			if RED.PlayerPointer.Money >= RED.GetItemById(1).Price && RED.PlayerPointer.Inventory[1].Quantity < byte(RED.PlayerPointer.InventorySize) {
				//Check si l'item existe déjà dans l'inventaire
				itemExists := false
				for _, slot := range RED.PlayerPointer.Inventory {
					if slot.Item == RED.GetItemById(1) {
						itemExists = true
						break
					}
				}
				if itemExists {
					RED.PlayerPointer.Inventory[1].Quantity++
				} else {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
						Item:     RED.GetItemById(1),
						Quantity: 1,
					})
				}
				RED.PlayerPointer.Money -= RED.GetItemById(1).Price
			}
			DisplayShop()
		} else if input == "2" {
			if RED.PlayerPointer.Money >= RED.GetItemById(2).Price && RED.PlayerPointer.Inventory[2].Quantity < byte(RED.PlayerPointer.InventorySize) {
				//Check si l'item existe déjà dans l'inventaire
				itemExists := false
				for _, slot := range RED.PlayerPointer.Inventory {
					if slot.Item == RED.GetItemById(2) {
						itemExists = true
						break
					}
				}
				if itemExists {
					RED.PlayerPointer.Inventory[2].Quantity++
				} else {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
						Item:     RED.GetItemById(2),
						Quantity: 1,
					})
				}
				RED.PlayerPointer.Money -= RED.GetItemById(2).Price
			}
			DisplayShop()
		} else if input == "3" {
			if RED.PlayerPointer.Money >= RED.GetItemById(3).Price && RED.PlayerPointer.Inventory[3].Quantity < byte(RED.PlayerPointer.InventorySize) {
				//Check si l'item existe déjà dans l'inventaire
				itemExists := false
				for _, slot := range RED.PlayerPointer.Inventory {
					if slot.Item == RED.GetItemById(3) {
						itemExists = true
						break
					}
				}
				if itemExists {
					RED.PlayerPointer.Inventory[3].Quantity++
				} else {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
						Item:     RED.GetItemById(3),
						Quantity: 1,
					})
				}
				RED.PlayerPointer.Money -= RED.GetItemById(3).Price
			}
			DisplayShop()
		} else if input == "4" {
			if RED.PlayerPointer.Money >= RED.GetItemById(4).Price && RED.PlayerPointer.Inventory[4].Quantity < byte(RED.PlayerPointer.InventorySize) {
				//Check si l'item existe déjà dans l'inventaire
				itemExists := false
				for _, slot := range RED.PlayerPointer.Inventory {
					if slot.Item == RED.GetItemById(4) {
						itemExists = true
						break
					}
				}
				if itemExists {
					RED.PlayerPointer.Inventory[4].Quantity++
				} else {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
						Item:     RED.GetItemById(4),
						Quantity: 1,
					})
				}
				RED.PlayerPointer.Money -= RED.GetItemById(4).Price
			}
			DisplayShop()
		} else if input == "4" {
			if RED.PlayerPointer.Money >= RED.GetItemById(99).Price && RED.PlayerPointer.Inventory[99].Quantity < byte(RED.PlayerPointer.InventorySize) {
				//Check si l'item existe déjà dans l'inventaire
				itemExists := false
				for _, slot := range RED.PlayerPointer.Inventory {
					if slot.Item == RED.GetItemById(99) {
						itemExists = true
						break
					}
				}
				if itemExists {
					RED.PlayerPointer.Inventory[99].Quantity++
				} else {
					RED.PlayerPointer.Inventory = append(RED.PlayerPointer.Inventory, RED.InventorySlot{
						Item:     RED.GetItemById(99),
						Quantity: 1,
					})
				}
				RED.PlayerPointer.Money -= RED.GetItemById(99).Price
			}
			DisplayShop()
		}
	}
	if input == "6" {
		MapNavigation("shop")
	} else {
		DisplayShop()
	}
}
