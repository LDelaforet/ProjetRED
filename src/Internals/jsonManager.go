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

func GetItemById(id byte) (ItemObject, error) {
	for _, item := range *ItemListPointer {
		if item.Id == id {
			return item, nil
		}
	}
	return ItemObject{}, fmt.Errorf("ID %s introuvable", id)
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
