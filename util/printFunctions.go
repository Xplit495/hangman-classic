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
	yellow = "\033[33m"
	red    = "\033[31m"
	green  = "\033[32m"
	reset  = "\033[0m"
)

func printAscii(pathAscii string, wordToPrintInAscii []string) {
	wordPartiallyRevealString := strings.Join(wordToPrintInAscii, "")
	arrRune := []rune(wordPartiallyRevealString)
	for i := 0; i < 9; i++ {
		for j := 0; j < len(wordToPrintInAscii); j++ {
			startLine := int((arrRune[j] - 32) * 9)
			endLine := int(((arrRune[j] + 1) - 32) * 9)
			file, _ := os.Open(pathAscii)
			scanner := bufio.NewScanner(file)
			currentLine := 0
			for scanner.Scan() {
				currentLine++
				if currentLine == startLine+i+1 {
					fmt.Print(scanner.Text())
					fmt.Print("  ")
					break
				}
				if currentLine >= endLine {
					err := file.Close()
					if err != nil {
						fmt.Println("Erreur lors de la fermeture du fichier ascii")
						return
					}
					break
				}
			}
			_, err := file.Seek(0, 0)
			if err != nil {
				fmt.Println("Erreur lors de la remise du pointeur du fichier ascii")
				return
			}
		}
		fmt.Println()
	}
}

func printWordPartiallyReveal(asciiMode string, pathAscii string, wordPartiallyReveal []string) {
	if asciiMode != "" {
		printAscii(pathAscii, wordPartiallyReveal)
	} else {
		for i := 0; i < len(wordPartiallyReveal); i++ {
			fmt.Print(wordPartiallyReveal[i])
		}
		fmt.Println("")
	}
}

func printWord(asciiMode string, pathAscii string, arrSelectWord []string) {
	if asciiMode != "" {
		printAscii(pathAscii, arrSelectWord)
	} else {
		for i := 0; i < len(arrSelectWord); i++ {
			fmt.Print(arrSelectWord[i])
		}
	}
	fmt.Println("")
}

func printLetterHistory(letterHistory []string) {
	for i := 0; i <= len(letterHistory)-1; i++ {
		fmt.Print(letterHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

func printWordHistory(wordHistory []string) {
	for i := 0; i < len(wordHistory); i++ {
		fmt.Print(wordHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erreur lors de l'éxécution de la commande de nettoyage du terminal")
		return
	}
}

func PrintRules(asciiMode string, pathAscii string) {
	currentDir, _ := os.Getwd()
	pathAsciiForBeggin := currentDir + "\\resources\\ascii\\standard.txt"
	printAscii(pathAsciiForBeggin, []string{"H", "A", "N", "G", "M", "A", "N", "-", "C", "L", "A", "S", "S", "I", "C"})

	fmt.Println("")
	fmt.Println("Bienvenue dans ce super jeu, les régles sont simples :")
	fmt.Println("- Vous pouvez proposer ou un mot ou une lettre")
	fmt.Println("- Une mauvaise lettre vous retire" + yellow + " une " + reset + "vie. Mais " + red + "attention" + reset + " car un mauvais mot vous en retire" + yellow + " 2" + reset + " !")
	fmt.Print("Appuyer sur entrer pour continuer : ")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée standard")
		return
	}
	selectDifficulty(asciiMode, pathAscii)
}
