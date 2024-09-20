package RED

import (
	"fmt"
	"strconv"
)

var enemyTurnCounter int

func BattleInit(en Enemy) {
	loadQuizData()

	CurrentEnemyPointer = &en
	enemyTurnCounter = 0
	BattleMain()
}

func BattleMain() {
	for {
		ClearScreen()

		NewLine(2)
		DisplayLine()
		DisplayText(DisplayTextOptions{
			TextToPrint: PlayerPointer.Name + " fait face à un " + CurrentEnemyPointer.Type + " !",
		})
		DisplayLine()
		NewLine(2)
		DisplayText(DisplayTextOptions{
			TextToPrint: PlayerPointer.Name + ": " + strconv.Itoa(int(PlayerPointer.Pv)) + " / " + strconv.Itoa(int(PlayerPointer.PvMax)),
		})
		DisplayText(DisplayTextOptions{
			TextToPrint: CurrentEnemyPointer.Type + ": " + strconv.Itoa(int(CurrentEnemyPointer.Pv)) + " / " + strconv.Itoa(int(CurrentEnemyPointer.PvMax)),
		})
		NewLine(2)
		BoxStrings([]string{"0: " + GetLineById("attack"), "1: " + GetLineById("defend"), "2: " + GetLineById("useItem")})
		NewLine(2)
		DisplayLine()

		fmt.Print(GetLineById("choice") + ": ")
		input := GetInput()
		switch input {
		case "0":
			enemyTurnCounter++
			attack()
		case "1":
			defend()
		case "2":
			item()
		}
		if CurrentEnemyPointer.Pv < 1 {
			PlayerPointer.Money += int(CurrentEnemyPointer.MoneyToDrop)

			DisplayText(DisplayTextOptions{
				TextToPrint: "Vous avez vaincu le " + CurrentEnemyPointer.Type + " !\n" + PlayerPointer.Name + " gagne " + strconv.Itoa(int(CurrentEnemyPointer.XpToDrop)) + " points d'expérience et " + strconv.Itoa(int(CurrentEnemyPointer.MoneyToDrop)) + " " + GetLineById("money"),
			})
			GetInput()
			break
		}
		if PlayerPointer.Pv < 1 {
			fmt.Println("Vous etes malhencontreusement décédé (RIP BOZO)")
			GetInput()
			ClearScreen()
			NewLine(2)
			DisplayLine()
			DisplayText(DisplayTextOptions{
				TextToPrint: "GAME OVER",
			})
			DisplayLine()
			NewLine(2)
			DisplayText(DisplayTextOptions{
				TextToPrint: "0: Recommencer le combat avec 50% de votre vie",
			})
			DisplayText(DisplayTextOptions{
				TextToPrint: "1: Quitter le jeu",
			})
			NewLine(2)
			DisplayLine()
			fmt.Print(GetLineById("choice") + ": ")
			input := GetInput()
			if input == "0" {
				PlayerPointer.Pv = PlayerPointer.PvMax / 2
				BattleInit(*CurrentEnemyPointer)
			} else {
				ClearScreen()
				break
			}
		}
	}
}

func attack() {
	ClearScreen()

	NewLine(2)
	DisplayLine()
	DisplayText(DisplayTextOptions{
		TextToPrint: "Répond juste à la question pour attaquer!",
	})
	DisplayLine()
	NewLine(2)

	question := getQuestion(mapToTheme[CurrentMapId], "Facile")
	DisplayText(DisplayTextOptions{
		TextToPrint: "Question: " + question.Question,
	})
	NewLine(1)
	for i, reponse := range question.Reponses {
		fmt.Printf("%d: %s\n", i, reponse)
	}
	NewLine(2)
	DisplayLine()
	fmt.Print(GetLineById("choice") + ": ")
	input := GetInput()
	if input == strconv.Itoa(int(question.ReponseIndex)) {
		fmt.Println("Vous mettez un coup au " + CurrentEnemyPointer.Type + ", il prends: " + strconv.Itoa(int(PlayerPointer.Damage)) + " dégats")
		// Si les degats que met le joueur sont superieurs aux pv de l'ennemi bah on le met direct a 0 pour eviter qu'ils soient négatifs car byte n'est pas signé
		if PlayerPointer.Damage >= CurrentEnemyPointer.Pv {
			CurrentEnemyPointer.Pv = 0
		} else {
			CurrentEnemyPointer.Pv -= PlayerPointer.Damage
		}
	} else {
		fmt.Println("Mauvaise réponse.")
	}

	if CurrentEnemyPointer.Pv > 0 {
		if enemyTurnCounter < 3 {
			fmt.Println("Le " + CurrentEnemyPointer.Type + " vous inflige " + strconv.Itoa(int(CurrentEnemyPointer.Damage)) + " dégats")
			if CurrentEnemyPointer.Damage >= PlayerPointer.Pv {
				PlayerPointer.Pv = 0
			} else {
				PlayerPointer.Pv -= CurrentEnemyPointer.Damage
			}
		} else {
			fmt.Println("Le " + CurrentEnemyPointer.Type + " vous inflige " + strconv.Itoa(int(CurrentEnemyPointer.Damage)*2) + " dégats critique")
			if CurrentEnemyPointer.Damage*2 >= PlayerPointer.Pv {
				PlayerPointer.Pv = 0
			} else {
				PlayerPointer.Pv -= CurrentEnemyPointer.Damage * 2
				enemyTurnCounter = 0
			}
		}
	}
	NewLine(2)
	DisplayText(DisplayTextOptions{
		TextToPrint: "Appuyez sur une touche pour continuer...",
	})
	GetInput()
}

func defend() {
	ClearScreen()
	if !(CurrentEnemyPointer.Damage < PlayerPointer.Defence) {
		if (CurrentEnemyPointer.Damage - PlayerPointer.Defence) >= PlayerPointer.Pv {
			PlayerPointer.Pv = 0
		} else {
			PlayerPointer.Pv -= (CurrentEnemyPointer.Damage - PlayerPointer.Defence)
		}
		DisplayText(DisplayTextOptions{
			TextToPrint: "Le " + CurrentEnemyPointer.Type + " vous porte un coup que vous bloquez, vous subissez hélas " + strconv.Itoa(int((CurrentEnemyPointer.Damage - PlayerPointer.Defence))) + " dégats",
		})
	} else {
		DisplayText(DisplayTextOptions{
			TextToPrint: "Le " + CurrentEnemyPointer.Type + " vous porte un coup que vous bloquez",
		})
	}
	DisplayText(DisplayTextOptions{
		TextToPrint: PlayerPointer.Name + " profite de cet instant pour se soigner " + strconv.Itoa(int(PlayerPointer.Heal)) + " " + GetLineById("hp"),
	})
	NewLine(2)
	DisplayText(DisplayTextOptions{
		TextToPrint: "Appuyez sur une touche pour continuer...",
	})
	if (PlayerPointer.Heal + PlayerPointer.Pv) > PlayerPointer.PvMax {
		PlayerPointer.Pv = PlayerPointer.PvMax
	} else {
		PlayerPointer.Pv += PlayerPointer.Heal
	}
	// Idem, le byte est non signé donc on dois checker qu'on soit a + de 0
	// Ok mec
	GetInput()
}

func item() {
	AccessInventory("battle")
}

// Systeme d'inventaire
func AccessInventory(currentState string) {
	ClearScreen()
	NewLine(2)
	DisplayLine()
	DisplayText(DisplayTextOptions{
		TextToPrint: GetLineById("inventory"),
	})
	DisplayLine()
	NewLine(2)
	if !(len(PlayerPointer.Inventory) == 0) {
		for i := 0; i < len(PlayerPointer.Inventory); i++ {
			DisplayText(DisplayTextOptions{
				TextToPrint: " - " + PlayerPointer.Inventory[i].Item.Name + " | x" + strconv.Itoa(int(PlayerPointer.Inventory[i].Quantity)) + "\n   - " + PlayerPointer.Inventory[i].Item.Description,
			})
			NewLine(1)
		}
		NewLine(1)
		DisplayText(DisplayTextOptions{
			TextToPrint: "0: " + GetLineById("useItem"),
		})
	} else {
		DisplayText(DisplayTextOptions{
			TextToPrint: GetLineById("emptyInventory"),
		})
	}
	NewLine(1)
	DisplayText(DisplayTextOptions{
		TextToPrint: "1: " + GetLineById("quit"),
	})
	NewLine(1)
	DisplayLine()
	fmt.Printf("Choix: ")
	input := GetInput()

	if input == "0" {
		ClearScreen()
		NewLine(2)
		DisplayLine()
		DisplayText(DisplayTextOptions{
			TextToPrint: GetLineById("inventory"),
		})
		DisplayLine()
		NewLine(2)

		for i := 0; i < len(PlayerPointer.Inventory); i++ {
			DisplayText(DisplayTextOptions{
				TextToPrint: strconv.Itoa(i) + ": " + PlayerPointer.Inventory[i].Item.Name + " | x" + strconv.Itoa(int(PlayerPointer.Inventory[i].Quantity)),
			})
			NewLine(1)
		}
		NewLine(1)
		DisplayLine()
		fmt.Printf("Choisir un objet à utiliser: ")
		input := GetInput()
		index, err := strconv.Atoi(input)
		if err == nil && index >= 0 && index < len(PlayerPointer.Inventory) {
			if PlayerPointer.Inventory[index].Item.Id == 0 {
				PlayerPointer.Pv += 5
				if PlayerPointer.Pv > PlayerPointer.PvMax {
					PlayerPointer.Pv = PlayerPointer.PvMax
				}
			} else if PlayerPointer.Inventory[index].Item.Id == 1 {
				PlayerPointer.Damage += 2
			} else if PlayerPointer.Inventory[index].Item.Id == 2 {
				PlayerPointer.PvMax += 5
			} else if PlayerPointer.Inventory[index].Item.Id == 9 {
				PlayerPointer.Pv = 0
			} else if PlayerPointer.Inventory[index].Item.Id == 99 {
				PlayerPointer.InventorySize += 5
			}
			if PlayerPointer.Inventory[index].Quantity > 1 {
				PlayerPointer.Inventory[index].Quantity -= 1
			} else {
				PlayerPointer.Inventory = append(PlayerPointer.Inventory[:index], PlayerPointer.Inventory[index+1:]...)
			}
		}
	}
	if input == "1" {
		if currentState == "battle" {
			BattleMain()
		} else {
			return
		}

	}

}

func AddItemToInventory(itemId byte, quantite byte) {
	for i, slot := range PlayerPointer.Inventory {
		if slot.Item.Id == itemId {
			PlayerPointer.Inventory[i].Quantity += quantite
			return
		}
	}
	PlayerPointer.Inventory = append(PlayerPointer.Inventory, InventorySlot{Item: ItemList[itemId], Quantity: quantite})
}

func RemoveItemFromInventory(itemId byte, quantite byte) {
	for i, slot := range PlayerPointer.Inventory {
		if slot.Item.Id == itemId {
			if slot.Quantity > quantite {
				PlayerPointer.Inventory[i].Quantity -= quantite
			} else {
				PlayerPointer.Inventory = append(PlayerPointer.Inventory[:i], PlayerPointer.Inventory[i+1:]...)
			}
			return
		}
	}
}

func checkIfItemInInventory(itemID byte, quantity byte) bool {
	for _, slot := range PlayerPointer.Inventory {
		if slot.Item.Id == itemID {
			if slot.Quantity >= quantity {
				return true
			}
		}
	}
	return false
}
