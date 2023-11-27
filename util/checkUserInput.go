package util

import "fmt"

func checkInputAlreadyUses(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, choiceToLowerStrings []string, letterFind bool, wordFind bool, letterAlreadyUse bool, wordAlreadyUse bool) {
	if letterAlreadyUse == true {
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé cette lettre, attention !" + reset)
		StartGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
	} else if wordAlreadyUse == true {
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé ce mot, attention !" + reset)
		StartGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
	}
	fmt.Println("")
	checkInputIsGood(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, choiceToLowerStrings, letterFind, wordFind)
}

func checkInputIsGood(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, choiceToLowerStrings []string, letterFind bool, wordFind bool) {
	if len(choiceToLowerStrings) == 1 {
		if letterFind == true {
			ClearTerminal()
			fmt.Println(green + "Bonne lettre !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
		} else if letterFind == false {
			liveJose--
			ClearTerminal()
			fmt.Println(red + "Mauvaise lettre !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
		}
	} else {
		if wordFind == false {
			liveJose = liveJose - 2
			ClearTerminal()
			fmt.Println(red + "Mauvais mot !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
		} else if wordFind == true {
			ClearTerminal()
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
		}
	}
}
