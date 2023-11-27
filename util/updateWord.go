package util

func updateWord(arrSelectWord []string, wordPartiallyReveal []string, choiceToLowerStrings []string) {
	letterFind := false
	wordFind := false
	if len(choiceToLowerStrings) == 1 {
		for index, letter := range arrSelectWord {
			if letter == choiceToLowerStrings[0] {
				wordPartiallyReveal[index] = letter
				letterFind = true
			}
		}
	} else {
		counter := 0
		for i := 0; i < len(arrSelectWord); i++ {
			if arrSelectWord[i] == choiceToLowerStrings[i] {
				counter++
			}
		}
		if counter == len(arrSelectWord) {
			wordPartiallyReveal = choiceToLowerStrings
			wordFind = true
		}
	}
	printWordPartiallyReveal(wordPartiallyReveal)
	histroy(choiceToLowerStrings, letterFind, wordFind, wordPartiallyReveal, arrSelectWord)
}
