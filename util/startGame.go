package util

import (
	"fmt"
	"strings"
)

func StartGame(asciiMode string, pathAscii string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {
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
	if len(letterHistory) > 0 {
		fmt.Print("Les lettres déjà essayé sont : ")
		printLetterHistory(letterHistory)
	}
	if len(wordHistory) > 0 {
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
		choiceToLower = strings.ToLower(choice)

		if choiceToLower == "stop" {
			createGameSave(arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
		}

		choiceToLowerRune = []rune(choiceToLower)
		if len(wordPartiallyReveal) == len(choiceToLowerRune) || len(choiceToLowerRune) == 1 {
			for j := 0; j < len(choiceToLowerRune); j++ {
				choiceToLowerStrings = append(choiceToLowerStrings, string(choiceToLowerRune[j]))
			}
			exit := true
			for k := 0; k < len(choiceToLowerRune); k++ {
				if choiceToLowerRune[k] >= rune(97) && choiceToLowerRune[k] <= rune(122) {
					if k+1 == len(choiceToLowerRune) {
						break
					}
				} else {
					ClearTerminal()
					fmt.Println("Merci de saisir" + red + " UNIQUEMENT " + reset + "des caractère de l'alphabet !")
					exit = false
					i--
				}
			}
			if exit == true {
				break
			}
		} else {
			ClearTerminal()
			fmt.Println("Merci de saisir " + red + "UNIQUEMENT " + reset + "une lettre ou un mot (de même longeur) !")
			i--
		}
	}
	if len(choiceToLowerStrings) == 1 {
		for i := 0; i < len(choiceToLowerStrings); i++ {
			letterHistory = append(letterHistory, choiceToLowerStrings[i])
		}
	} else {
		wordHistory = append(wordHistory, choiceToLower)
	}
	updateWord(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose, choiceToLowerStrings)
}
