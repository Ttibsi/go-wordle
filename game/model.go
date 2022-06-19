package game

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"

    "github.com/charmbracelet/lipgloss"
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

var style = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("63")). // Change this colour to be grey
    BorderTop(true).
    BorderLeft(true).
    BorderBottom(true).
    BorderRight(true)

type errMsg error

type model struct {
    answer []string
    guesses [][]string // user input guesses in char arrays
	textInput textinput.Model
    err error
}

func InitialModel() model {
    ti := textinput.New()
	ti.Placeholder = "Enter guess"
	ti.Focus()
	ti.CharLimit = 5
	ti.Width = 20

	return model{
        answer: generateAnswer(),
        guesses: [][]string{},
        textInput: ti,
        err: nil,
	}
}

func (m model) Init() tea.Cmd {
    //fmt.Println(lipgloss.NewStyle().Bold(true).Render("Hello, kitty."))
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

	switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
		case tea.KeyEnter:
            var userInput = m.textInput.Value()

            //checkGuess() // Need output from textbox
            //generate new UI - I think this is done automatically when you add a new item
        case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
        switch msg.String() {
        case "ctrl+c", "q":
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
    for _, item := range m.guesses {
        // Render row
        var row string
        for _, letter  := range item {
            row += style.Render(letter + " ")
        }

        s += fmt.Sprintf(row)
    }

    //entry box
    s += m.textInput.View()

    // The footer
    s += "\nq or ctrl+c to quit.\n"

    // Send the UI for rendering
    return s
}
