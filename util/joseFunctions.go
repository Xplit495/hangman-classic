package util

import (
	"bufio"
	"fmt"
	"os"
)

func PrintJose(startLine int, endLine int) {
	currentDir, _ := os.Getwd()                             // Get current directory
	absolutePath := currentDir + "\\resources\\hangman.txt" // Absolute path to the file
	file, _ := os.Open(absolutePath)                        // Open the file
	scanner := bufio.NewScanner(file)                       // Create a scanner to read the file
	currentLine := 0                                        // Current line of the file
	for scanner.Scan() {
		currentLine++
		if currentLine >= startLine && currentLine <= endLine { // If the current line is between the start and end line
			fmt.Println(red + scanner.Text() + reset) // Print the line
		}
		if currentLine > endLine { // If the current line is greater than the end line
			err := file.Close() // Close the file
			if err != nil {
				fmt.Println("Erreur lors de la fermeture du fichier hangman")
				return
			}
			break
		}
	}
}

func ChooseLiveJose(liveJose int) {
	switch liveJose { // Choose the right case in function of the number of lives
	case 10:
		fmt.Print("")
	case 9:
		PrintJose(1, 7)
	case 8:
		PrintJose(8, 14)
	case 7:
		PrintJose(15, 22)
	case 6:
		PrintJose(23, 30)
	case 5:
		PrintJose(31, 38)
	case 4:
		PrintJose(39, 46)
	case 3:
		PrintJose(47, 54)
	case 2:
		PrintJose(55, 62)
	case 1:
		PrintJose(63, 70)
	}
}
