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

func chooseDifficulty() {
	var difficulty int
	for i := 0; i <= 1; i++ {
		ClearTerminal()
		fmt.Println("")
		fmt.Print("Choose your level of difficulty (1-3), 1: Easy, 2: Medium, 3: Hard. What you choose : ")
		fmt.Scanln(&difficulty)
		if difficulty == 1 || difficulty == 2 || difficulty == 3 {
			break
		}
	}

	var filename string

	switch difficulty {
	case 1:
		filename = "C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words.txt"
	case 2:
		filename = "C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words2.txt"
	case 3:
		filename = "C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words3.txt"
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	wordlist := []string{}
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	randomWord := rand.Intn(len(wordlist))
	if randomWord < 0 {
		randomWord = 0
	}
	selectWord := wordlist[randomWord]

	hangman(selectWord)
}

func hangman(selectWord string) {
	var arrSelectWord []string
	liveJose := 0
	live := []string{"", "=========\n", "    |  \n    |  \n    |  \n    |  \n    |  \n=========\n", "  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n"}
	n := (len(selectWord) / 2) - 1 //LIS LA CONSIGNE POUR COMPRENDRE
	var randomIndexLetter []int
	usedIndices := make(map[int]bool)

	for i := 1; i <= n; i++ {
		var newIndex int
		for {
			newIndex = rand.Intn(len(arrSelectWord) - 1)

			if !usedIndices[newIndex] {
				usedIndices[newIndex] = true
				break
			}
		}
		randomIndexLetter = append(randomIndexLetter, newIndex)
	}

	sort.Ints(randomIndexLetter)

	values := 0 // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
	var wordPatriallyReveal []string
	if len(randomIndexLetter) == 0 {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			wordPatriallyReveal = append(wordPatriallyReveal, "_")
		}
	} else {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			if i == randomIndexLetter[values] { // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
				wordPatriallyReveal = append(wordPatriallyReveal, arrSelectWord[i])
				if values+1 >= len(randomIndexLetter) {
					values = 0
				} else {
					values += 1
				}
			} else {
				wordPatriallyReveal = append(wordPatriallyReveal, "_")
			} // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
		}
	}

	fmt.Print("\nLe mot avec le(s) indice(s) est : ")
	for i := 0; i < len(wordPatriallyReveal); i++ {
		fmt.Print(wordPatriallyReveal[i])
	}
	fmt.Println("")
	fmt.Println("")

	tryRemain := 10 // Nombre d'essais max

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

// NORMALEMENT LE CODE FONCTIONNE MAIS SI TU VOIS UN SOUCIS PREND UNE PHOTO ET ESSAYE DE LE CORRIGER APRES JE M'EN OCCUPERAIS SI T'AS PAS TROUVE !!
// BONNE CHANCE MAIS TU VAS VOIR C'EST PAS DUR A COMPRENDRE QUAND C'EST DEJA ECRIS ET LA LA MOITIE DU TRAVAIL EST DEJA FINIS NORMALMENT LE RESTE A L'AIR FACILE !!

/* A MODIFIER QUE POUR LE SYSTEME DE NIVEAU SINON TU T'OCCUPES QUE DE CE QUI EST AU DESSUS (ON FERA UN COPIER COLLER DU CODE POUR LES AUTRES NIVEAU COMME CA SERA LA MEME CHOSE DE TOUTE FACON !)
   	if difficulty == 2 {
   		f, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words2.txt")

   		defer f.Close()

   		scanner := bufio.NewScanner(f)
   		scanner.Split(bufio.ScanWords)
   		wordlist := []string{}
   		for scanner.Scan() {
   			wordlist = append(wordlist, scanner.Text())
   		}
   		indexword := rand.Intn(len(wordlist)) - 1
   		if indexword == 0 {
   			indexword += 1
   		}
   		selectword := wordlist[indexword]

   		for i := 0; i <= len(selectword)-1; i++ {
   			fmt.Print("_")
   		}
   	}

   	if difficulty == 3 {
   		f, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words3.txt")

   		defer f.Close()

   		scanner := bufio.NewScanner(f)
   		scanner.Split(bufio.ScanWords)
   		wordlist := []string{}
   		for scanner.Scan() {
   			wordlist = append(wordlist, scanner.Text())
   		}
   		indexword := rand.Intn(len(wordlist)) - 1
   		if indexword == 0 {
   			indexword += 1
   		}
   		selectword := wordlist[indexword]

   		for i := 0; i <= len(selectword)-1; i++ {
   			fmt.Print("_")
   		}
   	}

   }
*/
