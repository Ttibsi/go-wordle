package game

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

/*
We want an empty board with the header, footer, and input box to start
When the user types a word and presses the enter button, the input
is checked against the word, and each letter appears on the screen with
a border around it in the appropriate colour. If the word is otherwise
inappropriate, the screen will shake or there will be a warning of some
kind informing the user
*/

var style = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderTop(true).
	BorderLeft(true).
	BorderBottom(true).
	BorderRight(true).
	Width(10).
	Padding(2).
	MaxHeight(6).
	MaxWidth(5)

type errMsg error

type model struct {
	answer    []string
	guesses   [6][5]string // user input guesses in char arrays
	scores    [6][5]int    // 1 for red, 2 for yellow, 3 for green, 0 for grey
	turn      int
	textInput textinput.Model
	err       error
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter guess"
	ti.Focus()
	ti.CharLimit = 5
	ti.Width = 20

	return model{
		answer: generateAnswer(),
		guesses: [6][5]string{
			{" ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " "},
		},
		scores: [6][5]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		turn:      0,
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if !validateEntry(m.textInput.Value()) {
				m.textInput.Placeholder = "Invalid guess"
				m.textInput.SetValue("")
			} else {
				m.turn += 1
				var userInput = m.textInput.Value()
				m.scores[m.turn] = m.checkGuess(userInput)
				m.textInput.SetValue("")

				// Check if game over
				if hasWon(m.scores[m.turn]) {
					endGame(true, m.turn, strings.Join(m.answer, ""))
				} else if m.turn == 6 {
					endGame(false, m.turn, strings.Join(m.answer, ""))
				}
			}

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
		// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// User interface stored as a string
func (m model) View() string {
	// The header
	s := "Go-Wordle\n\n"

	// Words and display
	for idx, item := range m.guesses {
		// Render row
		var row string
		for idy, letter := range item {
			if m.scores[idx][idy] == 0 {
				row += style.BorderForeground(lipgloss.Color("#aaaaaa")).Render(letter)
			} else if m.scores[idx][idy] == 1 {
				row += style.BorderForeground(lipgloss.Color("#c93b22")).Render(letter)
			} else if m.scores[idx][idy] == 2 {
				row += style.BorderForeground(lipgloss.Color("#db9c27")).Render(letter)
			} else if m.scores[idx][idy] == 3 {
				row += style.BorderForeground(lipgloss.Color("#66b823")).Render(letter)
			}
		}

		s += fmt.Sprintf(row + "\n")
	}

	//entry box
	s += m.textInput.View()

	// The footer
	s += "\nq or ctrl+c to quit.\n"

	// Send the UI for rendering
	return s
}
