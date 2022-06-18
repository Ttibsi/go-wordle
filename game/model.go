package game

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type errMsg error

type model struct {
    answer []string
    guesses [][]string // user input guesses in char arrays
	textInput textinput.Model
    err error
}

func InitialModel() model {
    ti := textinput.New()
	ti.Placeholder = "Wordl"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
        answer: generateAnswer(),
        guesses: [][]string{},
        textInput: ti,
        err: nil,
	}
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

	switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
		case tea.KeyEnter:
            checkGuess()

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
            row += letter + " "
            // Still need to figure out cards
            // Potentially using lipgloss
        
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
