package game

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
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

func endGame(hasWon bool, turn int, answer string) string {
	if hasWon {
		var ret = "Contgratulations, you've won in " + strconv.Itoa(turn) + " turns"
		return ret
	} else {
		return "You lose. Feel free to try again!\nThe word was: " + answer
	}

}

// grid rendering based off of:
// https://github.com/nimblebun/wordle-cli/blob/master/game/grid.go
func (m model) renderTile(ch string, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(color).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(color).
		Render(string(ch))
}

func (m model) renderGridRow(rowIdx int, row [5]string) string {
	var output []string

	for colIdx, col := range row {
		letter_score := m.scores[rowIdx][colIdx]

		if letter_score == 0 {
			output = append(output, m.renderTile(col, lipgloss.Color("#aaa")))
		} else if letter_score == 1 {
			output = append(output, m.renderTile(col, lipgloss.Color("#cc2929")))
		} else if letter_score == 2 {
			output = append(output, m.renderTile(col, lipgloss.Color("#e09926")))
		} else if letter_score == 3 {
			output = append(output, m.renderTile(col, lipgloss.Color("#80bf02")))
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, output...)
}

func (m model) renderGrid() string {
	var output []string

	for idx, row := range m.guesses {
		renderRow := m.renderGridRow(idx, row)
		output = append(output, lipgloss.NewStyle().Padding(0, 1).Render(renderRow))
	}

	// unpack output using the ellipsis
	return lipgloss.JoinVertical(lipgloss.Top, output...)
}
