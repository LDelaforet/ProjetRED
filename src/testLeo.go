package main

import (
	RED "RED/Internals"
	"fmt"
)

// TODO, currently main.go is a test for the Internals functions

func main() {
	RED.PointersInit()

	testPerso := RED.PlayerPointer
	itemList := RED.ItemListPointer
	isGameInFrenchPointer := RED.IsGameInFrenchPointer

	testPerso.Name = "Test"

	RED.DisplayText(testPerso.Name)

	*isGameInFrenchPointer = true

	RED.ReadItemList()

	for _, item := range *itemList {
		fmt.Printf("Item ID: %d, Name: %s, Description: %s\n", item.Id, item.Name, item.Description)
	}
}

// poinTest est une fonction qui sert a tester l'utilité des pointeurs dans des cas précis
func poinTest() {
	RED.PointersInit()

	// Fous les pointeurs dans des var locales
	playerPointer := RED.PlayerPointer
	itemListPointer := RED.ItemListPointer
	isGameInFrenchPointer := RED.IsGameInFrenchPointer

	// Lis la liste des items
	RED.ReadItemList()

	fmt.Println(playerPointer.Name)

	// Affiche la liste des items
	for _, item := range *itemListPointer {
		fmt.Printf("Item ID: %d, Name: %s, Description: %s\n", item.Id, item.Name, item.Description)
	}

	fmt.Println(*isGameInFrenchPointer)

	fmt.Println("EDIT TIME !")

	RED.TestPlayerPointer(playerPointer)
	RED.TestItemListPointer(itemListPointer)
	RED.TestIsGameInFrenchPointer(isGameInFrenchPointer)

	fmt.Println(playerPointer.Name)

	// Affiche la liste des items
	for _, item := range *itemListPointer {
		fmt.Printf("Item ID: %d, Name: %s, Description: %s\n", item.Id, item.Name, item.Description)
	}

	fmt.Println(*isGameInFrenchPointer)
}
