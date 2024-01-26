package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	// ANSI escape sequence for gave text color (as they are defined in the util package, we can use them in all the other .go files into util package)
	yellow = "\033[33m"
	red    = "\033[31m"
	green  = "\033[32m"
	reset  = "\033[0m"
)

func printAscii(pathAscii string, wordToPrintInAscii []string) {
	wordToPrintInAsciiString := strings.Join(wordToPrintInAscii, "") // Convert the array of string into a string because cannot convert directly an array of string into an array of rune
	arrRune := []rune(wordToPrintInAsciiString)                      // Convert the string into an array of rune
	for i := 0; i < 9; i++ {
		for j := 0; j < len(wordToPrintInAscii); j++ {
			startLine := int((arrRune[j] - 32) * 9)     //Algorithm to find the start line of the letter in the ascii file
			endLine := int(((arrRune[j] + 1) - 32) * 9) //Algorithm to find the end line of the letter in the ascii file (beginning of the next letter)
			file, _ := os.Open(pathAscii)               // Open the ascii file
			scanner := bufio.NewScanner(file)           // Create a scanner to read the file
			currentLine := 0                            // Variable to know the current line of the file
			for scanner.Scan() {                        // Loop to read the file
				currentLine++
				if currentLine == startLine+i+1 { //This condition is here for print every line of every letter line by line (the + 1 is for ignore the first line of the letter because with this algorithm the line where beggin the start line is a blank line)(1/3)
					fmt.Print(scanner.Text()) //Because if we print directly an entire letter we can't go back to the good emplacement for well print the next letter (2/3)
					fmt.Print("  ")           //For example,  AB not going to be print like this : A B (horizontal) but like this : A (vertical) B (vertical) (3/3)
					break
				}
				if currentLine >= endLine { //Is here for close the file to prepare to pass to the next letter
					err := file.Close()
					if err != nil {
						fmt.Println("Erreur lors de la fermeture du fichier ascii")
						return
					}
					break
				}
			}
			_, err := file.Seek(0, 0) //Return to the beginning of the file
			if err != nil {
				fmt.Println("Erreur lors de la remise du pointeur du fichier ascii")
				return
			}
		}
		fmt.Println() //When the loop i is finished, we return to the line for print the next line of the letter
	}
}

func printWordPartiallyReveal(asciiMode string, pathAscii string, wordPartiallyReveal []string) {
	if asciiMode != "" { //If the user want to play with ascii mode we call the function printAscii
		printAscii(pathAscii, wordPartiallyReveal)
	} else {
		for i := 0; i < len(wordPartiallyReveal); i++ { //Else we print the word normally (with a loop cause the word is an array of string)
			fmt.Print(wordPartiallyReveal[i])
		}
		fmt.Println("")
	}
}

func printWord(asciiMode string, pathAscii string, arrSelectWord []string) {
	if asciiMode != "" { //If the user want to play with ascii mode we call the function printAscii
		printAscii(pathAscii, arrSelectWord)
	} else {
		for i := 0; i < len(arrSelectWord); i++ { //Else we print the word normally (with a loop cause the word is an array of string)
			fmt.Print(arrSelectWord[i])
		}
	}
	fmt.Println("")
}

func printLetterHistory(letterHistory []string) {
	for i := 0; i <= len(letterHistory)-1; i++ { //We print the letter history with a loop cause the letter history is an array of string
		fmt.Print(letterHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

func printWordHistory(wordHistory []string) {
	for i := 0; i < len(wordHistory); i++ { //We print the word history with a loop cause the word history is an array of string
		fmt.Print(wordHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" { //If the user is on windows we call the command cls
		cmd = exec.Command("cmd", "/c", "cls") //The command is cmd /c cls because the command cls is a command of cmd
	} else { //Else we call the command clear
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout //We set the output of the command to the output of the terminal
	err := cmd.Run()       //We execute the command
	if err != nil {
		fmt.Println("Erreur lors de l'éxécution de la commande de nettoyage du terminal")
		return
	}
}

func PrintRules(asciiMode string, pathAscii string) {
	currentDir, _ := os.Getwd()                                           //Get the current directory
	pathAsciiForBeggin := currentDir + "\\resources\\ascii\\standard.txt" //Create the path for the ascii file
	printAscii(pathAsciiForBeggin, []string{"H", "A", "N", "G", "M", "A", "N", "-", "C", "L", "A", "S", "S", "I", "C"})

	fmt.Println("")
	fmt.Println("Bienvenue dans ce super jeu, les régles sont simples :")
	fmt.Println("- Vous pouvez proposer ou un mot ou une lettre")
	fmt.Println("- Une mauvaise lettre vous retire" + yellow + " une " + reset + "vie. Mais " + red + "attention" + reset + " car un mauvais mot vous en retire" + yellow + " 2" + reset + " !")
	fmt.Print("Appuyer sur entrer pour continuer : ")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n') //Wait for the user press enter
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée standard")
		return
	}
	selectDifficulty(asciiMode, pathAscii)
}
