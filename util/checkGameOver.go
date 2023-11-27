package util

import "fmt"

func checkGameOver(wordPartiallyReveal []string, arrSelectWord []string) {
	wordFind := true
	for _, letter := range wordPartiallyReveal {
		if letter == "_" {
			wordFind = false
			break
		}
	}
	if wordFind == true {
		ClearTerminal()
		fmt.Println("\n" + green + "Vous avez deviné le mot !" + reset)
		if len(letterHistory) > 0 {
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory()
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory()
		}
		fmt.Print("Le mot était : ")
		printWord(arrSelectWord)
		fmt.Println("")
	} else if liveJose <= 0 {
		ClearTerminal()
		fmt.Print("\n" + red + "Vous n'avez plus de vie !" + reset + "\nLe mot était : ")
		printWord(arrSelectWord)
		printJose(71, 78)
		if len(letterHistory) > 0 {
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory()
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory()
		}
		fmt.Println("")
		fmt.Println(red + "Vous êtes pendu !" + reset)
	} else {
		StartGame(arrSelectWord, wordPartiallyReveal, liveJose)
	}
}
