package main

import (
	"fmt"
	"os"

	"imz/model"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".png"}
	fp.CurrentDirectory = "/home/gafoli/Pictures"

	m := model.NewModel("foo")
	m.Filepicker = fp

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error!!: %v", err)
		os.Exit(1)
	}
}
