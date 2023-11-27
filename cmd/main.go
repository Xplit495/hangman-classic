package main

import (
	"encoding/json"
	"flag"
	"fmt"
	Util "hangman-classic/util"
	"os"
)

var yellow = "\033[33m"
var red = "\033[31m"
var green = "\033[32m"
var reset = "\033[0m"

var liveJose = 10
var choiceToLowerRune []rune
var letterHistory []string
var wordHistory []string
var currentDir, _ = os.Getwd()
var startWith string
var asciiMode string
var pathAscii string

func main() {

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
			LiveJose            int      `json:"LiveJose"`
			Wordtofind          []string `json:"Wordtofind"`
			WordPartiallyReveal []string `json:"WordPartiallyReveal"`
			LetterHistory       []string `json:"LetterHistory"`
			WordHistory         []string `json:"WordHistory"`
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

		letterHistory = restart.LetterHistory
		wordHistory = restart.WordHistory
		liveJose = restart.LiveJose

		Util.ClearTerminal()
		fmt.Println("")
		fmt.Println("Bon retour parmis nous, votre sauvegarde à préalablement été sauvegardé et est prête à être utilisé !")

		Util.StartGame(restart.Wordtofind, restart.WordPartiallyReveal, liveJose)

	} else {
		Util.ClearTerminal()
		Util.PrintRules()
	}
}
