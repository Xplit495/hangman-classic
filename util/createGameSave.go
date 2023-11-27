package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func createGameSave(arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {

	type Gamestate struct {
		ArrSelectWord       []string
		WordPartiallyReveal []string
		LetterHistory       []string
		WordHistory         []string
		LiveJose            int
	}

	saveGame := Gamestate{
		ArrSelectWord:       arrSelectWord,
		WordPartiallyReveal: wordPartiallyReveal,
		LetterHistory:       letterHistory,
		WordHistory:         wordHistory,
		LiveJose:            liveJose,
	}

	save, err1 := json.Marshal(saveGame)
	if err1 != nil {
		fmt.Println("Erreur lors de la sauvegarde de la partie")
	}

	currentDir, _ := os.Getwd()
	err2 := os.WriteFile(currentDir+"\\resources\\save.txt", save, 0644)
	if err2 != nil {
		fmt.Println("Erreur lors de la sauvegarde de la partie")
		return
	}

	ClearTerminal()

	fmt.Println("Votre partie a été sauvegardé, à bientôt !")

	os.Exit(0)
}
