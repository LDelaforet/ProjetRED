package main

import (
	RED "RED/Internals"
	"fmt"
)

// TODO, currently main.go is a test for the Internals functions

func main() {
	testSlot := RED.InventorySlot{
		Item:     RED.ItemObject{},
		Quantity: 1,
	}

	testPerso := RED.Perso{
		Name:      "Test",
		Class:     1,
		Level:     1,
		Xp:        1,
		PvMax:     1,
		Pv:        1,
		Inventory: []RED.InventorySlot{testSlot},
	}
	RED.DisplayText(testPerso.Name)
	itemList := RED.ReadItemList()
	for _, item := range itemList {
		fmt.Printf("Item ID: %d, Name: %s, Description: %s\n", item.Id, item.Name, item.Description)
	}
}
