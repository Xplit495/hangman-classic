package util

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func addStats(asciiMode string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, liveJose int, wordFind bool) {
	var (
		result        string
		currentDir, _ = os.Getwd()
	)

	if wordFind == true {
		result = "Oui"
	} else {
		result = "Non"
	}

	newLine := "\n##############################################################" + //Text to add in file
		"\nPartie du : " + time.Now().Format("02/01/2006 15:04:05") +
		"\nAvez-vous deviné le mot : " + result +
		"\nLe mot était : " + strings.Join(arrSelectWord, "") +
		"\nVotre mot ressemblait à ça : " + strings.Join(wordPartiallyReveal, "") + //Text to add in file
		"\nVous avez essayé les lettres : " + strings.Join(letterHistory, "") +
		"\nVous avez essayé les mots : " + strings.Join(wordHistory, "") +
		"\nIl vous restait : " + strconv.Itoa(liveJose) + " vies" +
		"\nVous avez joué avec le mode ASCII : " + asciiMode +
		"\n##############################################################\n" //Text to add in file

	file, err := os.OpenFile(currentDir+"\\resources\\stats.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644) //Open file for add text (in append mode for not earse the content)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(newLine)
	if err != nil {
		return
	}
}
