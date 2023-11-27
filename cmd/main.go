package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	Util "hangman-classic/util"
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

		ClearTerminal()
		fmt.Println("")
		fmt.Println("Bon retour parmis nous, votre sauvegarde à préalablement été sauvegardé et est prête à être utilisé !")

		startGame(restart.Wordtofind, restart.WordPartiallyReveal, liveJose)

	} else {
		ClearTerminal()
		Util.Rules()
	}
}

func startGame(arrSelectWord []string, wordPartiallyReveal []string, liveJose int) {
	var choiceToLowerStrings []string
	var choiceToLower string
	var choice string

	printLive(liveJose)
	fmt.Println("")
	fmt.Printf("Il vous reste "+yellow+"%d vie "+reset+"avant d'être pendu !\n", liveJose)
	fmt.Println("")
	if len(letterHistory) > 0 {
		fmt.Print("Les lettres déjà essayé sont : ")
		printLetterHistory()
	}
	if len(wordHistory) > 0 {
		fmt.Print("Les mots déjà essayé sont : ")
		printWordHistory()
	}

	for i := 0; i <= 1; i++ {
		choiceToLowerStrings = nil
		fmt.Print("Entrez votre lettre ou votre mot : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée standard")
			return
		}
		choiceToLower = strings.ToLower(choice)
		if choiceToLower == "stop" {

			type Gamestate struct {
				LiveJose            int
				Wordtofind          []string
				WordPartiallyReveal []string
				LetterHistory       []string
				WordHistory         []string
			}

			saveGame := Gamestate{
				LiveJose:            liveJose,
				Wordtofind:          arrSelectWord,
				WordPartiallyReveal: wordPartiallyReveal,
				LetterHistory:       letterHistory,
				WordHistory:         wordHistory,
			}

			save, err1 := json.Marshal(saveGame)
			if err1 != nil {
				fmt.Println("Erreur lors de la sauvegarde de la partie")
			}

			err2 := os.WriteFile(currentDir+"\\resources\\save.txt", save, 0644)
			if err2 != nil {
				fmt.Println("Erreur lors de la sauvegarde de la partie")
				return
			}

			ClearTerminal()

			fmt.Println("Votre partie a été sauvegardé, à bientôt !")

			os.Exit(0)

		}

		choiceToLowerRune = []rune(choiceToLower)
		if len(wordPartiallyReveal) == len(choiceToLowerRune) || len(choiceToLowerRune) == 1 {
			for j := 0; j < len(choiceToLowerRune); j++ {
				choiceToLowerStrings = append(choiceToLowerStrings, string(choiceToLowerRune[j]))
			}
			exit := true
			for k := 0; k < len(choiceToLowerRune); k++ {
				if choiceToLowerRune[k] >= rune(97) && choiceToLowerRune[k] <= rune(122) {
					if k+1 == len(choiceToLowerRune) {
						break
					}
				} else {
					ClearTerminal()
					fmt.Println("Merci de saisir" + red + " UNIQUEMENT " + reset + "des caractère de l'alphabet !")
					exit = false
					i--
				}
			}
			if exit == true {
				break
			}
		} else {
			ClearTerminal()
			fmt.Println("Merci de saisir " + red + "UNIQUEMENT " + reset + "une lettre ou un mot (de même longeur) !")
			i--
		}
	}
	if len(choiceToLowerStrings) == 1 {
		for i := 0; i < len(choiceToLowerStrings); i++ {
			letterHistory = append(letterHistory, choiceToLowerStrings[i])
		}
	} else {
		wordHistory = append(wordHistory, choiceToLower)
	}
	refreshWord(arrSelectWord, wordPartiallyReveal, choiceToLowerStrings)
}

func refreshWord(arrSelectWord []string, wordPartiallyReveal []string, choiceToLowerStrings []string) {
	letterFind := false
	wordFind := false
	if len(choiceToLowerStrings) == 1 {
		for index, letter := range arrSelectWord {
			if letter == choiceToLowerStrings[0] {
				wordPartiallyReveal[index] = letter
				letterFind = true
			}
		}
	} else {
		counter := 0
		for i := 0; i < len(arrSelectWord); i++ {
			if arrSelectWord[i] == choiceToLowerStrings[i] {
				counter++
			}
		}
		if counter == len(arrSelectWord) {
			wordPartiallyReveal = choiceToLowerStrings
			wordFind = true
		}
	}
	printWordPartiallyReveal(wordPartiallyReveal)
	histroy(choiceToLowerStrings, letterFind, wordFind, wordPartiallyReveal, arrSelectWord)
}

func histroy(choiceToLowerStrings []string, letterFind bool, wordFind bool, wordPartiallyReveal []string, arrSelectWord []string) {
	counter := 0
	letterAlreadyUse := false
	wordAlreadyUse := false
	if len(choiceToLowerStrings) == 1 {
		for _, char := range letterHistory {
			for i := 0; i < len(choiceToLowerStrings); i++ {
				if choiceToLowerStrings[i] == char {
					counter++
					if counter > 1 {
						counter = 0
						letterAlreadyUse = true
						if len(letterHistory) > 0 {
							letterHistory = letterHistory[:len(letterHistory)-1]
						}
					}
				}
			}
		}
	} else {
		var wordTry string
		for i := 0; i < len(choiceToLowerStrings); i++ {
			wordTry = wordTry + choiceToLowerStrings[i]
		}
		if len(wordHistory) > 1 {
			for i := 0; i < len(wordHistory); i++ {
				if wordHistory[i] == wordTry {
					counter++
					if counter > 1 {
						counter = 0
						wordAlreadyUse = true
						if len(wordHistory) > 0 {
							wordHistory = wordHistory[:len(wordHistory)-1]
						}
					}
				}
			}
		}
	}
	checkElementUses(choiceToLowerStrings, letterFind, wordFind, letterAlreadyUse, wordPartiallyReveal, arrSelectWord, wordAlreadyUse)
}

func checkElementUses(choiceToLowerStrings []string, letterFind bool, wordFind bool, letterAlreadyUse bool, wordPartiallyReveal []string, arrSelectWord []string, wordAlreadyUse bool) {
	if letterAlreadyUse == true {
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé cette lettre, attention !" + reset)
		startGame(arrSelectWord, wordPartiallyReveal, liveJose)
	} else if wordAlreadyUse == true {
		ClearTerminal()
		fmt.Println(red + "Vous avez déjà essayé ce mot, attention !" + reset)
		startGame(arrSelectWord, wordPartiallyReveal, liveJose)
	}
	fmt.Println("")
	findLetter(choiceToLowerStrings, wordPartiallyReveal, arrSelectWord, letterFind, wordFind)
}

func findLetter(choiceToLowerStrings []string, wordPartiallyReveal []string, arrSelectWord []string, letterFind bool, wordFind bool) {
	if len(choiceToLowerStrings) == 1 {
		if letterFind == true {
			ClearTerminal()
			fmt.Println(green + "Bonne lettre !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		} else if letterFind == false {
			liveJose--
			ClearTerminal()
			fmt.Println(red + "Mauvaise lettre !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		}
	} else {
		if wordFind == false {
			liveJose = liveJose - 2
			ClearTerminal()
			fmt.Println(red + "Mauvais mot !" + reset)
			fmt.Println("")
			fmt.Printf("Pour le moment le mot ressemble à ca -> ")
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		} else if wordFind == true {
			ClearTerminal()
			printWordPartiallyReveal(wordPartiallyReveal)
			checkWordFind(wordPartiallyReveal, arrSelectWord)
		}
	}
}

func checkWordFind(wordPartiallyReveal []string, arrSelectWord []string) {
	wordFind := true
	for _, letter := range wordPartiallyReveal {
		if letter == "_" {
			wordFind = false
			break
		}
	}
	if wordFind == true {
		ClearTerminal()
		fmt.Println("\n" + green + "Vous avez deviné le mot !" + reset)
		if len(letterHistory) > 0 {
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory()
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory()
		}
		fmt.Print("Le mot était : ")
		printWord(arrSelectWord)
		fmt.Println("")
	} else if liveJose <= 0 {
		ClearTerminal()
		fmt.Print("\n" + red + "Vous n'avez plus de vie !" + reset + "\nLe mot était : ")
		printWord(arrSelectWord)
		printJose(71, 78)
		if len(letterHistory) > 0 {
			fmt.Println("")
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistory()
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory()
		}
		fmt.Println("")
		fmt.Println(red + "Vous êtes pendu !" + reset)
	} else {
		startGame(arrSelectWord, wordPartiallyReveal, liveJose)
	}
}

// Début des fonctions d'affichage

// A séparer dans des dossiers

//Fin des fonctions d'affichage

// Debut fonction position Jose
func printJose(startLine int, endLine int) {
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

//Fin fonction position Jose
