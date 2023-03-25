package sw

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdb/seaweed-cli/internal/config"
)

type SW struct {
	config config.Config
}

func New(c config.Config) *SW {
	return &SW{
		c,
	}
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
	return s.config.Spots[0].ID
}
