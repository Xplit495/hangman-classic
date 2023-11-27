package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

var yellow = "\033[33m"
var red = "\033[31m"
var green = "\033[32m"
var reset = "\033[0m"

var liveJose = 10
var choiceToLowerRune []rune
var letterHistory []string
var letterHistoryEnd []string
var wordHistory []string
var currentDir, _ = os.Getwd()
var startWith string
var asciiMode string
var pathAscii string

func main() {

	flag.StringVar(&asciiMode, "letterFile", "", "Select Ascii Mode")
	flag.StringVar(&startWith, "startWith", "", "Start with the save file")

	flag.Parse()

	if asciiMode == "standard.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\standard.txt"
	} else if asciiMode == "shadow.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\shadow.txt"
	} else if asciiMode == "thinkertoy.txt" {
		pathAscii = currentDir + "\\resources\\ascii\\thinkertoy.txt"
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

		file, _ := os.Open("\\resources\\save.txt")
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
		rules()
	}
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

func rules() {
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
	chooseDifficulty()
}

func chooseDifficulty() {
	var difficulty int
	for i := 0; i <= 1; i++ {
		ClearTerminal()
		fmt.Println("")
		fmt.Print("Choissisez votre niveau de difficulté (1-3), 1: Facile, 2: Moyen, 3: Difficile. Que choissisez-vous : ")
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée standard")
			return
		}
		if difficulty != 1 && difficulty != 2 && difficulty != 3 {
			i--
		} else {
			break
		}
	}
	selectDictionnary(difficulty)
}

func selectDictionnary(difficulty int) {
	switch difficulty {
	case 1:
		absolutePath := currentDir + "\\resources\\dictionnary\\words.txt"
		selectRandomWordIntoDictionnary(absolutePath)
	case 2:
		absolutePath := currentDir + "\\resources\\dictionnary\\words2.txt"
		selectRandomWordIntoDictionnary(absolutePath)
	case 3:
		absolutePath := currentDir + "\\resources\\dictionnary\\words3.txt"
		selectRandomWordIntoDictionnary(absolutePath)
	}
}

func selectRandomWordIntoDictionnary(absolutePath string) {
	var (
		arrSelectWord   []string
		word            string
		numberOfWords   int
		indexRandomWord int
	)

	f, _ := os.Open(absolutePath)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numberOfWords++
	}
	err := f.Close()
	if err != nil {
		fmt.Println("Erreur lors de la fermeture du fichier de dictionnaire")
		return
	}
	indexRandomWord = rand.Intn(numberOfWords)

	currentLine := 0
	f2, _ := os.Open(absolutePath)
	scanner2 := bufio.NewScanner(f2)
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		currentLine++
		if currentLine == indexRandomWord {
			word = scanner2.Text()
			break
		}
	}
	err2 := f2.Close()
	if err2 != nil {
		fmt.Println("Erreur lors de la fermeture du fichier de dictionnaire")
		return
	}
	arrSelectWord = strings.Split(word, "")

	findWordClue(arrSelectWord)
}

func findWordClue(arrSelectWord []string) {
	var (
		randomClues []int
		n           = (len(arrSelectWord) / 2) - 1
	)

	usedClues := make(map[int]bool)
	for i := 1; i <= n; i++ {
		var newClue int
		for {
			newClue = rand.Intn(len(arrSelectWord) - 1)
			if !usedClues[newClue] {
				usedClues[newClue] = true
				break
			}
		}
		randomClues = append(randomClues, newClue)
	}
	sort.Ints(randomClues)

	associateClueToWord(randomClues, arrSelectWord)
}

func associateClueToWord(randomClues []int, arrSelectWord []string) {
	var (
		values              = 0
		wordPartiallyReveal []string
	)

	if len(randomClues) == 0 {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			wordPartiallyReveal = append(wordPartiallyReveal, "_")
		}
	} else {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			if i == randomClues[values] { // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
				wordPartiallyReveal = append(wordPartiallyReveal, arrSelectWord[i])
				if values+1 >= len(randomClues) {
					values = 0
				} else {
					values += 1
				}
			} else {
				wordPartiallyReveal = append(wordPartiallyReveal, "_")
			} // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
		}
	}

	fmt.Println("")
	fmt.Print("\nLe mot avec le(s) indice(s) est : ")
	printWordPartiallyReveal(wordPartiallyReveal)
	fmt.Println("")
	startGame(arrSelectWord, wordPartiallyReveal, 10)
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
		printLetterHistoryInGame()
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

			err2 := os.WriteFile("save.txt", save, 0644)
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
			letterHistoryEnd = append(letterHistoryEnd, choiceToLowerStrings[i])
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
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistoryInGame()
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory()
		}
		fmt.Print("Le mot était : ")
		printWordPartiallyReveal(wordPartiallyReveal)
		fmt.Println("")
	} else if liveJose <= 0 {
		ClearTerminal()
		fmt.Print("\n" + red + "Vous n'avez plus de vie !" + reset + "\nLe mot était : ")
		printWordPartiallyReveal(wordPartiallyReveal)
		fmt.Println("")
		fmt.Println("")
		printJose(71, 78)
		if len(letterHistory) > 0 {
			fmt.Print("Les lettres essayés ont été : ")
			printLetterHistoryInGame()
		}
		if len(wordHistory) > 0 {
			fmt.Print("Les mots essayés ont été : ")
			printWordHistory()
		}
		fmt.Println(red + "Vous êtes pendu !" + reset)
	} else {
		startGame(arrSelectWord, wordPartiallyReveal, liveJose)
	}
}

// Début des fonctions d'affichage
func printWordPartiallyReveal(wordPartiallyReveal []string) {
	wordPartiallyRevealString := strings.Join(wordPartiallyReveal, "")
	arrRune := []rune(wordPartiallyRevealString)
	if asciiMode != "" {
		for i := 0; i < 9; i++ {
			for j := 0; j < len(wordPartiallyReveal); j++ {
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
	} else {
		for i := 0; i < len(wordPartiallyReveal); i++ {
			fmt.Print(wordPartiallyReveal[i])
		}
		fmt.Println("")
	}
}

func printLetterHistoryInGame() {
	for i := 0; i <= len(letterHistory)-1; i++ {
		fmt.Print(letterHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

func printWordHistory() {
	for i := 0; i < len(wordHistory); i++ {
		fmt.Print(wordHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

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