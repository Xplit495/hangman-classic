package util

import "fmt"

func checkGameOver(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {
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
			printLetterHistory(letterHistory)
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory(wordHistory)
		}
		fmt.Print("Le mot était : ")
		printWord(asciiMode, pathAscii, arrSelectWord)
		fmt.Println("")
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
	} else {
		StartGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
	}
}
