package main

import (
	"fmt"
	"os"

	"github.com/Ttibsi/go-wordle/game"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(game.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
