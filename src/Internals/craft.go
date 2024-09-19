package RED

import (
	"fmt"
	"strconv"
)

// craftingMenu devrait output un truc comme ca:
// Recettes disponibles:
// 1: Epée: 2x Fer, 1x Cuir
// choix:
// Ensuite il recupère l'input du joueur

func CraftingMenu() {
	fmt.Println("Recettes disponibles:")
	for i, recipe := range *CraftingRecipesPointer {
		fmt.Print(i+1, ": ")
		if !checkIfCraftable(recipe) {
			fmt.Print("[PAS ASSEZ DE RESSOURCES] ")
		}
		for j, ingredient := range recipe.Ingredients {
			if j != 0 {
				fmt.Print(", ")
			}
			fmt.Print(GetItemById(byte(ingredient.ItemID)).Name)
			fmt.Print(" x", ingredient.Quantity)
		}
		fmt.Println(" -> ", GetItemById(recipe.Result).Name)
	}
	fmt.Print("choix: ")
	input := GetInput()
	recipeIndex, _ := strconv.Atoi(input)
	if recipeIndex > 0 && recipeIndex <= len(*CraftingRecipesPointer) {
		if checkIfCraftable((*CraftingRecipesPointer)[recipeIndex-1]) {
			CraftItem((*CraftingRecipesPointer)[recipeIndex-1])
		} else {
			fmt.Println("Recette non craftable")
		}
	}
}

func CraftItem(recipe CraftingRecipe) {
	for _, ingredient := range recipe.Ingredients {
		RemoveItemFromInventory(byte(ingredient.ItemID), byte(ingredient.Quantity))
	}
	AddItemToInventory(recipe.Result, 1)
}

func checkIfCraftable(recipe CraftingRecipe) bool {
	for _, ingredient := range recipe.Ingredients {
		if !checkIfItemInInventory(byte(ingredient.ItemID), byte(ingredient.Quantity)) {
			return false
		}
	}
	return true
}
