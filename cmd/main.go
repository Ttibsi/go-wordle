package main

import (
	"fmt"
	"os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/Ttibsi/go-wordle/tree/tui/pkg"
)

func main() {
    p := tea.NewProgram(pkg.initialModel())
        if err := p.Start(); err != nil {
            fmt.Printf("Alas, there's been an error: %v", err)
            os.Exit(1)
        }
}
