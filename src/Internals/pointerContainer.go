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

var CraftingRecipes = []CraftingRecipe{}
var CraftingRecipesPointer = &CraftingRecipes

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
	// liste tous les items d'itemlist
	Chests[0] = []InventorySlot{
		{Item: ItemList[0], Quantity: 1},
		{Item: ItemList[10], Quantity: 2},
	}
	Chests[1] = []InventorySlot{
		{Item: ItemList[10], Quantity: 1},
		{Item: ItemList[11], Quantity: 1},
	}
	Chests[2] = []InventorySlot{
		{Item: ItemList[11], Quantity: 3},
		{Item: ItemList[2], Quantity: 1},
	}
	Chests[3] = []InventorySlot{
		{Item: ItemList[2], Quantity: 1},
		{Item: ItemList[10], Quantity: 2},
	}
	Chests[4] = []InventorySlot{
		{Item: ItemList[9], Quantity: 1},
	}
}

func EnnemisInit() {
	// Initialisation des ennemis
	Ennemis = append(Ennemis, Enemy{
		Type:        "GO-blin",
		PvMax:       15,
		Pv:          15,
		Damage:      5,
		Defence:     3,
		XpToDrop:    5,
		MoneyToDrop: 3,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Vampire",
		PvMax:       15,
		Pv:          15,
		Damage:      6,
		Defence:     0,
		XpToDrop:    8,
		MoneyToDrop: 5,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Elfe",
		PvMax:       18,
		Pv:          18,
		Damage:      9,
		Defence:     5,
		XpToDrop:    10,
		MoneyToDrop: 5,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Java-Gicien",
		PvMax:       20,
		Pv:          20,
		Damage:      8,
		Defence:     2,
		XpToDrop:    12,
		MoneyToDrop: 7,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Momie",
		PvMax:       30,
		Pv:          30,
		Damage:      5,
		Defence:     10,
		XpToDrop:    15,
		MoneyToDrop: 10,
	})
	Ennemis = append(Ennemis, Enemy{
		Type:        "Training Dummy",
		PvMax:       100,
		Pv:          100,
		Damage:      1,
		Defence:     0,
		XpToDrop:    0,
		MoneyToDrop: 0,
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
