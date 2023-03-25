package sw

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdb/seaweed"
	"github.com/mdb/seaweed-cli/internal/config"
)

type SW struct {
	config    config.Config
	client    *seaweed.Client
	forecasts []seaweed.Forecast
}

func New(debug bool, c config.Config) *SW {
	sw := &SW{
		config: c,
		client: seaweed.NewClient(os.Getenv("MAGIC_SEAWEED_API_KEY")),
	}

	return sw
}

func (s *SW) Init() tea.Cmd {
	f, _ := s.client.Forecast(s.config.Spots[0].ID)
	s.forecasts = f

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
	d := time.Unix(s.forecasts[0].LocalTimestamp, 0)

	return fmt.Sprintf("%s: %s", s.config.Spots[0].Name, d.Format("Monday, January 2, 2006"))
}
