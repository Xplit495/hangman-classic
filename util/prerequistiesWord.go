package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func chooseDifficulty() {
	var difficulty int
	for i := 0; i <= 1; i++ {
		util.ClearTerminal()
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

func selectDictionnary(difficulty int) {
	currentDir, _ := os.Getwd()
	dictionnaryPath := currentDir + "\\resources\\dictionnary"

	switch difficulty {
	case 1:
		absolutePath := dictionnaryPath + "words.txt"
		selectRandomWordIntoDictionnary(absolutePath)
	case 2:
		absolutePath := dictionnaryPath + "words2.txt"
		selectRandomWordIntoDictionnary(absolutePath)
	case 3:
		absolutePath := dictionnaryPath + "words3.txt"
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
	util.printWordPartiallyReveal(wordPartiallyReveal)
	fmt.Println("")
	startGame(arrSelectWord, wordPartiallyReveal, 10)
}
