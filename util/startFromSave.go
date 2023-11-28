package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func StartFromSave(asciiMode string, pathAscii string) {
	var currentDir, _ = os.Getwd()

	type Gamestate struct {
		ArrSelectWord       []string `json:"ArrSelectWord"`
		WordPartiallyReveal []string `json:"WordPartiallyReveal"`
		LetterHistory       []string `json:"LetterHistory"`
		WordHistory         []string `json:"WordHistory"`
		LiveJose            int      `json:"LiveJose"`
	}

	var restart Gamestate

	file, _ := os.Open(currentDir + "\\resources\\save.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erreur lors de la fermeture du fichier de sauvegarde")
		}
	}(file)

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&restart)
	if err != nil {
		fmt.Println("Erreur lors du décodage du fichier de sauvegarde")
	}

	ClearTerminal()
	fmt.Println("")
	fmt.Println("Bon retour parmis nous, votre sauvegarde à préalablement été sauvegardé et est prête à être utilisé !")

	startGame(asciiMode, pathAscii, restart.ArrSelectWord, restart.WordPartiallyReveal, restart.LetterHistory, restart.WordHistory, restart.LiveJose)
}
