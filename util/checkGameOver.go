package util

import (
	"fmt"
)

func checkGameOver(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {

	wordFind := true                             //If word is find
	for _, letter := range wordPartiallyReveal { //Move in the wordPartiallyReveal array
		if letter == "_" { //If there is a "_" in the array
			wordFind = false //Word is not find
			break
		}
	}
	if wordFind == true { //If word is find
		ClearTerminal()
		fmt.Print("\n" + green + "Vous avez deviné le mot !" + reset + "\nLe mot était : ") //Print success message
		printWord(asciiMode, pathAscii, arrSelectWord)                                      //Print the word completely reveal
		fmt.Println("")

		if len(letterHistory) > 0 { //If letterHistory array is not empty show the letter history
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory(letterHistory)
		}
		if len(wordHistory) > 0 { //If wordHistory array is not empty show the word history
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory(wordHistory)
		}

		addStats(asciiMode, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, wordFind) //Call function to add stats
		restartGame(asciiMode, pathAscii)                                                                       //Call function to restart the game

	} else if liveJose <= 0 { //If live is equal or less than 0
		ClearTerminal()
		fmt.Print("\n" + red + "Vous n'avez plus de vie !" + reset + "\nLe mot était : ") //Print lose message
		printWord(asciiMode, pathAscii, arrSelectWord)                                    //Print the word completely reveal
		printJose(71, 78)                                                                 //Print Jose completely hang

		if len(letterHistory) > 0 { //If letterHistory array is not empty show the letter history
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory(letterHistory)
		}

		if len(wordHistory) > 0 { //If wordHistory array is not empty show the word history
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory(wordHistory)
		}
		fmt.Println("")
		fmt.Println(red + "Vous êtes pendu !" + reset) //Print lose message

		addStats(asciiMode, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, wordFind) //Call function to add stats
		restartGame(asciiMode, pathAscii)                                                                       //Call function to restart the game

	} else {
		startGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose) //If game is not over recall the startGame function
	}
}
