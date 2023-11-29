package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func StartFromSave(asciiMode string, pathAscii string) {
	var currentDir, _ = os.Getwd() // Get the current directory

	type Gamestate struct { // Create a struct to store the gamestate of the save file
		ArrSelectWord       []string `json:"ArrSelectWord"`
		WordPartiallyReveal []string `json:"WordPartiallyReveal"`
		LetterHistory       []string `json:"LetterHistory"`
		WordHistory         []string `json:"WordHistory"`
		LiveJose            int      `json:"LiveJose"`
	}

	var restart Gamestate // Create a variable to store the gamestate of the save file

	file, _ := os.Open(currentDir + "\\resources\\save.txt") // Open the save file
	defer func(file *os.File) {                              // Defer the closing of the save file
		err := file.Close()
		if err != nil {
			fmt.Println("Erreur lors de la fermeture du fichier de sauvegarde")
		}
	}(file)

	decoder := json.NewDecoder(file) // Create a decoder to decode json of  the save file

	err := decoder.Decode(&restart) // Decode the save file
	if err != nil {
		fmt.Println("Erreur lors du décodage du fichier de sauvegarde")
	}

	ClearTerminal()
	fmt.Println("")
	fmt.Println("Bon retour parmis nous, votre sauvegarde à préalablement été sauvegardé et est prête à être utilisé !")

	startGame(asciiMode, pathAscii, restart.ArrSelectWord, restart.WordPartiallyReveal, restart.LetterHistory, restart.WordHistory, restart.LiveJose) // Start the game with the gamestate of the save file
}
