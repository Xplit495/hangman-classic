package main

import (
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

	} else if asciiMode == "standard.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\standard.txt"
	} else if asciiMode == "shadow.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\shadow.txt"
	} else if asciiMode == "thinkertoy.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\thinkertoy.txt"
	}

	if startWith != "" && startWith != "save.txt" {
		fmt.Println("Le fichier de sauvegarde spécifié n'existe pas, merci de relancer le programme avec un fichier existant")
		return

	} else if startWith == "save.txt" {
		Util.StartFromSave(asciiMode, pathAscii)
	} else {
		Util.ClearTerminal()
		Util.PrintRules(asciiMode, pathAscii)
	}
}
