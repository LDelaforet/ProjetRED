package RED

import (
	"fmt"
	"strconv"
)

func BattleInit(en Enemy) {
	loadQuizData()

	PlayerPointer.Name = "Player"
	PlayerPointer.Damage = 20
	PlayerPointer.PvMax = 100
	PlayerPointer.Pv = 25

	CurrentEnemyPointer = &en

	CurrentEnemyPointer.Pv = 60

	BattleMain()
}

func BattleMain() {
	for {
		ClearScreen()
		DisplayText(DisplayTextOptions{
			TextToPrint: "Joueur: " + PlayerPointer.Name + " | " + strconv.Itoa(int(PlayerPointer.Pv)) + " / " + strconv.Itoa(int(PlayerPointer.PvMax)),
		})
		DisplayText(DisplayTextOptions{
			TextToPrint: "Ennemi: " + CurrentEnemyPointer.Type + " | " + strconv.Itoa(int(CurrentEnemyPointer.Pv)) + " / " + strconv.Itoa(int(CurrentEnemyPointer.PvMax)),
		})

		fmt.Println("1: Attaquer")
		fmt.Println("2: Défendre")
		fmt.Println("3: Objet ")
		fmt.Print("Choix: ")
		input := GetInput()
		switch input {
		case "1":
			attack()
		case "2":
			defend()
		case "3":
			item()
		}
		if CurrentEnemyPointer.Pv == 0 {
			break
		}
		if PlayerPointer.Pv == 0 {
			fmt.Println("T'es mort")
			break
		}
	}
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
		fmt.Println("Vous mettez un coup au Gobelin, il prends: " + strconv.Itoa(int(PlayerPointer.Damage)) + " dégats")
		// Si les degats que met le joeur sont superieurs aux pv de l'ennemi bah on le met direct a 0 pour eviter qu'ils soient négatifs car byte n'est pas signé
		if PlayerPointer.Damage >= CurrentEnemyPointer.Pv {
			CurrentEnemyPointer.Pv = 0
		} else {
			CurrentEnemyPointer.Pv -= PlayerPointer.Damage
		}
	} else {
		fmt.Println("Mauvaise réponse.")
	}

	if CurrentEnemyPointer.Pv > 0 {
		fmt.Println("Le gobelin vous met un coup, vous prenez: " + strconv.Itoa(int(CurrentEnemyPointer.Damage)) + " dégats")

		// Idem, le byte est non signé donc on dois checker qu'on soit a + de 0
		if CurrentEnemyPointer.Damage >= PlayerPointer.Pv {
			PlayerPointer.Pv = 0
		} else {
			PlayerPointer.Pv -= CurrentEnemyPointer.Damage
		}
	}

	input = GetInput()

}

func defend() {

}

func item() {

}
