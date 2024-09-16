package RED

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Fonctions pour fouiller les Listes de json parsés
func GetLineById(id string) string {
	for _, item := range *MenuLinesPointer {
		if item.Id == id {
			return item.Line
		}
	}
	return "/[" + id + "]\\"
}

func GetItemById(id byte) ItemObject {
	for _, item := range *ItemListPointer {
		if item.Id == id {
			return item
		}
	}
	return ItemObject{}
}

func GetMapById(id int) Map {
	if id < len(*MapListPointer) {
		return (*MapListPointer)[id]
	}
	return Map{}
}

func GetMapTileById(id int) MapTile {
	if *CurrentMapIdPointer < len(*MapListPointer) {
		for _, tile := range (*MapListPointer)[*CurrentMapIdPointer].Tiles {
			if tile.Id == id {
				return tile
			}
		}
	}
	return MapTile{}
}

// Parseurs de json

func ReadMenuStrings() {
	MenuLines = nil // Vide MenuLines

	// Ouvre le JSON et le met dans content
	content, _ := ioutil.ReadFile("./Database/menu.json")

	// Définir un dico dans lequel stocker les données du JSON
	var data map[string]map[string]string

	// Décode le JSON
	_ = json.Unmarshal(content, &data)

	// Check les données par rapport à la langue
	var languageData map[string]string
	if IsGameInFrench {
		languageData = data["i18n_fr"]
	} else {
		languageData = data["i18n_en"]
	}

	// Remplir MenuLines avec des TextLine
	for key, value := range languageData {
		MenuLines = append(MenuLines, TextLine{
			Id:   key,
			Line: value,
		})
	}
}

func ReadItemList() {
	ItemList = nil // Vide ItemList

	// Ouvre le JSON et le met dans content
	content, _ := ioutil.ReadFile("./Database/items.json")

	// Définir un dico dans lequel stocker les données du JSON
	var data map[string][]map[string]interface{}

	// Décode le JSON
	_ = json.Unmarshal(content, &data)

	// Check les données par rapport a la langue
	var languageData []map[string]interface{}
	if IsGameInFrench {
		languageData = data["i18n_fr"]
	} else {
		languageData = data["i18n_en"]
	}

	// Remplir ItemList avec des ItemObject
	for _, item := range languageData {
		id := item["Id"].(float64)
		name := item["Name"].(string)
		description := item["Description"].(string)

		ItemList = append(ItemList, ItemObject{
			Id:          byte(int(id)),
			Name:        name,
			Description: description,
		})
	}
}

func ReadMapLists() {
	content, err := ioutil.ReadFile("./Database/maps.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var mapData []interface{}
	err = json.Unmarshal(content, &mapData)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Clear MapLists before adding new maps
	*MapListPointer = nil

	for _, item := range mapData {
		// Type assertion to ensure the correct format
		mapArray, ok := item.([]interface{})
		if !ok || len(mapArray) < 2 {
			fmt.Println("Unexpected data format")
			continue
		}

		// Decode Tiles
		tilesData, ok := mapArray[0].([]interface{})
		if !ok {
			fmt.Println("Unexpected tiles data format")
			continue
		}

		var tiles []MapTile
		for _, tile := range tilesData {
			tileMap, ok := tile.(map[string]interface{})
			if !ok {
				fmt.Println("Unexpected tile data format")
				continue
			}

			var mapTile MapTile
			err := mapToStruct(tileMap, &mapTile)
			if err != nil {
				fmt.Println("Error converting tile map to struct:", err)
				continue
			}

			tiles = append(tiles, mapTile)
		}

		// Decode Grid
		gridData, ok := mapArray[1].([]interface{})
		if !ok {
			fmt.Println("Unexpected grid data format")
			continue
		}

		var grid [][]int
		for _, row := range gridData {
			rowData, ok := row.([]interface{})
			if !ok {
				fmt.Println("Unexpected row data format")
				continue
			}

			var gridRow []int
			for _, cell := range rowData {
				cellValue, ok := cell.(float64) // JSON numbers are decoded as float64
				if !ok {
					fmt.Println("Unexpected cell value format")
					continue
				}

				gridRow = append(gridRow, int(cellValue))
			}

			grid = append(grid, gridRow)
		}

		*MapListPointer = append(*MapListPointer, Map{Tiles: tiles, Grid: grid})
	}
}

// Helper function to convert map to struct
func mapToStruct(m map[string]interface{}, s interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}

// Partie relative au quizz:
var quizData Quiz

func loadQuizData() {
	content, _ := ioutil.ReadFile("./Database/quizz.json")
	json.Unmarshal(content, &quizData)
}

func getQuestion(category string, difficulty string) QuizzQuestion {
	var questions []QuizzQuestion
	switch category {
	case "GO":
		questions = quizData.GO[difficulty]
	case "Git":
		questions = quizData.Git[difficulty]
	default:
		return QuizzQuestion{}
	}

	if len(questions) == 0 {
		return QuizzQuestion{}
	}

	return questions[0]
}
