package RED

import (
	"encoding/json"
	"os"
)

type SaveData struct {
	PlayerPointer  *Perso `json:"player_pointer"`
	IsGameInFrench *bool  `json:"is_game_in_french"`
	Discovered     *[]int `json:"discovered"`
	CurrentMapId   *int   `json:"current_map_id"`
	CurrentTileId  *int   `json:"current_tile_id"`
}

func SaveToJSON() error {
	data := SaveData{
		PlayerPointer:  PlayerPointer,
		IsGameInFrench: IsGameInFrenchPointer,
		Discovered:     DiscoveredPointer,
		CurrentMapId:   CurrentMapIdPointer,
		CurrentTileId:  CurrentTileIdPointer,
	}

	file, err := os.Create("save.sav")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func LoadFromJSON() error {
	file, err := os.Open("save.sav")
	if err != nil {
		return err
	}
	defer file.Close()

	var data SaveData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	PlayerPointer = data.PlayerPointer
	IsGameInFrenchPointer = data.IsGameInFrench
	DiscoveredPointer = data.Discovered
	CurrentMapIdPointer = data.CurrentMapId
	CurrentTileIdPointer = data.CurrentTileId

	return nil
}
