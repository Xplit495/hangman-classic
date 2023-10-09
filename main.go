package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	hangman()
}

// live := []string{"=========", "    |  \n    |  \n    |  \n    |  \n    |  \n=========", "  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n"}
func hangman() {
	var difficulty int // CHOISIR NIVEAU DIFFICULTE (POUR LE MOMENT UTILISER SEULEMENT DIIFUCLTE 1 ON CODERA LES AUTRES PLUS TARD)
	for i := 0; i <= 1; i++ {
		fmt.Println("")
		fmt.Print("Choose your level of difficulty (1-3), 1: Easy, 2: Medium, 3: Hard. What you choose : ") // CHOISIR NIVEAU DIFFICULTE (POUR LE MOMENT UTILISER SEULEMENT DIIFUCLTE 1 ON CODERA LES AUTRES PLUS TARD)
		fmt.Scanln(&difficulty)
		if difficulty != 1 && difficulty != 2 && difficulty != 3 {
			i = -1
		} else {
			break
		} // CHOISIR NIVEAU DIFFICULTE (POUR LE MOMENT UTILISER SEULEMENT DIIFUCLTE 1 ON CODERA LES AUTRES PLUS TARD)
	}
	if difficulty == 1 { // SELECTIONNE UN MOT AU HASARD DANS LE DICTIONNAIRE ET LE MET DANS UN TABLEAU
		f, _ := os.Open("C:\\Ytrack\\tls-challenge-go-23-24\\hangman-classic\\words.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		wordlist := []string{}
		for scanner.Scan() {
			wordlist = append(wordlist, scanner.Text())
		}
		indexword := rand.Intn(len(wordlist)) - 1 // SELECTIONNE UN MOT AU HASARD DANS LE DICTIONNAIRE ET LE MET DANS UN TABLEAU
		if indexword <= 0 {
			indexword += 1
		}
		selectword := wordlist[indexword]
		tabselectword := []string{}
		for j := 0; j <= len(selectword)-1; j++ {
			tabselectword = append(tabselectword, string(selectword[j]))
		} // SELECTIONNE UN MOT AU HASARD DANS LE DICTIONNAIRE ET LE MET DANS UN TABLEAU

		n := (len(selectword) / 2) - 1 //LIS LA CONSIGNE POUR COMPRENDRE
		randomindexletter := []int{}
		// Debut CHATGPT
		usedIndices := make(map[int]bool)
		for k := 1; k <= n; k++ {
			var newIndex int
			for {
				newIndex = rand.Intn(len(tabselectword) - 1)

				if !usedIndices[newIndex] {
					usedIndices[newIndex] = true
					break
				}
			}
			randomindexletter = append(randomindexletter, newIndex)
		}
		// Fin ChatGpt

		for l := 0; l < len(randomindexletter)-1; l++ { // TRIE LES INDEX DES LETTRES A AFFICHER DANS L'ORDRE CROISSANT SINON J'AVAIS DES BUGS DANS L'AFFICHAGE
			if randomindexletter[l+1] < randomindexletter[l] {
				randomindexletter[l+1], randomindexletter[l] = randomindexletter[l], randomindexletter[l+1]
			} // TRIE LES INDEX DES LETTRES A AFFICHER DANS L'ORDRE CROISSANT SINON J'AVAIS DES BUGS DANS L'AFFICHAGE

		}

		values := 0 // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
		word_partially_reveal := []string{}
		if len(randomindexletter) == 0 {
			for m := 0; m <= len(tabselectword)-1; m++ {
				word_partially_reveal = append(word_partially_reveal, "_")
			}
		} else {
			for o := 0; o <= len(tabselectword)-1; o++ {
				if o == randomindexletter[values] { // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
					word_partially_reveal = append(word_partially_reveal, tabselectword[o])
					if values+1 >= len(randomindexletter) {
						values = 0
					} else {
						values += 1
					}
				} else {
					word_partially_reveal = append(word_partially_reveal, "_")
				} // SERT A AFFICHER SEULEMENT LES LETTRES ALEATOIRES CHOISIS PRECEDEMENT
			}
		}

		fmt.Println(tabselectword)     // SERT A DEBUG (TU VAS COMPRENDRE)
		fmt.Println(randomindexletter) // SERT A DEBUG (TU VAS COMPRENDRE)
		fmt.Print("\nLe mot avec le(s) indice(s) est : ")
		for p := 0; p < len(word_partially_reveal); p++ {
			fmt.Print(word_partially_reveal[p])
		}
		fmt.Println("")
		fmt.Println("")

		tryremain := 6 // Nombre d'essais max

		for tryremain > 0 {
			fmt.Printf("Il vous reste %d vie avant d'être pendu !\n", tryremain)
			fmt.Print("Entrez une lettre : ")
			var choicerune rune
			for q := 0; q < 1; q++ {
				fmt.Scan(&choicerune)
				if choicerune < rune(97) || choicerune > rune(122) {
					fmt.Println("\nMerci de sélectionnez une seule et unique lettre de l'alphabet et en minuscule !")
					fmt.Printf("Il vous reste %d vie avant d'être pendu !\n", tryremain)
					fmt.Print("Entrez votre lettre (minuscule) : ")
					fmt.Scan(&choicerune)
					if choicerune < 'a' && choicerune > 'z' {
						q--
					} else {
						break
					}
				}
			}

			choice := string(choicerune)

			letterfind := false

			for index, lettre := range tabselectword {
				if lettre == choice {
					word_partially_reveal[index] = lettre
					letterfind = true
				}
			}

			if letterfind == true {
				fmt.Println("Bonne lettre !")
			} else {
				tryremain--
				fmt.Println("Mauvaise lettre.")
			}

			// Vérifier si le mot est complètement deviné
			wordfind := true
			for _, lettre := range word_partially_reveal {
				if lettre == "_" {
					wordfind = false
					break
				}
			}

			if wordfind == true {
				fmt.Printf("\nVous avez deviné le mot !")
				fmt.Println("")
				return
			}
		}

		fmt.Print("\n\nVous n'avez plus de vie. Le mot était : ")
		for r := 0; r < len(tabselectword); r++ {
			fmt.Print(tabselectword[r])
		}
	}
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

   // live := []string{"=========", "    |  \n    |  \n    |  \n    |  \n    |  \n=========", "  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n"}
*/
