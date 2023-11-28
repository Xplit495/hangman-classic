package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func selectDifficulty(asciiMode string, pathAscii string) {
	var difficulty int
	for i := 0; i <= 1; i++ {
		ClearTerminal()
		fmt.Println("")
		fmt.Print("Choissisez votre mode (1-3), 1: Facile, 2: Moyen, 3: Difficile, 4: Multijoueur. Que choissisez-vous : ")
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée standard")
		}
		if difficulty != 1 && difficulty != 2 && difficulty != 3 && difficulty != 4 {
			i--
		} else {
			break
		}
	}
	selectDictionnaryPath(asciiMode, pathAscii, difficulty)
}

func selectDictionnaryPath(asciiMode string, pathAscii string, difficulty int) {
	currentDir, _ := os.Getwd()
	dictionnaryPath := currentDir + "\\resources\\dictionnary\\"

	switch difficulty {
	case 1:
		absolutePath := dictionnaryPath + "words.txt"
		selectRandomWordIntoDictionnary(asciiMode, pathAscii, absolutePath)
	case 2:
		absolutePath := dictionnaryPath + "words2.txt"
		selectRandomWordIntoDictionnary(asciiMode, pathAscii, absolutePath)
	case 3:
		absolutePath := dictionnaryPath + "words3.txt"
		selectRandomWordIntoDictionnary(asciiMode, pathAscii, absolutePath)
	case 4:
		playMultiPlayer(asciiMode, pathAscii)
	}
}

func playMultiPlayer(asciiMode string, pathAscii string) {
	var (
		toFind            string
		arrSelectWord     []string
		arrSelectWordRune []rune
	)
	for i := 0; i < 1; i++ {
		fmt.Print("\nBienvenue dans le mode multijoueur ! Choissisez un mot pour votre adversaire : ")
		_, err := fmt.Scanln(&toFind)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée standard")
		}
		toFind = strings.ToLower(toFind)
		arrSelectWordRune = []rune(toFind)

		for j := 0; j < len(arrSelectWordRune); j++ {
			if arrSelectWordRune[j] >= rune(97) && arrSelectWordRune[j] <= rune(122) {
			} else {
				ClearTerminal()
				fmt.Println("Merci de saisir" + red + " UNIQUEMENT " + reset + "des caractère de l'alphabet !")
				i--
				break
			}
		}
	}

	arrSelectWord = strings.Split(toFind, "")

	generateWordClue(asciiMode, pathAscii, arrSelectWord)
}

func selectRandomWordIntoDictionnary(asciiMode string, pathAscii string, absolutePath string) {
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
	}
	indexRandomWord = rand.Intn(numberOfWords) + 1

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
	}
	arrSelectWord = strings.Split(word, "")

	generateWordClue(asciiMode, pathAscii, arrSelectWord)
}

func generateWordClue(asciiMode string, pathAscii string, arrSelectWord []string) {
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

	associateClueToWord(asciiMode, pathAscii, randomClues, arrSelectWord)
}

func associateClueToWord(asciiMode string, pathAscii string, randomClues []int, arrSelectWord []string) {
	var (
		values              = 0
		wordPartiallyReveal []string
		letterHistory       []string
		wordHistory         []string
		liveJose            = 10
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
	printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
	fmt.Println("")
	startGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
}
