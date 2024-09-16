package main

import (
	RED "RED/Internals"
	"fmt"
)

// TODO: Une fonction TitleScreen (force a toi Yolan)

func main() {
	TitleScreen()

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

func TitleScreen() {
	RED.PointersInit()
	RED.ReadItemList()
	RED.ReadMapLists()
	RED.ReadMenuStrings()
	// Emule le menu principal
	RED.PlayerPointer.Name = "Test"
	RED.PlayerPointer.Pv = 100
	RED.PlayerPointer.PvMax = 100
	RED.PlayerPointer.Damage = 10
	RED.PlayerPointer.Defence = 10
	RED.PlayerPointer.Heal = 10
	RED.PlayerPointer.Money = 0
	*RED.CurrentMapIdPointer = 0
	*RED.CurrentTileIdPointer = 1
	*RED.CurrentMapPointer = RED.GetMapById(*RED.CurrentMapIdPointer)
}
