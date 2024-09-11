package RED

type InventorySlot struct {
	Id       byte
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
