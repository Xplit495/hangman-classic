package util

func UpdateWord(arrSelectWord []string, wordPartiallyReveal []string, choiceToLowerStrings []string) ([]string, bool, bool) {
	var (
		letterFind = false
		wordFind   = false
	)

	if len(choiceToLowerStrings) == 1 { //If user input is a letter
		for index, letter := range arrSelectWord {
			if letter == choiceToLowerStrings[0] { //If letter of user input is equal to a letter of the word
				wordPartiallyReveal[index] = letter //Replace the "_" of wordPartiallyReveal by the good letter
				letterFind = true                   //The letter is find
			}
		}
	} else { //If user input is a word
		counter := 0
		for i := 0; i < len(arrSelectWord); i++ {
			if arrSelectWord[i] == choiceToLowerStrings[i] { //If letter of index[i] of the user input is equal to the letter of index[i] counter increase
				counter++ //Necessary to know if every letter of the word is find
			}
		}
		if counter == len(arrSelectWord) { //If counter is equal to the length of the word, the word is find
			wordPartiallyReveal = choiceToLowerStrings //So, replace the "_" of wordPartiallyReveal by the word
			wordFind = true
		}
	}

	return wordPartiallyReveal, letterFind, wordFind
}
