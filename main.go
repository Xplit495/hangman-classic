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

func hangman() {
	//live := []string{"=========", "    |  \n    |  \n    |  \n    |  \n    |  \n=========", "  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========\n", "  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n"}
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

		essaisRestants := 6 // Nombre d'essais max

		motPartiel := make([]string, len(tabselectword))
		for i := range motPartiel {
			motPartiel[i] = "_"
		}

		// Révéler deux lettres au hasard au début du jeu
		for i := 0; i < 2; i++ {
			randomIndex := rand.Intn(len(tabselectword))
			motPartiel[randomIndex] = tabselectword[randomIndex]
		}

		for essaisRestants > 0 {
			fmt.Println("\n\nMot partiellement révélé :", motPartiel)
			fmt.Printf("Il vous reste %d essais.\n", essaisRestants)
			fmt.Print("Entrez une lettre : ")
			var choix string
			fmt.Scan(&choix)
			ptrChoix := &choix
			lettreTrouvee := false

			for i, lettre := range tabselectword {
				if lettre == *ptrChoix {
					motPartiel[i] = lettre
					lettreTrouvee = true
				}
			}
			if essaisRestants > 0 {
				for i := 0; i <= 10; i++ {
				}
			}

			if lettreTrouvee {
				fmt.Println("Bonne lettre !")
			} else {
				essaisRestants--
				fmt.Println("Mauvaise lettre.")
			}

			// Vérifier si le mot est complètement deviné
			wordDevine := true
			for _, lettre := range motPartiel {
				if lettre == "_" {
					wordDevine = false
					break
				}
			}

			if wordDevine {
				fmt.Println("\nVous avez deviné le mot !")
				return
			}
		}

		fmt.Println("\n\nVous avez utilisé tous vos essais. Le mot était :", tabselectword)
	}
}
