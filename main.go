package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	entries := 0

	fmt.Println("Welcome to go-wordle")
	fmt.Print("All answers are in complete lowercase\n\n")

	// Get random word from list as array of char
	var answer []string = generateAnswer()
	answerWord := strings.Join(answer, "")

	for entries < 6 {
		// Ask user to enter a word
		var userInput string
		fmt.Print("Enter a word: ")
		fmt.Scan(&userInput)

		// Validate word is in valid list
		charArray := []string(strings.Split(userInput, ""))
		if len(charArray) != 5 {
			fmt.Println("Word invalid")
			continue
		} else {
			entries += 1
		}

		// Check each letter in input against answer
		var score [5]string
		for idx, val := range charArray {
			if val == answer[idx] {
				score[idx] = "Y"
			} else if checkElsewhere(val, answer) {
				score[idx] = "O"
			} else {
				score[idx] = "N"
			}
		}

		fmt.Println(score)

		if checkScore(score) {
			break
		}
	}

	fmt.Println("\nThe word was:", answerWord)
	if entries < 6 {
		fmt.Printf("You won in %v guesses\n", entries)
	} else {
		fmt.Println("You lose this round.")
	}
}

func check(e error) {
	if e != nil {
		log.Fatal()
	}
}

func generateAnswer() []string {
	// Open wordlist.txt
	file, err := os.Open("wordlist.txt")
	check(err)
	defer file.Close() // This will close the file when the function ends

	sc := bufio.NewScanner(file)
	var words []string

	for sc.Scan() {
		words = append(words, sc.Text())
	}

	// return a random word from the file
	rand.Seed(time.Now().UnixNano())
	var selectedWord = words[rand.Intn(len(words))]

	return []string(strings.Split(selectedWord, ""))
}

func checkElsewhere(val string, charArray []string) bool {
	for _, letter := range charArray {
		if letter == val {
			return true
		}
	}

	return false
}

func checkScore(score [5]string) bool {
	correctPoints := 0

	for _, val := range score {
		if val == "Y" {
			correctPoints += 1
		}
	}

	if correctPoints == 5 {
		return true
	} else {
		return false
	}
}
