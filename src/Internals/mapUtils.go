package RED

import (
	"fmt"
)

var numLetters = map[int]string{
	10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j",
}

var boxTemplate = []string{"+-+", "| |", "+-+"}

// Relatif a l'affichage de la carte dans l'option Carte
func createBoxedMap(mapData [][]int) [][][]string {
	boxedMap := [][][]string{}
	for _, row := range mapData {
		newRow := [][]string{}
		for _, num := range row {
			if num == 0 || !contains(*DiscoveredPointer, num) {
				// Affiche le numéro de la case pour les cases non découvertes [TOGGLEABLE]

				// newRow = append(newRow, []string{"   ", " " + strconv.Itoa(num) + " ", "   "})
				newRow = append(newRow, []string{"   ", "   ", "   "})
			} else if num == CurrentTileId {
				newRow = append(newRow, []string{boxTemplate[0], "|X|", boxTemplate[2]})
			} else {
				newRow = append(newRow, []string{boxTemplate[0], boxTemplate[1], boxTemplate[2]})
			}
		}
		boxedMap = append(boxedMap, newRow)
	}
	return boxedMap
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func MapDisplay() {
	if CurrentMapId < len(MapList) {
		boxedMap := createBoxedMap(MapList[CurrentMapId].Grid)
		ClearScreen()
		for _, row := range boxedMap {
			for line := 0; line < len(boxTemplate); line++ {
				for _, box := range row {
					fmt.Print(box[line] + " ")
				}
				fmt.Println()
			}
		}
	}

	fmt.Print("Appuyez sur Entrée pour continuer.")
	fmt.Scanln()
}

func LiteMapDisplay() {
	if CurrentMapId < len(MapList) {
		boxedMap := createBoxedMap(MapList[CurrentMapId].Grid)
		for _, row := range boxedMap {
			for line := 0; line < len(boxTemplate); line++ {
				for _, box := range row {
					fmt.Print(box[line] + " ")
				}
				fmt.Println()
			}
		}
	}
}
