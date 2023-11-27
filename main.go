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

var yellow = "\033[33m"
var red = "\033[31m"
var green  = "\033[32m"
var reset  = "\033[0m"
var liveJose = 10
var letterHistory []string
var letterHistoryEnd []string
var currentDir, _ = os.Getwd()
var choiceToLowerRune []rune
var choiceToLowerStrings  []string

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
var difficulty int
	for i := 0; i <= 1; i++ {
		ClearTerminal()
		fmt.Println("")
		fmt.Print("Choissisez votre niveau de difficulté (1-3), 1: Facile, 2: Moyen, 3: Difficile. Que choissisez-vous : ")
		fmt.Scanln(&difficulty)
		if difficulty != 1 && difficulty != 2 && difficulty != 3 {
			i--
		} else {
			break
		}
	}
	selectDictionnary(difficulty)
}

func selectDictionnaryPath(absolutePath string){
	f, _ := os.Open(absolutePath)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	var wordList []string
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	randomWord(wordList)
}

func selectDictionnary(difficulty int){
	switch difficulty {
	case 1:
		relativePath := "words.txt"
		absolutePath := currentDir + "\\" + relativePath
		selectDictionnaryPath(absolutePath)
	case 2:
		relativePath := "words2.txt"
		absolutePath := currentDir + "\\" + relativePath
		selectDictionnaryPath(absolutePath)
	case 3:
		relativePath := "words3.txt"
		absolutePath := currentDir + "\\" + relativePath
		selectDictionnaryPath(absolutePath)
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

	fmt.Print("\nLe mot avec le(s) indice(s) est : ")
	printWordPartiallyReveal(wordPartiallyReveal)
	fmt.Println("")
	startGame(arrSelectWord,wordPartiallyReveal,10)
}

func startGame(arrSelectWord []string, wordPartiallyReveal [] string, liveJose int) {
	printLive(liveJose)
	fmt.Println("")
	fmt.Printf("Il vous reste "+yellow+"%d vie "+reset+"avant d'être pendu !\n", liveJose)
	fmt.Print("Les lettres déjà essayé sont : ")
	printLetterHistoryInGame()
	var choice string
	var choiceToLower string
	for i := 0; i <= 1; i++ {
		fmt.Print("Entrez votre lettre : ")
		fmt.Scan(&choice)
		choiceToLower = strings.ToLower(choice)
		choiceToLowerRune = []rune(choiceToLower)
		for j := 0; j < len(choiceToLowerRune); j++ {
			choiceToLowerStrings = append(choiceToLowerStrings, string(choiceToLowerRune[j]))
		}
		exit := true
		for k := 0; k < len(choiceToLowerRune); k++ {
			if choiceToLowerRune[k] >= rune(97) && choiceToLowerRune[k] <= rune(122) {
				if k+1 == len(choiceToLowerRune) {
					break
				}
			}else{
				ClearTerminal()
				fmt.Println("Merci de saisir" + red + " UNIQUEMENT " + reset + "des caractère de l'alphabet !")
				exit = false
				i--
				}
			}
			if exit == true {
				break
			}
		}
		letterHistory = append(letterHistory, choiceToLower)
		letterHistoryEnd = append(letterHistoryEnd, choiceToLower)
		refreshWord(arrSelectWord,choiceToLower,wordPartiallyReveal,choiceToLowerStrings)
}

func refreshWord(arrSelectWord []string, choiceToLower string, wordPartiallyReveal []string,choiceToLowerStrings []string) {
	//letterFind := false
	for index, letter := range arrSelectWord {
		for i := 0; i < len(choiceToLowerStrings); i++ {
			if letter == choiceToLowerStrings[i] {
				wordPartiallyReveal[index] = letter
				//letterFind = true
			}
		}

	}
	printWordPartiallyReveal(wordPartiallyReveal)
	//histroy(choiceToLower, letterFind ,wordPartiallyReveal,arrSelectWord)
}

func histroy (choiceToLower string, letterFind bool, wordPartiallyReveal []string, arrSelectWord []string){
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
	checkLetterUses(letterFind, letterAlreadyUse, wordPartiallyReveal,arrSelectWord)
}

func checkLetterUses(letterFind bool, alreadyUse bool, wordPartiallyReveal []string, arrSelectWord []string) {
	if alreadyUse == true {
			ClearTerminal()
			fmt.Println(red+"Vous avez déjà essayé cette lettre !"+reset) //Pas fonctionnel mais à laisser pour les advandec features !!!
			startGame(arrSelectWord,wordPartiallyReveal,liveJose)
	}
	fmt.Println("")
	findLetter(wordPartiallyReveal,arrSelectWord,letterFind)
}

func findLetter(wordPartiallyReveal []string,arrSelectWord []string, letterFind bool){
	if letterFind == true{
		ClearTerminal()
		fmt.Println(green+"Bonne lettre !"+reset)
		fmt.Println("")
		fmt.Printf("Pour le moment le mot ressemble à ca -> ")
		printWordPartiallyReveal(wordPartiallyReveal)
		checkWordFind(wordPartiallyReveal,arrSelectWord)
	}else{
		liveJose--
		ClearTerminal()
		fmt.Println(red+"Mauvaise lettre !"+reset)
		fmt.Println("")
		fmt.Printf("Pour le moment le mot ressemble à ca -> ")
		printWordPartiallyReveal(wordPartiallyReveal)
		checkWordFind(wordPartiallyReveal,arrSelectWord)
	}
}

func checkWordFind(wordPartiallyReveal []string,arrSelectWord []string) {
	wordFind := true
	for _, letter := range wordPartiallyReveal {
		if letter == "_" {
			wordFind = false
			break
		}
	}
	if wordFind == true {
		ClearTerminal()
		fmt.Println("\nVous avez deviné le mot !")
		fmt.Print("Vos propositions ont étaient : ")
		printLetterHistoryEnd()
		fmt.Print("Le mot était : ")
		printWord(arrSelectWord)
	}else if liveJose == 0{
		ClearTerminal()
		fmt.Print("\nVous n'avez plus de vie. Le mot était : ")
		printWord(arrSelectWord)
		fmt.Println("")
		fmt.Println("")
		printJose(71,78)
		fmt.Print("Vos propositions ont étaient : ")
		printLetterHistoryEnd()
		fmt.Println("Vous êtes pendu !")
	}else{
		startGame(arrSelectWord,wordPartiallyReveal,liveJose)
	}
}





//Début des fonctions d'affichage
func printWordPartiallyReveal(wordPartiallyReveal []string){
	for i := 0; i < len(wordPartiallyReveal); i++ {
		fmt.Print(wordPartiallyReveal[i])
	}
	fmt.Println("")
}

func printLetterHistoryInGame(){
	for i := 0; i <= len(letterHistory)-1; i++ {
		fmt.Print(letterHistory[i])
		fmt.Print(" ")
	}
	fmt.Println("")
	fmt.Println("")
}
																		//A séparer dans des dossiers
func printLetterHistoryEnd()  {
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
//Fin des fonctions d'affichage








//Debut fonction position Jose
func printJose(startLine int ,endLine int){
	relativePath := "hangman.txt"
	absolutePath := currentDir + "\\" + relativePath
	file, _ := os.Open(absolutePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		currentLine++
		if currentLine >= startLine && currentLine <= endLine {
			fmt.Println(scanner.Text())
		}
		if currentLine > endLine {
			file.Close()
			break
		}
	}
}
																	//A séparer dans des dossiers
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