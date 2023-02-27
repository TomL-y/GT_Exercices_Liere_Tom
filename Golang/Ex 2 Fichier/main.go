package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Recup(nom_fichier string) (string, error) {
	fichier, err := os.Open(nom_fichier)
	if err != nil {
		return "", err
	}
	defer fichier.Close()
	contenu, err := ioutil.ReadAll(fichier)
	if err != nil {
		return "", err
	}
	texte := string(contenu)
	return texte, nil
}

func lunch_Recup() {
	texte, err := Recup("document.txt")
	if err != nil {
		panic(err)
	}
	println(texte)
}

func AddTxt(nom_fichier string, texte_a_ajouter string) error {
	fichier, err := os.OpenFile(nom_fichier, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()
	_, err = fichier.WriteString(texte_a_ajouter)
	if err != nil {
		return err
	}
	return nil
}

func lunch_AddTxt() {
	err := AddTxt("document.txt", "Texte ajouté")
	if err != nil {
		panic(err)
	}
}

func Delete_contfichier(nom_fichier string) error {
	fichier, err := os.OpenFile(nom_fichier, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()
	return nil
}

func lunch_Delete_contfichier() {
	err := Delete_contfichier("document.txt")
	if err != nil {
		panic(err)
	}
}

func Replace_contfichier(nom_fichier string, texte_a_remplacer string) error {
	err := Delete_contfichier(nom_fichier)
	if err != nil {
		return err
	}
	err = AddTxt(nom_fichier, texte_a_remplacer)
	if err != nil {
		return err
	}
	return nil
}

func lunch_Replace_contfichier() {
	err := Replace_contfichier("document.txt", "Texte remplacé")
	if err != nil {
		panic(err)
	}
}

// Faire un menu pour choisir entre les fonctions et relancer le menu apres chaque choix rajouter un choix quitter le programme
func menu() {
	println("1 - Afficher le contenu d'un fichier")
	println("2 - Ajouter du texte")
	println("3 - Supprimer le contenu d'un fichier")
	println("4 - Remplacer le contenu d'un fichier")
	println("5 - Quitter")
}

func lunch_menu() {
	menu()
	var choix int
	println("Entrez votre choix")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		lunch_Recup()
		lunch_menu()
	case 2:
		lunch_AddTxt()
		lunch_menu()
	case 3:
		lunch_Delete_contfichier()
		lunch_menu()
	case 4:
		lunch_Replace_contfichier()
		lunch_menu()
	case 5:
		println("Au revoir")
		os.Exit(0)
	default:
		println("Choix invalide")
	}
}

func main() {
	lunch_menu()
}
