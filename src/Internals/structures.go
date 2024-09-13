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
	Name        string
	Description string
}

type TextLine struct {
	Id   string
	Line string
}

type MapTiles struct {
	Id        byte
	EventType byte
	ToLeftID  byte
	ToRightID byte
	ToUpID    byte
	ToDownID  byte
}
type InventorySlot struct {
	Item     ItemObject
	Quantity byte
}

type Perso struct {
	Name      string
	Class     byte
	Level     byte
	Xp        byte
	PvMax     byte
	Pv        byte
	Inventory []InventorySlot
	Damage    byte
	Defence   byte
}

type Enemy struct {
	Type    string
	PvMax   byte
	Pv      byte
	Damage  byte
	Defence byte
}
