package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

var liveJose = 10
var letterHistory []string
var letterHistoryEnd []string

func main() {
	chooseDifficulty()
}

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func chooseDifficulty(){
var difficulty int // CHOISIR NIVEAU DIFFICULTE (POUR LE MOMENT UTILISER SEULEMENT DIIFUCLTE 1 ON CODERA LES AUTRES PLUS TARD)
	for i := 0; i <= 1; i++ {
		ClearTerminal()
		fmt.Println("")
		fmt.Print("Choose your level of difficulty (1-3), 1: Easy, 2: Medium, 3: Hard. What you choose : ") // CHOISIR NIVEAU DIFFICULTE (POUR LE MOMENT UTILISER SEULEMENT DIIFUCLTE 1 ON CODERA LES AUTRES PLUS TARD)
		fmt.Scanln(&difficulty)
		if difficulty != 1 && difficulty != 2 && difficulty != 3 {
			i--
		} else {
			break
		} // CHOISIR NIVEAU DIFFICULTE (POUR LE MOMENT UTILISER SEULEMENT DIIFUCLTE 1 ON CODERA LES AUTRES PLUS TARD)
	}
	selectDictionnary(difficulty)
}

func selectDictionnary(difficulty int){
	switch difficulty {
	case 1:
		f, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		wordList := []string{}
		for scanner.Scan() {
			wordList = append(wordList, scanner.Text())
		}
		randomWord(wordList)
	case 2:
		f, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words2.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		wordList := []string{}
		for scanner.Scan() {
			wordList = append(wordList, scanner.Text())
		}
		randomWord(wordList)
	case 3:
		f, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words3.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		wordList := []string{}
		for scanner.Scan() {
			wordList = append(wordList, scanner.Text())
		}
		randomWord(wordList)
	}

}

func randomWord(wordList []string){
	indexRandomWord := rand.Intn(len(wordList)) - 1 // SELECTIONNE UN MOT AU HASARD DANS LE DICTIONNAIRE ET LE MET DANS UN TABLEAU
	if indexRandomWord <= 0 {
		indexRandomWord += 1
	}
	selectWord := wordList[indexRandomWord]
	var arrSelectWord []string
	for i := 0; i < len(selectWord); i++ {
		arrSelectWord = append(arrSelectWord, string(selectWord[i]))
	} // SELECTIONNE UN MOT AU HASARD DANS LE DICTIONNAIRE ET LE MET DANS UN TABLEAU
	fmt.Println(arrSelectWord)
	findWordClue(arrSelectWord)
}

func findWordClue(arrSelectWord []string){
	n := (len(arrSelectWord) / 2) - 1 //LIS LA CONSIGNE POUR COMPRENDRE
	var randomClues []int

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
	fmt.Print(randomClues)
	associateClueToWord(randomClues, arrSelectWord)
}

func associateClueToWord(randomClues []int, arrSelectWord []string){
	values := 0 // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
	var wordPartiallyReveal []string
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
			}else {
				wordPartiallyReveal = append(wordPartiallyReveal, "_")
			} // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
		}
	}

	fmt.Println("")
	ClearTerminal()
	fmt.Print("\nLe mot avec le(s) indice(s) est : ")
	showWordPartiallyReveal(wordPartiallyReveal)
	fmt.Println("")
	startGame(arrSelectWord,wordPartiallyReveal,10)
}

func printLive(liveJose int){
	file, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\hangman.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentline := 0
	switch liveJose {
		case 10:
		fmt.Print("")
		case 9:
			startline := 1
			endline := 7
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					file.Close()
					break
				}
			}
		case 8:
			startline := 8
			endline := 14
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 7:
			startline := 15
			endline := 22
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 6:
			startline := 23
			endline := 30
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 5:
			startline := 31
			endline := 38
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 4:
			startline := 39
			endline := 46
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 3:
			startline := 47
			endline := 54
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 2:
			startline := 55
			endline := 62
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 1:
			startline := 63
			endline := 70
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
		case 0:
			startline := 71
			endline := 78
			for scanner.Scan() {
				currentline++
				if currentline >= startline && currentline <= endline {
					fmt.Println(scanner.Text())
				}
				if currentline > endline {
					break
				}
			}
	}
}


func startGame(arrSelectWord []string, wordPartiallyReveal [] string, liveJose int){
		printLive(liveJose)
		fmt.Println("")
		fmt.Printf("Il vous reste %d vie avant d'être pendu !\n", liveJose)
		var choice string
		var choiceToLower string
		for i := 0; i <= 1; i++ {
			fmt.Print("Entrez votre lettre : ")
			fmt.Scan(&choice)
			choiceToLower = strings.ToLower(choice)
			choiceToLowerRune := []rune(choiceToLower)
			if len(choiceToLowerRune) > 1 || (choiceToLowerRune[0] >= rune(0) && choiceToLowerRune[0] <= rune(96) || (choiceToLowerRune[0] > rune(122)))  {
				ClearTerminal()
				fmt.Println("Merci de saisir UN seul caractère de l'alphabet !")
				i--
			} else {
				break
			}
		}
		letterHistory = append(letterHistory, choiceToLower)
		letterHistoryEnd = append(letterHistoryEnd, choiceToLower)
		refreshWord(arrSelectWord,choiceToLower,wordPartiallyReveal,letterHistory,letterHistoryEnd)
}

func refreshWord(arrSelectWord []string, choiceToLower string, wordPartiallyReveal []string, letterHistory []string, letterHistoryEnd []string) {
	letterFind := false
	for index, letter := range arrSelectWord {
		if letter == choiceToLower {
			wordPartiallyReveal[index] = letter
			letterFind = true
		}
	}
	histroy(letterHistory, choiceToLower, letterFind ,wordPartiallyReveal,letterHistoryEnd,arrSelectWord)
}

func histroy (letterHistory []string, choiceToLower string, letterFind bool, wordPartiallyReveal []string, letterHistoryEnd []string, arrSelectWord []string){
	counter := 0
	letterAlreadyUse := false
	for _, char := range letterHistory {
		if choiceToLower == char {
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
	checkLetterUse(letterFind, letterAlreadyUse, wordPartiallyReveal, letterHistory, letterHistoryEnd, arrSelectWord)
}

func checkLetterUse(letterFind bool, alreadyUse bool, wordPartiallyReveal []string,letterHistory []string,letterHistoryEnd []string, arrSelectWord []string) {
	if letterFind == true && alreadyUse == true {
			ClearTerminal()
			fmt.Println("VOUS AVEZ DEJA ESSAYEZ CETTE LETTRE ! FAITES ATTENTION.")
	}
	if letterFind == true{
		findLetterYes(wordPartiallyReveal, letterHistory, letterHistoryEnd, arrSelectWord)
	}else{
		findLetterNo(wordPartiallyReveal, letterHistory, letterHistoryEnd, arrSelectWord)
	}
}

func findLetterYes(wordPartiallyReveal []string, letterHistory []string, letterHistoryEnd []string,arrSelectWord []string){
	fmt.Println("")
	ClearTerminal()
	fmt.Println("Bonne lettre !")
	fmt.Println("")
	fmt.Printf("Pour le moment le mot ressemble à ca -> ")
	showWordPartiallyReveal(wordPartiallyReveal)
	fmt.Print("Les lettres déjà essayé sont : ")
	showLetterHistoryInGame(letterHistory)
	checkWordFind(wordPartiallyReveal, letterHistoryEnd,arrSelectWord)
}

func findLetterNo(wordPartiallyReveal []string, letterHistory []string, letterHistoryEnd []string,arrSelectWord []string){
	liveJose--
	ClearTerminal()
	fmt.Println("Mauvaise lettre.")
	fmt.Println("")
	fmt.Printf("Pour le moment le mot ressemble à ca -> ")
	showWordPartiallyReveal(wordPartiallyReveal)
	fmt.Print("Les lettres déjà essayé sont : ")
	showLetterHistoryInGame(letterHistory)
	checkWordFind(wordPartiallyReveal, letterHistoryEnd,arrSelectWord)
}

func showWordPartiallyReveal(wordPartiallyReveal []string){
	for i := 0; i < len(wordPartiallyReveal); i++ {
		fmt.Print(wordPartiallyReveal[i])
	}
	fmt.Println("")
}

func showLetterHistoryInGame(letterHistory []string){
	for i := 0; i <= len(letterHistory)-1; i++ {
		fmt.Print(letterHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
	fmt.Println("")
}

func showLetterHistoryEnd(letterHistoryEnd []string)  {
	for i := 0; i <= len(letterHistoryEnd)-1; i++ {
		fmt.Print(letterHistoryEnd[i])
		fmt.Print(" ")
	}
	fmt.Println("")
	fmt.Println("")
}

func printWord(arrSelectWord []string){
	for i := 0; i < len(arrSelectWord); i++ {
		fmt.Print(arrSelectWord[i])
	}
}


func checkWordFind(wordPartiallyReveal []string, letterHistoryEnd []string,arrSelectWord []string) {
	wordFind := true
	for _, letter := range wordPartiallyReveal {
		if letter == "_" {
			wordFind = false
			break
		}
	}
	if wordFind == true {
		ClearTerminal()
		fmt.Printf("\nVous avez deviné le mot !")
		fmt.Println("")
		fmt.Print("Vos propositions ont étaient : ")
		showLetterHistoryEnd(letterHistoryEnd)
	}else if liveJose == 0{
		ClearTerminal()
		fmt.Print("\nVous n'avez plus de vie. Le mot était : ")
		printWord(arrSelectWord)
		fmt.Println("")
		fmt.Println("")
		printLive(0)
		fmt.Print("Vos propositions ont étaient : ")
		showLetterHistoryEnd(letterHistoryEnd)
		fmt.Println("Vous êtes pendu !")
	}else{
		startGame(arrSelectWord,wordPartiallyReveal,liveJose)
	}
}

/*
live := []string{"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n","=========\n","  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n" ,"  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n" , "  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n","  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n" ,"  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n" , "    |  \n    |  \n    |  \n    |  \n    |  \n=========\n" ,""}
*/