package game

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    answer []string
    guesses [][]string // user input guesses in char arrays
}

func initialModel() model {
	return model{
        answer: generateAnswer(),
        guesses: [][]string{},
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
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
