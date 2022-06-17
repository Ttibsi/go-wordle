package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    guesses []string // Words already guessed by user
    answer string // seed word generated from txtfile
}

func initialModel() model {
	return model{
        guesses: []string{},
        answer: generateAnswer(),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func generateAnswer() string {
    // Is there seriously no easier way to readlines a file?
    var wordlist []string
    file, err := os.Open("wordlist.txt")
    if err != nil { log.Fatal() }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wordlist = append(wordlist, scanner.Text())
    }

	rand.Seed(time.Now().UnixNano())
	return wordlist[rand.Intn(len(wordlist))]
}

//Update function
