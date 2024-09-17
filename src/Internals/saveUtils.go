package RED

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct to hold the pointers
type Pointers struct {
	PlayerPointer         *string `json:"player_pointer"`
	ItemListPointer       *string `json:"item_list_pointer"`
	IsGameInFrenchPointer *bool   `json:"is_game_in_french_pointer"`
	MenuLinesPointer      *string `json:"menu_lines_pointer"`
	MapListPointer        *string `json:"map_list_pointer"`
	CurrentEnemyPointer   *string `json:"current_enemy_pointer"`
	DiscoveredPointer     *string `json:"discovered_pointer"`
	CurrentMapPointer     *string `json:"current_map_pointer"`
	CurrentMapIdPointer   *int    `json:"current_map_id_pointer"`
	CurrentTileIdPointer  *int    `json:"current_tile_id_pointer"`
}

func SavePointers() {
	// Initialize pointers with some values
	player := "Player1"
	itemList := "Sword, Shield"
	isGameInFrench := true
	menuLines := "Start, Load, Exit"
	mapList := "Map1, Map2"
	currentEnemy := "Dragon"
	discovered := "Treasure"
	currentMap := "Map1"
	currentMapId := 1
	currentTileId := 101

	pointers := Pointers{
		PlayerPointer:         &player,
		ItemListPointer:       &itemList,
		IsGameInFrenchPointer: &isGameInFrench,
		MenuLinesPointer:      &menuLines,
		MapListPointer:        &mapList,
		CurrentEnemyPointer:   &currentEnemy,
		DiscoveredPointer:     &discovered,
		CurrentMapPointer:     &currentMap,
		CurrentMapIdPointer:   &currentMapId,
		CurrentTileIdPointer:  &currentTileId,
	}

	// Convert the struct to JSON
	jsonData, err := json.MarshalIndent(pointers, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Save the JSON to a file
	file, err := os.Create("save.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Pointers saved to save.json")
}
