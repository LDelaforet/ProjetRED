package RED

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadItemList(textToPrint string) []ItemObject {
	itemList := []ItemObject{}

	content, err := ioutil.ReadFile("../Database/items.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload ItemObject
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return (itemList)
}
