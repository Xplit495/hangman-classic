package util

func updateHistroy(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, choiceToLowerStrings []string, letterFind bool, wordFind bool) {
	var (
		counter          = 0
		letterAlreadyUse = false
		wordAlreadyUse   = false
		wordTry          string
	)

	if len(choiceToLowerStrings) == 1 { // if the user input is a letter
		for _, char := range letterHistory { //Check every element of the letterHistory array
			for i := 0; i < len(choiceToLowerStrings); i++ {
				if choiceToLowerStrings[i] == char { // if the user input is already in the letterHistory array
					counter++ // increment the counter
					if counter > 1 {
						counter = 0                 // reset the counter to avoid a bug
						letterAlreadyUse = true     // set the letterAlreadyUse to true
						if len(letterHistory) > 0 { // if the letterHistory array is not empty
							letterHistory = letterHistory[:len(letterHistory)-1] // remove the last element of the letterHistory array
						}
					}
				}
			}
		}
	} else {
		for i := 0; i < len(choiceToLowerStrings); i++ {
			wordTry = wordTry + choiceToLowerStrings[i] //WordTry is an empty string, we add the user input to it to compare it with the wordHistory array
		}
		if len(wordHistory) > 1 { // if the wordHistory array is not empty
			for i := 0; i < len(wordHistory); i++ {
				if wordHistory[i] == wordTry { // if the user input is already in the wordHistory array
					counter++ // increment the counter
					if counter > 1 {
						counter = 0               // reset the counter to avoid a bug
						wordAlreadyUse = true     // set the wordAlreadyUse to true
						if len(wordHistory) > 0 { // if the wordHistory array is not empty
							wordHistory = wordHistory[:len(wordHistory)-1] // remove the last element of the wordHistory array
						}
					}
				}
			}
		}
	}
	checkInputAlreadyUses(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, choiceToLowerStrings, letterFind, wordFind, letterAlreadyUse, wordAlreadyUse)
}
