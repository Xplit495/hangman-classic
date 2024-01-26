package util

import (
	"fmt"
	"os"
	"strings"
)

func restartGame(asciiMode string, pathAscii string) {
	var restart string

	fmt.Println("")
	fmt.Print("Voulez-vous rejouer ? (o/n) : ") //Ask if user want to restart
	_, err := fmt.Scan(&restart)                //Scan user input
	if err != nil {
		fmt.Println("Erreur lors de la saisie de la réponse")
	}
	restart = strings.ToLower(restart) //Convert user input to lowercase
	for i := 0; i <= 1; i++ {

		if restart == "o" { //If user want to restart
			PrintRules(asciiMode, pathAscii) //Print rules (to restart the game)

		} else if restart == "n" { //If user don't want to restart
			fmt.Println("\nMerci d'avoir joué !") //Print end message
			fmt.Println("")
			os.Exit(0) //Exit the program

		} else {
			i-- //If user input is not valid, restart the loop
			ClearTerminal()
			fmt.Print("Veuillez entrer une réponse valide (o/n) : ")
			_, err2 := fmt.Scan(&restart)
			if err2 != nil {
				fmt.Println("Erreur lors de la saisie de la réponse")
			}
			restart = strings.ToLower(restart) //Convert user input to lowercase
		}
	}
}
