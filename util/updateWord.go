package util

func updateWord(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, choiceToLowerStrings []string) {
	var (
		letterFind = false
		wordFind   = false
	)

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
	printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
	updateHistroy(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, choiceToLowerStrings, letterFind, wordFind)
}
