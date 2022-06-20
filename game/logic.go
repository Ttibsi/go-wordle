package game

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generateAnswer() []string {
	// Is there seriously no easier way to readlines a file?
	var wordlist []string
	file, err := os.Open("game/wordlist.txt")
	if err != nil {
		log.Println(err)
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

func validateEntry(inp string) bool {
	file, err := os.Open("game/wordlist.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == inp {
			return true
		}
	}

	return false
}

func (m model) checkGuess(inp string) [5]int {
	var answer = m.answer
	var guessScore [5]int

	for idx, char := range inp {
		var char = string(char)
		if char == answer[idx] {
			guessScore[idx] = 3
		} else if checkElsewhere(char, answer) {
			guessScore[idx] = 2
		} else {
			guessScore[idx] = 1
		}
	}

	return guessScore
}

func checkElsewhere(char string, ans []string) bool {
	for _, letter := range ans {
		if letter == char {
			return true
		}
	}

	return false
}

func hasWon(score [5]int) bool {
	for _, val := range score {
		if val != 3 { // 3 == green square
			return false
		}
	}

	return true
}

func endGame(hasWon bool, turn int, answer string) {
	if hasWon {
		fmt.Printf("Congratulations, you won in %v turns\n", turn)
	} else {
		fmt.Println("You lose. Feel free to try again!")
		fmt.Println("The word was: ", answer)
	}

}
