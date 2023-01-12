package module

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(" salut  %v ", name)
	return message, nil
}

func Saisie() {

	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("Entrez quelque chose : ")
	scanner.Scan()                      // lancement du scanner
	entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
	fmt.Println(entreeUtilisateur)

}
