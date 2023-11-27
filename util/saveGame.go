package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func createGameSave(arrSelectWord []string, wordPartiallyReveal []string, liveJose int) {

	type Gamestate struct {
		LiveJose            int
		Wordtofind          []string
		WordPartiallyReveal []string
		LetterHistory       []string
		WordHistory         []string
	}

	saveGame := Gamestate{
		LiveJose:            liveJose,
		Wordtofind:          arrSelectWord,
		WordPartiallyReveal: wordPartiallyReveal,
		LetterHistory:       letterHistory,
		WordHistory:         wordHistory,
	}

	save, err1 := json.Marshal(saveGame)
	if err1 != nil {
		fmt.Println("Erreur lors de la sauvegarde de la partie")
	}

	err2 := os.WriteFile(currentDir+"\\resources\\save.txt", save, 0644)
	if err2 != nil {
		fmt.Println("Erreur lors de la sauvegarde de la partie")
		return
	}

	ClearTerminal()

	fmt.Println("Votre partie a été sauvegardé, à bientôt !")

	os.Exit(0)
}
