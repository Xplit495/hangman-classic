package util

import (
	"fmt"
	"os"
	"strings"
)

func checkGameOver(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {
	var restart string

	wordFind := true
	for _, letter := range wordPartiallyReveal {
		if letter == "_" {
			wordFind = false
			break
		}
	}
	if wordFind == true {
		ClearTerminal()
		fmt.Print("\n" + green + "Vous avez deviné le mot !" + reset + "\nLe mot était : ")
		printWord(asciiMode, pathAscii, arrSelectWord)
		fmt.Println("")

		if len(letterHistory) > 0 {
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory(letterHistory)
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory(wordHistory)
		}

		fmt.Println("")
		fmt.Print("Voulez-vous rejouer ? (o/n) : ")
		fmt.Scan(&restart)
		restart = strings.ToLower(restart)
		for i := 0; i <= 1; i++ {

			if restart == "o" {
				PrintRules(asciiMode, pathAscii)

			} else if restart == "n" {
				fmt.Println("\nMerci d'avoir joué !")
				fmt.Println("")
				os.Exit(0)

			} else {
				i--
				ClearTerminal()
				fmt.Print("Veuillez entrer une réponse valide (o/n) : ")
				fmt.Scan(&restart)
				restart = strings.ToLower(restart)
			}
		}

	} else if liveJose <= 0 {
		ClearTerminal()
		fmt.Print("\n" + red + "Vous n'avez plus de vie !" + reset + "\nLe mot était : ")
		printWord(asciiMode, pathAscii, arrSelectWord)
		printJose(71, 78)

		if len(letterHistory) > 0 {
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory(letterHistory)
		}

		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory(wordHistory)
		}
		fmt.Println("")
		fmt.Println(red + "Vous êtes pendu !" + reset)

		fmt.Println("")
		fmt.Print("Voulez-vous rejouer ? (o/n) : ")
		fmt.Scan(&restart)
		restart = strings.ToLower(restart)
		for i := 0; i <= 1; i++ {

			if restart == "o" {
				PrintRules(asciiMode, pathAscii)

			} else if restart == "n" {
				fmt.Println("\nMerci d'avoir joué !")
				fmt.Println("")
				os.Exit(0)

			} else {
				i--
				ClearTerminal()
				fmt.Print("Veuillez entrer une réponse valide (o/n): ")
				fmt.Scanln(&restart)
				restart = strings.ToLower(restart)
			}
		}

	} else {
		startGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
	}
}
