package tui

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

func generateAnswer() []string {
	// Is there seriously no easier way to readlines a file?
	var wordlist []string
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
    answer := wordlist[rand.Intn(len(wordlist))]
    return strings.Split(answer, "")
}
