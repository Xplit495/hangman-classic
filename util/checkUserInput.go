package util

import "fmt"

func checkInputAlreadyUses(choiceToLowerStrings []string, letterFind bool, wordFind bool, letterAlreadyUse bool, wordPartiallyReveal []string, arrSelectWord []string, wordAlreadyUse bool) {
	if letterAlreadyUse == true {
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé cette lettre, attention !" + reset)
		StartGame(arrSelectWord, wordPartiallyReveal, liveJose)
	} else if wordAlreadyUse == true {
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé ce mot, attention !" + reset)
		StartGame(arrSelectWord, wordPartiallyReveal, liveJose)
	}
	fmt.Println("")
	checkFindLetter(choiceToLowerStrings, wordPartiallyReveal, arrSelectWord, letterFind, wordFind)
}

func checkInputIsGood(choiceToLowerStrings []string, wordPartiallyReveal []string, arrSelectWord []string, letterFind bool, wordFind bool) {
	if len(choiceToLowerStrings) == 1 {
		if letterFind == true {
			ClearTerminal()
			fmt.Println(green + "Bonne lettre !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		} else if letterFind == false {
			liveJose--
			ClearTerminal()
			fmt.Println(red + "Mauvaise lettre !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		}
	} else {
		if wordFind == false {
			liveJose = liveJose - 2
			ClearTerminal()
			fmt.Println(red + "Mauvais mot !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		} else if wordFind == true {
			ClearTerminal()
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		}
	}
}
