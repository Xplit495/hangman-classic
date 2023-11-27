package util

import (
	"bufio"
	"fmt"
	"os"
)

func printJose(startLine int, endLine int) {
	currentDir, _ := os.Getwd()
	absolutePath := currentDir + "\\resources\\hangman.txt"
	file, _ := os.Open(absolutePath)
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		currentLine++
		if currentLine >= startLine && currentLine <= endLine {
			fmt.Println(scanner.Text())
		}
		if currentLine > endLine {
			err := file.Close()
			if err != nil {
				fmt.Println("Erreur lors de la fermeture du fichier hangman")
				return
			}
			break
		}
	}
}

// A séparer dans des dossiers
func printLive(liveJose int) {
	switch liveJose {
	case 10:
		fmt.Print("")
	case 9:
		printJose(1, 7)
	case 8:
		printJose(8, 14)
	case 7:
		printJose(15, 22)
	case 6:
		printJose(23, 30)
	case 5:
		printJose(31, 38)
	case 4:
		printJose(39, 46)
	case 3:
		printJose(47, 54)
	case 2:
		printJose(55, 62)
	case 1:
		printJose(63, 70)
	}
}
