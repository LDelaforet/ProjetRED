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

var Chests = map[int][]InventorySlot{}
var ChestsPointer = &Chests

var Ennemis = []Enemy{}
var EnnemisPointer = &Ennemis

var mapToTheme = map[int]string{0: "GO", 1: "Git", 2: "Misc", 3: "Java", 4: "Javascript"}

func PointersInit() {
	PlayerPointer = &Perso{}
	ItemListPointer = &ItemList
	IsGameInFrenchPointer = &IsGameInFrench
	MenuLinesPointer = &MenuLines
	MapListPointer = &MapList
	CurrentEnemyPointer = &CurrentEnemy
}

func ChestsInit() {
	// Initialisation du contenu des coffres
	Chests[0] = []InventorySlot{
		{Item: ItemList[8], Quantity: 1},
		{Item: ItemList[11], Quantity: 1},
		{Item: ItemList[5], Quantity: 1},
	}
	Chests[1] = []InventorySlot{}
	Chests[2] = []InventorySlot{}
	Chests[3] = []InventorySlot{}
	Chests[4] = []InventorySlot{}
}

func EnnemisInit() {
	// Initialisation des ennemis
	Ennemis = append(Ennemis, Enemy{
		Type:        "GO-blin",
		PvMax:       15,
		Pv:          10,
		Damage:      3,
		Defence:     2,
		XpToDrop:    5,
		MoneyToDrop: 3,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Vampire",
		PvMax:       15,
		Pv:          10,
		Damage:      3,
		Defence:     2,
		XpToDrop:    5,
		MoneyToDrop: 3,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Elfe",
		PvMax:       15,
		Pv:          10,
		Damage:      3,
		Defence:     2,
		XpToDrop:    5,
		MoneyToDrop: 3,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Java-Gicien",
		PvMax:       15,
		Pv:          10,
		Damage:      3,
		Defence:     2,
		XpToDrop:    5,
		MoneyToDrop: 3,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Momie",
		PvMax:       15,
		Pv:          10,
		Damage:      3,
		Defence:     2,
		XpToDrop:    5,
		MoneyToDrop: 3,
	})
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
