package RED

import (
	"encoding/json"
	"io/ioutil"
)

func ReadItemList() []ItemObject {
	// ItemList est une liste contenant des ItemObjects
	var ItemList []ItemObject

	// Ouvre le JSON et le met dans content
	content, _ := ioutil.ReadFile("./Database/items.json")

	// Définis un dico dans lequel stocker les données du JSON
	var data map[string][]map[string]interface{}

	// Rend le JSON lisible
	_ = json.Unmarshal(content, &data)

	// Check les données par rapport a la langue
	var languageData []map[string]interface{}
	if isGameInFrench {
		languageData = data["i18n_fr"]
	} else {
		languageData = data["i18n_en"]
	}

	// Remplis ItemList avec des ItemObject
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
	return ItemList
}
