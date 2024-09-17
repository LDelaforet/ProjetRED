package main

import (
	RED "RED/Internals"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	TestSavePointers()
}

func TestSavePointers() {
	// Run the SavePointers function
	RED.SavePointers()

	// Check if the file save.json exists
	if _, err := os.Stat("save.json"); os.IsNotExist(err) {
		fmt.Printf("Expected file save.json to be created\n")
		return
	}

	// Read the file content
	data, err := ioutil.ReadFile("save.json")
	if err != nil {
		fmt.Printf("Error reading save.json: %v\n", err)
		return
	}

	// Unmarshal the JSON data
	var pointers RED.Pointers
	err = json.Unmarshal(data, &pointers)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Check the values
	if *pointers.PlayerPointer != "Player1" {
		fmt.Printf("Expected PlayerPointer to be 'Player1', got '%s'\n", *pointers.PlayerPointer)
	}
	if *pointers.ItemListPointer != "Sword, Shield" {
		fmt.Printf("Expected ItemListPointer to be 'Sword, Shield', got '%s'\n", *pointers.ItemListPointer)
	}
}
