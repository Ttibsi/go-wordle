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
    // Need to call in the tea model and add local import above too
	entries := 0

	fmt.Println("Welcome to go-wordle")
	fmt.Print("All answers are in complete lowercase\n\n")

	// Get random word from list as array of char
    var validWords []string = getValidWords()
	var answer string = generateAnswer(validWords)
    var answerArray []string = strings.Split(answer, "") 

	for entries < 6 {
		// Ask user to enter a word
		var userInput string
		fmt.Print("Enter a word: ")
		fmt.Scan(&userInput)

		// Validate word is in valid list
		inputWordArray := []string(strings.Split(userInput, ""))
		if len(inputWordArray) != 5 {
			fmt.Println("Word invalid length")
			continue
        } else if !wordOnFile(validWords, userInput) {
			fmt.Println("Word not present")
			continue
		} else {
			entries += 1
		}

		// Check each letter in input against answer
		var score [5]string
		for idx, val := range inputWordArray {
			if val == answerArray[idx] {
				score[idx] = "Y"
			} else if checkElsewhere(val, answerArray) {
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

	fmt.Println("\nThe word was:", answer)
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

func getValidWords() []string {
	// Open wordlist.txt
	file, err := os.Open("wordlist.txt")
	check(err)
	defer file.Close() // This will close the file when the function ends

	sc := bufio.NewScanner(file)
	var words []string

	for sc.Scan() {
		words = append(words, sc.Text())
	}

    return words
}

func generateAnswer(words []string) string {
	// return a random word from the file
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func wordOnFile(validWords []string, userInput string) bool {
    for _, word := range validWords {
        if word == userInput {
            return true
        }
    }

    return false
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
