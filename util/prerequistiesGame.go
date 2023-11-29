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
	for i := 0; i <= 1; i++ { //Loop to check if the user enter a correct value
		ClearTerminal()
		fmt.Println("")
		fmt.Print("Choissisez votre mode (1-3), 1: Facile, 2: Moyen, 3: Difficile, 4: Multijoueur. Que choissisez-vous : ")
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée standard")
		}
		if difficulty != 1 && difficulty != 2 && difficulty != 3 && difficulty != 4 { //Check if the user enter a correct value
			i-- //If the user enter a wrong value, the loop restart
		} else {
			break
		}
	}
	selectDictionnaryPath(asciiMode, pathAscii, difficulty)
}

func selectDictionnaryPath(asciiMode string, pathAscii string, difficulty int) {
	currentDir, _ := os.Getwd()                                  //Get the current directory
	dictionnaryPath := currentDir + "\\resources\\dictionnary\\" //Set the path to the dictionnary

	switch difficulty {
	case 1:
		absolutePath := dictionnaryPath + "easy.txt" //Set the exact path of the dictionnary in function of the difficulty
		selectRandomWordIntoDictionnary(asciiMode, pathAscii, absolutePath)
	case 2:
		absolutePath := dictionnaryPath + "medium.txt" //Set the exact path of the dictionnary in function of the difficulty
		selectRandomWordIntoDictionnary(asciiMode, pathAscii, absolutePath)
	case 3:
		absolutePath := dictionnaryPath + "hard.txt" //Set the exact path of the dictionnary in function of the difficulty
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
		toFind = strings.ToLower(toFind)   //Convert the word to lowercase to avoid worries
		arrSelectWordRune = []rune(toFind) //Convert the word to rune to check if the user enter only letters

		for j := 0; j < len(arrSelectWordRune); j++ {
			if arrSelectWordRune[j] >= rune(97) && arrSelectWordRune[j] <= rune(122) { //Check if the user enter only letters
			} else {
				ClearTerminal()
				fmt.Println("Merci de saisir" + red + " UNIQUEMENT " + reset + "des caractère de l'alphabet !")
				i-- //If the user enter a wrong value, the loop restart
				break
			}
		}
	}

	arrSelectWord = strings.Split(toFind, "") //Convert the word to an array of string
	ClearTerminal()

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
		numberOfWords++ //Count the number of words in the dictionnary
	}
	err := f.Close()
	if err != nil {
	}
	indexRandomWord = rand.Intn(numberOfWords) + 1 //Select a random number between 1 and the number of words in the dictionnary

	currentLine := 0
	f2, _ := os.Open(absolutePath)
	scanner2 := bufio.NewScanner(f2)
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		currentLine++
		if currentLine == indexRandomWord { //Select the word at the line of the random number
			word = scanner2.Text()
			break
		}
	}
	err2 := f2.Close()
	if err2 != nil {
	}
	arrSelectWord = strings.Split(word, "") //Convert the word to an array of string

	generateWordClue(asciiMode, pathAscii, arrSelectWord)
}

func generateWordClue(asciiMode string, pathAscii string, arrSelectWord []string) {
	var (
		randomClues []int
		n           = (len(arrSelectWord) / 2) - 1 //Ytrack condition
	)

	usedClues := make(map[int]bool) //Create a map to check if the random number is already used
	for i := 1; i <= n; i++ {
		var newClue int
		for { //Loop without condition of exit need to wait break statement to exit (while loop in other languages)
			newClue = rand.Intn(len(arrSelectWord) - 1) //Select a random number between 0 and the length of the word
			if usedClues[newClue] == false {            //Check if the random number is already used
				usedClues[newClue] = true //If the random number is not used, add it to the map
				break
			}
		}
		randomClues = append(randomClues, newClue) //Add the random number to the array
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
			wordPartiallyReveal = append(wordPartiallyReveal, "_") //Replace every letter by an underscore
		}
	} else {
		for i := 0; i <= len(arrSelectWord)-1; i++ {
			if i == randomClues[values] { //Check if the index of the letter is equal to the random number
				wordPartiallyReveal = append(wordPartiallyReveal, arrSelectWord[i]) //If yes, add the letter to the array
				if values+1 >= len(randomClues) {
					values = 0 //If the random number is the last of the array, reset the value to 0 to avoid an error
				} else {
					values += 1
				}
			} else {
				wordPartiallyReveal = append(wordPartiallyReveal, "_") //If no, replace the letter by an underscore
			}
		}
	}

	fmt.Println("")
	fmt.Print("\nLe mot avec le(s) indice(s) est : ")
	printWordPartiallyReveal(asciiMode, pathAscii, wordPartiallyReveal)
	fmt.Println("")
	startGame(asciiMode, pathAscii, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, liveJose)
}
