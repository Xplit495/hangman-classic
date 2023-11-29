package util

import "fmt"

func checkInputAlreadyUses(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, choiceToLowerStrings []string, letterFind bool, wordFind bool, letterAlreadyUse bool, wordAlreadyUse bool) {
	if letterAlreadyUse == true { //If user input is already use
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé cette lettre, attention !" + reset)
		fmt.Printf("\nPour le moment le mot ressemble à ca -> ")
		printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)                                       //Print error message
		startGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //Restart to select a new letter or word
	} else if wordAlreadyUse == true { //If user input is already use
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé ce mot, attention !" + reset) //Print error message
		fmt.Printf("\nPour le moment le mot ressemble à ca -> ")
		printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
		startGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //Restart to select a new letter or word
	}
	fmt.Println("")
	checkInputIsGood(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, choiceToLowerStrings, letterFind, wordFind) //If not already use, check if input is good
}

func checkInputIsGood(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, choiceToLowerStrings []string, letterFind bool, wordFind bool) {
	if len(choiceToLowerStrings) == 1 { //If user input is a letter
		if letterFind == true { //If letter is find
			ClearTerminal()
			fmt.Println(green + "Bonne lettre !" + reset) //Print success message
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)                                           //Show the word with the letter find
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //Check if game is over
		} else if letterFind == false { //If letter is not find
			liveJose-- //Remove one live
			ClearTerminal()
			fmt.Println(red + "Mauvaise lettre !" + reset) //Print error message
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ") //Show the word
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //Check if game is over
		}
	} else {
		if wordFind == false { //If user input is a word and word is not find
			liveJose = liveJose - 2 //Remove two live
			ClearTerminal()
			fmt.Println(red + "Mauvais mot !" + reset) //Print error message
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ") //Show the word
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //Check if game is over
		} else if wordFind == true { //If user input is a word and word is find
			ClearTerminal()
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)                                           //Show the word
			checkGameOver(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //Check if game is over
		}
	}
}
