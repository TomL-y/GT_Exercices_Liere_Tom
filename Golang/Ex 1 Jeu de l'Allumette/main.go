package main

import (
	"fmt"
)

func main() {
	var allumettes int
	fmt.Println("Choisis un nombre d'allumettes supérieur à 4 :")
	fmt.Scanln(&allumettes)

	if allumettes < 4 {
		fmt.Println("Le nombre d'allumettes doit être supérieur à 4")
		fmt.Scanln(&allumettes)
	}
	player1 := "Joueur 1"
	player2 := "Joueur 2"
	player_actuel := player1

	for allumettes > 0 {
		allumettes_choisies := 0
		allumettes_restantes := 0
		fmt.Println("Au tour du ", player_actuel)
		fmt.Println("Choisissez entre 1 et 3 allumettes :")
		fmt.Scanln(&allumettes_choisies)
		if allumettes_choisies < 1 || allumettes_choisies > 3 {
			fmt.Println("Vous devez chosir entre 1 et 3 allumettes :")
			fmt.Scanln(&allumettes_choisies)
		} else {
			allumettes_restantes = allumettes - allumettes_choisies
		}
		if player_actuel == player1 && allumettes_restantes > 0 {
			player_actuel = player2
		} else if allumettes_restantes > 0 {
			player_actuel = player1
		}
		fmt.Println("Il reste", allumettes_restantes, "allumettes")
		allumettes = allumettes_restantes
		if allumettes_restantes == 0 {
			fmt.Println("Le ", player_actuel, "a perdu !")
		}
	}
}
