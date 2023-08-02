package main

// basic scrollbar demo

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhengkyl/pearls/scrollbar"
)

type model struct {
	scrollbar *scrollbar.Model
}

func newModel() model {
	s := scrollbar.New()
	s.Height = 15

	// How many scroll "states"
	// Generally, number of overflowing items + 1
	s.NumPos = 20
	return model{s}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.scrollbar.Pos > 0 {
				m.scrollbar.Pos--
			}
		case "down", "j":
			if m.scrollbar.Pos < m.scrollbar.NumPos-1 {
				m.scrollbar.Pos++
			}
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.scrollbar.View()
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
