package game

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/textinput"
)

/*
We want an empty board with the header, footer, and input box to start
When the user types a word and presses the enter button, the input
is checked against the word, and each letter appears on the screen with
a border around it in the appropriate colour. If the word is otherwise
inappropriate, the screen will shake or there will be a warning of some
kind informing the user
*/

type errMsg error

type model struct {
	answer    []string
	guesses   [6][5]string // user input guesses in char arrays
	scores    [6][5]int    // 1 for red, 2 for yellow, 3 for green, 0 for grey
	turn      int
	textInput textinput.Model
	err       error
	output    string
	gameOver  bool
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
		output:    "",
		gameOver:  false,
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
				userInputArr := (*[5]string)(strings.Split(userInput, ""))
				m.guesses[m.turn-1] = *userInputArr
				m.scores[m.turn-1] = m.checkGuess(userInput)
				m.textInput.SetValue("")

				// Check if game over
				if hasWon(m.scores[m.turn-1]) {
					m.gameOver = true
					m.output = endGame(true, m.turn, strings.Join(m.answer, ""))
					return m, tea.Quit
				} else if m.turn == 6 {
					m.gameOver = true
					m.output = endGame(false, m.turn, strings.Join(m.answer, ""))
					return m, tea.Quit
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

	// For testing
	s += strings.Join(m.answer, "")
	s += "\n"

	//Render grid
	s += m.renderGrid()
	s += "\n"

	//entry box
	if !m.gameOver {
		s += m.textInput.View()
	}

	// The footer
	s += "\nq or ctrl+c to quit.\n"
	s += m.output

	// Send the UI for rendering
	return s
}
