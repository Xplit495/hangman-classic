package util

import (
	"fmt"
	"strings"
)

func startGame(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {
	var (
		choice               string
		choiceToLower        string
		choiceToLowerStrings []string
		choiceToLowerRune    []rune
	)

	chooseLiveJose(liveJose)
	fmt.Println("")
	fmt.Printf("Il vous reste "+yellow+"%d vie "+reset+"avant d'être pendu !\n", liveJose)
	fmt.Println("")
	if len(letterHistory) > 0 { //If letterHistory is not empty
		fmt.Print("Les lettres déjà essayé sont : ")
		printLetterHistory(letterHistory)
	}
	if len(wordHistory) > 0 { //If wordHistory is not empty
		fmt.Print("Les mots déjà essayé sont : ")
		printWordHistory(wordHistory)
	}

	for i := 0; i <= 1; i++ {
		choiceToLowerStrings = nil
		fmt.Print("Entrez votre lettre ou votre mot : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée standard")
			return
		}
		choiceToLower = strings.ToLower(choice) //Convert choice to lowercase

		if choiceToLower == "stop" { //If choiceToLower is equal to "stop" call the function createGameSave
			createGameSave(arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
		}

		choiceToLowerRune = []rune(choiceToLower)                                              //Convert choiceToLower to rune
		if len(wordPartiallyReveal) == len(choiceToLowerRune) || len(choiceToLowerRune) == 1 { //If wordPartiallyReveal is equal to choiceToLower or choiceToLower is equal to 1 (if user input is a word of equal size than the word to find or if user input is a letter)
			for j := 0; j < len(choiceToLowerRune); j++ { //For each letter of choiceToLowerRune
				choiceToLowerStrings = append(choiceToLowerStrings, string(choiceToLowerRune[j])) //Append the letter to choiceToLowerStrings
			}
			exit := true
			for k := 0; k < len(choiceToLowerRune); k++ {
				if choiceToLowerRune[k] >= rune(97) && choiceToLowerRune[k] <= rune(122) { //Check if the letter is between a and z
					if k+1 == len(choiceToLowerRune) {
						break
					}
				} else { //If the letter is not between a and z prevent the user and ask him to enter only letters
					ClearTerminal()
					fmt.Println("Merci de saisir" + red + " UNIQUEMENT " + reset + "des caractère de l'alphabet !")
					fmt.Print("\nPour le moment le mot ressemble à ca -> ")
					printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
					fmt.Println("")
					exit = false //Set exit false to prevent the loop to break
					i--
				}
			}
			if exit == true {
				break
			}
		} else {
			ClearTerminal()
			fmt.Println("Merci de saisir " + red + "UNIQUEMENT " + reset + "une lettre ou un mot (de même longeur) !")
			fmt.Print("\nPour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
			fmt.Println("")
			i-- //Decrement i to ask the user to enter a letter or a word again
		}
	}
	if len(choiceToLowerStrings) == 1 { //If choiceToLowerStrings is equal to 1 (if user input is a letter)
		for i := 0; i < len(choiceToLowerStrings); i++ {
			letterHistory = append(letterHistory, choiceToLowerStrings[i]) //Append the letter to letterHistory
		}
	} else {
		wordHistory = append(wordHistory, choiceToLower) //Else, append the word to wordHistory
	}
	updateWord(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, choiceToLowerStrings)
}
