package RED

import (
	color "github.com/fatih/color"
)

// Structures spécifiques a des fonctions
type DisplayTextOptions struct {
	TextToPrint string
	IsCentered  bool
	Offset      int
	FgColor     color.Attribute
	BgColor     color.Attribute
	Bold        bool
	Underline   bool
}

// Structures générales
type ItemObject struct {
	Id          byte
	ItemType    byte   // 0: Tout, 1: Arme
	Name        string // 2: Armure,
	Description string //3: Consommables
	Price       int    //4: Materiaux
}

type CraftingMaterial struct {
	ItemID   int
	Quantity int
}

type CraftingRecipe struct {
	Ingredients []CraftingMaterial
	Result      byte
}

type QuizzQuestion struct {
	Id           byte
	Question     string
	Reponses     []string
	ReponseIndex byte
}

type Quiz struct {
	GO         map[string][]QuizzQuestion
	Git        map[string][]QuizzQuestion
	Misc       map[string][]QuizzQuestion
	Java       map[string][]QuizzQuestion
	Javascript map[string][]QuizzQuestion
}

type TextLine struct {
	Id   string
	Line string
}

type MapTile struct {
	Id        int
	EventType int
	ToLeftID  int
	ToRightID int
	ToUpID    int
	ToDownID  int
}
type Map struct {
	Tiles []MapTile
	Grid  [][]int
}

type InventorySlot struct {
	Item     ItemObject
	Quantity byte
}

type Perso struct {
	Name          string
	Class         byte // 0: none, 2: Attaque, 3: Défense, 4: soin
	Level         byte
	Xp            byte
	PvMax         byte
	Pv            byte
	Inventory     []InventorySlot
	EquipedItem   ItemObject
	Damage        byte
	Defence       byte
	Heal          byte
	Money         int
	EquippedArmor Equipment
	InventorySize int
}

type Equipment struct {
	Helmet ArmorSlot
	Chest  ArmorSlot
	Boots  ArmorSlot
}
type ArmorSlot struct {
	Name         string
	SlotType     int
	PvBoost      int
	DamageBoost  int
	DefenceBoost int
}

type Enemy struct {
	Type        string
	PvMax       byte
	Pv          byte
	Damage      byte
	Defence     byte
	XpToDrop    byte
	MoneyToDrop byte
}
