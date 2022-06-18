package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
    guesses [6][5]string // Words already guessed by user
	answer  []string   // seed word generated from txtfile as char array
	textInput textinput.Model
}

func initialModel() model {
	return model{
		guesses: [6][5]string{},
		answer:  generateAnswer(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
        case tea.KeyMsg:
            switch msg.String() {
            case "ctrl+c", "q":
                return m, tea.Quit
            case "enter":
                var newGuess = getNewGuess()
                // Take text entry from elsewhere and trigger a function to check word
                // append as an array
            }
	}

    return m, nil
}

// User interface stored as a string
func (m model) View() string {
    // The header
    s := "Go-Wordle\n\n"

    for _, item := range m.guesses {
        // Render row
        var row string
        for _, letter  := range item {
            row += letter + " "
           // Still need to figure out cards
        }
        s += fmt.Sprintf(row)

        //s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nq or ctrl+c to quit.\n"

    // Send the UI for rendering
    return s
}

