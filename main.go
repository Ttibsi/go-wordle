package main

import (
    "fmt"
    "os"

    "github.com/Ttibsi/go-wordle/game"
)

func main() {
    p := tea.NewProgram(initialModel())
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
