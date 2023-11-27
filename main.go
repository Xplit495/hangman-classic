package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
)

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

//LiveJose a la fin du code !!!!

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
	var wordPatriallyReveal []string
	if len(randomClues) == 0 {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			wordPatriallyReveal = append(wordPatriallyReveal, "_")
		}
	} else {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			if i == randomClues[values] { // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
				wordPatriallyReveal = append(wordPatriallyReveal, arrSelectWord[i])
				if values+1 >= len(randomClues) {
					values = 0
				} else {
					values += 1
				}
			}else {
				wordPatriallyReveal = append(wordPatriallyReveal, "_")
			} // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
		}
	}

	fmt.Println("")
	fmt.Print("\nLe mot avec le(s) indice(s) est : ")
	for i := 0; i < len(wordPatriallyReveal); i++ {
		fmt.Print(wordPatriallyReveal[i])
	}
	fmt.Println("")
	fmt.Println("")
}

func startGame(){
	//Commence à partir d'ici (tu peux tester le reste mais normallement tout fonctionne regarde comment j'ai fais et tu continues pour le reste vasy bonne chance je viens a 15h.)
	// ET SURTOUT MODIFIE PAS MON CODE AU DESSUS A MOINS QUE CA SOIT VRAIMENT OBLIGATOIRE !!!
}


/*
Fais juste des copier coller du code en dessous.

liveJose := 0
live := []string{"", "=========\n", "    |  \n    |  \n    |  \n    |  \n    |  \n=========\n", "  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n"}

	tryRemain := 10 // Nombre d'essais maximum
	var letterHistory []string
	var letterHistoryEnd []string
	for tryRemain > 0 {
		if liveJose > 0 {
			fmt.Println(live[liveJose])
		}
		fmt.Printf("Il vous reste %d vie avant d'être pendu !\n", tryRemain)

		var choice string
		var choiceToLower string
		for i := 0; i <= 1; i++ {
			fmt.Print("Entrez votre lettre : ")
			fmt.Scanln(&choice)
			choiceToLower = strings.ToLower(choice)
			choiceToLowerRune := []rune(choiceToLower)
			if len(choiceToLowerRune) > 1 || (choiceToLowerRune[0] >= rune(0) && choiceToLowerRune[0] <= rune(96) || (choiceToLowerRune[0] > rune(122))) {
				ClearTerminal()
				fmt.Println("Merci de saisir UN seul caractère de l'alphabet !.")
				i--
			} else {
				break
			}
		}

		letterHistory = append(letterHistory, choiceToLower)
		letterHistoryEnd = append(letterHistoryEnd, choiceToLower)
		letterFind := false

		for index, letter := range arrSelectWord {
			if letter == choiceToLower {
				wordPatriallyReveal[index] = letter
				letterFind = true
			}
		}

		counter := 0
		alreadyUse := false
		for _, char := range letterHistory {
			if choiceToLower == char {
				counter++
				if counter > 1 {
					counter = 0
					alreadyUse = true
					if len(letterHistory) > 0 {
						letterHistory = letterHistory[:len(letterHistory)-1]
					}
				}
			}
		}

			if letterFind == true {
				ClearTerminal()
				if alreadyUse == true {
					fmt.Println("VOUS AVEZ DEJA ESSAYEZ CETTE LETTRE ! FAITES ATTENTION.")
				}
				fmt.Println("")
				fmt.Println("Bonne lettre !")
				fmt.Println("")
				fmt.Printf("Pour le moment le mot ressemble à ca -> ")
				for i := 0; i < len(wordPatriallyReveal); i++ {
					fmt.Print(wordPatriallyReveal[i])
				}
				fmt.Println("")
				fmt.Print("Les lettres déjà essayé sont : ")
				for i := 0; i <= len(letterHistory)-1; i++ {
					fmt.Print(letterHistory[i])
					fmt.Print(" ")
				}
				fmt.Println("")
				fmt.Println("")

			} else {
				ClearTerminal()
				if alreadyUse == true {
					fmt.Println("VOUS AVEZ DÉJA ESSAYEZ CETTE LETTRE ! FAITES ATTENTION.")
				}
				tryRemain--
				liveJose++
				fmt.Println("")
				fmt.Println("Mauvaise lettre.")
				fmt.Println("")
				fmt.Printf("Pour le moment le mot ressemble à ca -> ")
				for i := 0; i < len(wordPatriallyReveal); i++ {
					fmt.Print(wordPatriallyReveal[i])
				}
				fmt.Println("")
				fmt.Print("L'historique de vos lettres est : ")
				for i := 0; i <= len(letterHistory)-1; i++ {
					fmt.Print(letterHistory[i])
					fmt.Print(" ")
				}
				fmt.Println("")
				fmt.Println("")
			}

			// Vérifier si le mot est complètement deviné
			wordFind := true
			for _, letter := range wordPatriallyReveal {
				if letter == "_" {
					wordFind = false
					break
				}
			}
			if wordFind == true {
				fmt.Printf("\nVous avez deviné le mot !")
				fmt.Println("")
				fmt.Print("Vos propositions ont étaient : ")
				for i := 0; i <= len(letterHistoryEnd)-1; i++ {
					fmt.Print(letterHistoryEnd[i])
					fmt.Print(" ")
				}
				fmt.Println("")
				fmt.Println("")
				return
			}
		}

		fmt.Print("\nVous n'avez plus de vie. Le mot était : ")
		for i := 0; i < len(arrSelectWord); i++ {
			fmt.Print(arrSelectWord[i])

		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println(live[10])
		fmt.Print("Vos propositions ont étaient : ")
		for i := 0; i <= len(letterHistoryEnd)-1; i++ {
			fmt.Print(letterHistoryEnd[i])
			fmt.Print(" ")
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Vous êtes pendu !")
	}
}
*/
