package main

import (
	"flag"
	"fmt"
	Util "hangman-classic/util"
	"os"
)

func main() {

	var (
		currentDir, _ = os.Getwd() // Get the current directory
		startWith     string
		asciiMode     string
		pathAscii     string
	)

	flag.StringVar(&asciiMode, "letterFile", "", "Select Ascii Mode")       // Flag to select the ascii mode
	flag.StringVar(&startWith, "startWith", "", "Start with the save file") // Flag to start with the save file
	flag.Parse()                                                            // Parse (read) the flags

	if asciiMode != "" && asciiMode != "standard.txt" && asciiMode != "shadow.txt" && asciiMode != "thinkertoy.txt" { // Check if the ascii mode is available
		fmt.Println("Le fichier ascii spécifié n'existe pas, merci de relancer le programme avec un fichier existant")
		return

	} else if asciiMode == "standard.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\standard.txt"
	} else if asciiMode == "shadow.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\shadow.txt" // Set the path of the ascii mode if is available
	} else if asciiMode == "thinkertoy.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\thinkertoy.txt"
	}

	if startWith != "" && startWith != "save.txt" { // Check if the save file is available
		fmt.Println("Le fichier de sauvegarde spécifié n'existe pas, merci de relancer le programme avec un fichier existant")
		return

	} else if startWith == "save.txt" {
		Util.StartFromSave(asciiMode, pathAscii) // Go in the function to start from the save file if is available
	} else {
		Util.ClearTerminal()
		Util.PrintRules(asciiMode, pathAscii) // Else, go in the function to print the rules
	}
}
