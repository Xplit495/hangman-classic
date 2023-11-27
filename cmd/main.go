package main

import (
	"encoding/json"
	"flag"
	"fmt"
	Util "hangman-classic/util"
	"os"
)

func main() {
	var (
		currentDir, _ = os.Getwd()
		startWith     string
		asciiMode     string
		pathAscii     string
	)

	flag.StringVar(&asciiMode, "letterFile", "", "Select Ascii Mode")
	flag.StringVar(&startWith, "startWith", "", "Start with the save file")

	flag.Parse()

	if asciiMode != "" && asciiMode != "standard.txt" && asciiMode != "shadow.txt" && asciiMode != "thinkertoy.txt" {
		fmt.Println("Le fichier ascii spécifié n'existe pas, merci de relancer le programme avec un fichier existant")
		return
	}

	if asciiMode == "standard.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\standard.txt"
	} else if asciiMode == "shadow.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\shadow.txt"
	} else if asciiMode == "thinkertoy.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\thinkertoy.txt"
	}

	if startWith != "" && startWith != "save.txt" {
		fmt.Println("Le fichier de sauvegarde spécifié n'existe pas, merci de relancer le programme avec un fichier existant")
		return
	}
	if startWith == "save.txt" {

		type Gamestate struct {
			ArrSelectWord       []string `json:"Wordtofind"`
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

		Util.ClearTerminal()
		fmt.Println("")
		fmt.Println("Bon retour parmis nous, votre sauvegarde à préalablement été sauvegardé et est prête à être utilisé !")

		Util.StartGame(asciiMode, pathAscii, restart.ArrSelectWord, restart.WordPartiallyReveal, restart.LetterHistory, restart.WordHistory, restart.LiveJose)

	} else {
		Util.ClearTerminal()
		Util.PrintRules(asciiMode, pathAscii)
	}
}
