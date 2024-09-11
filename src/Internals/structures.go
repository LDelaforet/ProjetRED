package RED

// Structures générales
type ItemObject struct {
	Id          byte
	Name        string
	Description string
}

// Structures relatives au personnage
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
}
