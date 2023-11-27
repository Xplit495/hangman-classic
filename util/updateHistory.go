package util

func updateHistroy(choiceToLowerStrings []string, letterFind bool, wordFind bool, wordPartiallyReveal []string, arrSelectWord []string) {
	counter := 0
	letterAlreadyUse := false
	wordAlreadyUse := false
	if len(choiceToLowerStrings) == 1 {
		for _, char := range letterHistory {
			for i := 0; i < len(choiceToLowerStrings); i++ {
				if choiceToLowerStrings[i] == char {
					counter++
					if counter > 1 {
						counter = 0
						letterAlreadyUse = true
						if len(letterHistory) > 0 {
							letterHistory = letterHistory[:len(letterHistory)-1]
						}
					}
				}
			}
		}
	} else {
		var wordTry string
		for i := 0; i < len(choiceToLowerStrings); i++ {
			wordTry = wordTry + choiceToLowerStrings[i]
		}
		if len(wordHistory) > 1 {
			for i := 0; i < len(wordHistory); i++ {
				if wordHistory[i] == wordTry {
					counter++
					if counter > 1 {
						counter = 0
						wordAlreadyUse = true
						if len(wordHistory) > 0 {
							wordHistory = wordHistory[:len(wordHistory)-1]
						}
					}
				}
			}
		}
	}
	checkElementUses(choiceToLowerStrings, letterFind, wordFind, letterAlreadyUse, wordPartiallyReveal, arrSelectWord, wordAlreadyUse)
}
