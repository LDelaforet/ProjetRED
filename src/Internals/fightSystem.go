package RED

import (
	"fmt"
	"strconv"
)

func BattleInit(en Enemy) {
	loadQuizData()

	CurrentEnemyPointer = &en

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
			TextToPrint: "Joueur: " + PlayerPointer.Name + " | " + strconv.Itoa(int(PlayerPointer.Pv)) + " / " + strconv.Itoa(int(PlayerPointer.PvMax)),
		})
		DisplayText(DisplayTextOptions{
			TextToPrint: "Ennemi: " + CurrentEnemyPointer.Type + " | " + strconv.Itoa(int(CurrentEnemyPointer.Pv)) + " / " + strconv.Itoa(int(CurrentEnemyPointer.PvMax)),
		})
		NewLine(2)
		BoxStrings([]string{"0: " + GetLineById("attack"), "1: " + GetLineById("defend"), "2: " + GetLineById("useItem")})
		NewLine(2)
		DisplayLine()

		fmt.Print("Choix: ")
		input := GetInput()
		switch input {
		case "0":
			attack()
		case "1":
			defend()
		case "2":
			item()
		}
		if CurrentEnemyPointer.Pv == 0 {
			break
		}
		if PlayerPointer.Pv == 0 {
			fmt.Println("Vous etes décédé malhencontreusement (RIP BOZO)")
			break
		}
	}
	NewLine(2)
	DisplayLine()
	_ = GetInput()
}

func attack() {
	ClearScreen()

	question := getQuestion("GO", "Facile")

	fmt.Printf("%s: \n", question.Question)
	for i, reponse := range question.Reponses {
		fmt.Printf("%d: %s\n", i, reponse)
	}
	fmt.Print("Choix: ")
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
		fmt.Println("Le " + CurrentEnemyPointer.Type + " vous inflige " + strconv.Itoa(int(CurrentEnemyPointer.Damage)) + " dégats")

		// Idem, le byte est non signé donc on dois checker qu'on soit a + de 0
		if CurrentEnemyPointer.Damage >= PlayerPointer.Pv {
			PlayerPointer.Pv = 0
		} else {
			PlayerPointer.Pv -= CurrentEnemyPointer.Damage
		}
	}
	NewLine(2)
	DisplayText(DisplayTextOptions{
		TextToPrint: "Appuyez sur une touche pour continuer...",
	})

	_ = GetInput()

}

func defend() {
	ClearScreen()
	DisplayText(DisplayTextOptions{
		TextToPrint: "Le " + CurrentEnemyPointer.Type + " vous porte un coup que vous bloquez, vous subissez hélas " + strconv.Itoa(int((CurrentEnemyPointer.Damage - PlayerPointer.Defence))) + " dégats",
	})
	if (CurrentEnemyPointer.Damage - PlayerPointer.Defence) >= PlayerPointer.Pv {
		PlayerPointer.Pv = 0
	} else {
		PlayerPointer.Pv -= (CurrentEnemyPointer.Damage - PlayerPointer.Defence)
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
	_ = GetInput()
}

func item() {
	InventoryDisplay()
}

func InventoryDisplay() {
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
				TextToPrint: " - " + PlayerPointer.Inventory[i].Item.Name + " | " + strconv.Itoa(int(PlayerPointer.Inventory[i].Quantity)),
			})
		}
	} else {
		DisplayText(DisplayTextOptions{
			TextToPrint: GetLineById("emptyInventory"),
		})
	}

	NewLine(2)
	DisplayLine()
	fmt.Println("Choix: ")
	_ = GetInput()
	BattleMain()
}
