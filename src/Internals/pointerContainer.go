package RED

// pointerContainer est un fichier dans lequel on va déclarer toutes les variables en pointeurs afin de les utiliser globalement a travers tout le code.
// Afin de l'utiliser, il faut tout d'abord appeler PointeursInit(), il se charge d'assigner aux variables containant les pointeurs aux dits pointeurs

var PlayerPointer *Perso

var ItemList []ItemObject
var ItemListPointer *[]ItemObject

var IsGameInFrench bool = true
var IsGameInFrenchPointer *bool

var MenuLines []TextLine
var MenuLinesPointer *[]TextLine

var MapList []Map
var MapListPointer *[]Map

var CurrentEnemy Enemy
var CurrentEnemyPointer *Enemy

var Discovered = []int{}
var DiscoveredPointer = &Discovered

var CurrentMap = Map{}
var CurrentMapPointer = &CurrentMap

var CurrentMapId = 0
var CurrentMapIdPointer = &CurrentMapId

var CurrentTileId = 0
var CurrentTileIdPointer = &CurrentTileId

func PointersInit() {
	PlayerPointer = &Perso{}
	ItemListPointer = &ItemList
	IsGameInFrenchPointer = &IsGameInFrench
	MenuLinesPointer = &MenuLines
	MapListPointer = &MapList
	CurrentEnemyPointer = &CurrentEnemy
}

// Partie test de pointeurs, sera amenée a changer

func TestPlayerPointer(persoPointer *Perso) {
	persoPointer.Name = "EDITED NAME"
}

func TestItemListPointer(itemListPointer *[]ItemObject) {
	*itemListPointer = append(*itemListPointer, ItemObject{
		Id:          164,
		Name:        "EDITED",
		Description: "EDITED",
	})
}

func TestIsGameInFrenchPointer(isGameInFrenchPointer *bool) {
	*isGameInFrenchPointer = false
}
