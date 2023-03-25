package sw

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SW struct {
	mouse tea.MouseMsg
}

func New() *SW {
	return &SW{}
}

func (s *SW) Init() tea.Cmd {
	return tea.ClearScreen
}

func (s *SW) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s *SW) View() string {
	return "hello"
}
