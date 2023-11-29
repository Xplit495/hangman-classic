package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func createGameSave(arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int) {

	type Gamestate struct { //Create a struct to save the game
		ArrSelectWord       []string
		WordPartiallyReveal []string
		LetterHistory       []string
		WordHistory         []string
		LiveJose            int
	}

	saveGame := Gamestate{ //Create a variable to store the gamestate
		ArrSelectWord:       arrSelectWord,
		WordPartiallyReveal: wordPartiallyReveal,
		LetterHistory:       letterHistory,
		WordHistory:         wordHistory,
		LiveJose:            liveJose,
	}

	save, err1 := json.Marshal(saveGame) //Convert the gamestate to json.Marshal
	if err1 != nil {
		fmt.Println("Erreur lors de la sauvegarde de la partie")
	}

	currentDir, _ := os.Getwd()                                          //Get the current directory
	err2 := os.WriteFile(currentDir+"\\resources\\save.txt", save, 0644) //Write the save in the file save.txt
	if err2 != nil {
		fmt.Println("Erreur lors de la sauvegarde de la partie")
		return
	}

	ClearTerminal()

	fmt.Println("Votre partie a été sauvegardé, à bientôt !")

	os.Exit(0) //Exit the program
}
