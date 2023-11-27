package util

import (
	"bufio"
	"fmt"
	Util "hangman-classic/util"
	"os"
)

var (
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
)

func Rules() {
	fmt.Println("")
	fmt.Println("Bienvenue dans ce super jeu, les régles sont simples :")
	fmt.Println("- Vous pouvez proposer ou un mot ou une lettre")
	fmt.Println("- Une mauvaise lettre vous retire" + yellow + " une " + reset + "vie. Mais " + red + "attention" + reset + " car un mauvais mot vous en retire" + yellow + " 2" + reset + " !")
	fmt.Print("Appuyer sur entrer pour continuer : ")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée standard")
		return
	}
	Util.chooseDifficulty()
}
