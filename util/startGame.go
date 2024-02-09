package util

import (
	"strings"
)

func StartGame(choice string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string) bool {
	var (
		choiceToLower        string
		choiceToLowerStrings []string
		choiceToLowerRune    []rune
		goodInput            bool
	)

	choiceToLowerStrings = nil
	choiceToLower = strings.ToLower(choice) //Convert choice to lowercase

	choiceToLowerRune = []rune(choiceToLower)                                              //Convert choiceToLower to rune
	if len(wordPartiallyReveal) == len(choiceToLowerRune) || len(choiceToLowerRune) == 1 { //If wordPartiallyReveal is equal to choiceToLower or choiceToLower is equal to 1 (if user input is a word of equal size than the word to find or if user input is a letter)
		for j := 0; j < len(choiceToLowerRune); j++ { //For each letter of choiceToLowerRune
			choiceToLowerStrings = append(choiceToLowerStrings, string(choiceToLowerRune[j])) //Append the letter to choiceToLowerStrings
		}

		goodInput = true
		for k := 0; k < len(choiceToLowerRune); k++ {
			if choiceToLowerRune[k] >= rune(97) && choiceToLowerRune[k] <= rune(122) { //Check if the letter is between a and z
				if k+1 == len(choiceToLowerRune) {
					break
				}
			} else { //If the letter is not between a and z prevent the user and ask him to enter only letters
				goodInput = false
			}
		}
	} else {
		goodInput = false
		return goodInput
	}

	if goodInput == true {
		if len(choiceToLowerStrings) == 1 { //If choiceToLowerStrings is equal to 1 (if user input is a letter)
			for i := 0; i < len(choiceToLowerStrings); i++ {
				letterHistory = append(letterHistory, choiceToLowerStrings[i]) //Append the letter to letterHistory
			}
		} else {
			wordHistory = append(wordHistory, choiceToLower) //Else, append the word to wordHistory
		}
	}
	return goodInput
}
