package main

import (
	RED "RED/Internals"
	"fmt"

	"github.com/fatih/color"
)

// TODO, currently main.go is a test for the Internals functions

func main() {
	RED.PointersInit()

	goblin := RED.Enemy{
		Type:    "Goblin",
		PvMax:   15,
		Pv:      15,
		Damage:  10,
		Defence: 3,
	}

	RED.BattleInit(goblin)
	//return

	testPerso := RED.PlayerPointer
	// itemList := RED.ItemListPointer
	isGameInFrenchPointer := RED.IsGameInFrenchPointer
	// i18nLinePointer := RED.MenuLinesPointer

	testPerso.Name = "Test"

	*isGameInFrenchPointer = true

	RED.ReadItemList()
	RED.ReadMenuStrings()

	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("newGame"),
		FgColor:     color.FgRed,
	})

	*isGameInFrenchPointer = false

	RED.ReadItemList()
	RED.ReadMenuStrings()

	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: RED.GetLineById("newGame"),
		FgColor:     color.FgRed,
		BgColor:     color.BgBlue,
	})

}

// poinTest est une fonction qui sert a tester l'utilité des pointeurs dans des cas précis
func poinTest() {
	RED.PointersInit()

	// Fous les pointeurs dans des var locales
	playerPointer := RED.PlayerPointer
	itemListPointer := RED.ItemListPointer
	isGameInFrenchPointer := RED.IsGameInFrenchPointer
	i18nLinePointer := RED.MenuLinesPointer

	// Lis la liste des items
	RED.ReadItemList()
	RED.ReadMenuStrings()

	fmt.Println(playerPointer.Name)

	// Affiche la liste des items
	for _, item := range *itemListPointer {
		fmt.Printf("Item ID: %d, Name: %s, Description: %s\n", item.Id, item.Name, item.Description)
	}

	fmt.Println(playerPointer.Name)

	for _, item := range *i18nLinePointer {
		fmt.Printf("Item ID: %s, Name: %s\n", item.Id, item.Line)
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
	for _, item := range *i18nLinePointer {
		fmt.Printf("Item ID: %s, Name: %s\n", item.Id, item.Id, item.Line)
	}

	fmt.Println(*isGameInFrenchPointer)
}

// Série de tests basics pour voir si rien n'est pété
func basicTests() {

	// poinTest()

	RED.PointersInit()

	testPerso := RED.PlayerPointer
	itemList := RED.ItemListPointer
	isGameInFrenchPointer := RED.IsGameInFrenchPointer
	i18nLinePointer := RED.MenuLinesPointer

	testPerso.Name = "Test"

	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: testPerso.Name,
		FgColor:     color.FgRed,
		Underline:   true,
	})

	RED.DisplayText(RED.DisplayTextOptions{
		TextToPrint: testPerso.Name,
		FgColor:     color.FgRed,
		Underline:   false,
	})

	*isGameInFrenchPointer = false

	RED.ReadItemList()
	RED.ReadMenuStrings()

	for _, item := range *itemList {
		fmt.Printf("Item ID: %d, Name: %s, Description: %s\n", item.Id, item.Name, item.Description)
	}

	for _, item := range *i18nLinePointer {
		fmt.Printf("Id: %s, Line: %s\n", item.Id, item.Line)
	}

}
