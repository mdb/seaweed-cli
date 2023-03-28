package sw

import (
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdb/seaweed"
	"github.com/mdb/seaweed-cli/internal/config"
)

type SW struct {
	config     config.Config
	client     *seaweed.Client
	forecasts  []seaweed.Forecast
	tabs       []string
	tabContent []string
	activeTab  int
}

func New(debug bool, c config.Config) *SW {
	sw := &SW{
		config: c,
		client: seaweed.NewClient(os.Getenv("MAGIC_SEAWEED_API_KEY")),
	}

	return sw
}

func contains(s []string, str string) (int, bool) {
	for i, v := range s {
		if v == str {
			return i, true
		}
	}

	return 0, false
}

func (s *SW) Init() tea.Cmd {
	f, _ := s.client.Forecast(s.config.Spots[0].ID)
	s.forecasts = f
	s.tabs = []string{}
	s.tabContent = []string{}

	for _, forecast := range s.forecasts {
		fTS := time.Unix(forecast.LocalTimestamp, 0)
		now := time.Now()

		// TODO: perhaps we should only continue if the time is before today (and not before now)
		if now.After(fTS) {
			continue
		}

		day := time.Unix(forecast.LocalTimestamp, 0).Weekday().String()

		i, contains := contains(s.tabs, day)
		if !contains {
			s.tabs = append(s.tabs, day)
			// TODO: append real forecast rendering
			s.tabContent = append(s.tabContent, "hello")
		} else {
			// TODO: append real forecast rendering
			s.tabContent[i] += " hello"
		}
	}

	return tea.ClearScreen
}

func (s *SW) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return s, tea.Quit
		case "right", "l", "n", "tab":
			s.activeTab = min(s.activeTab+1, len(s.tabs)-1)
			return s, nil
		case "left", "h", "p", "shift+tab":
			s.activeTab = max(s.activeTab-1, 0)
			return s, nil
		}
	}

	return s, nil
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right

	return border
}

func (s *SW) View() string {
	/*
		sb := strings.Builder{}
		d := time.Unix(s.forecasts[0].LocalTimestamp, 0)
		sb.WriteString(s.config.Spots[0].Name + "\n")
		sb.WriteString(d.Format("Monday, January 2, 2006"))
	*/

	doc := strings.Builder{}

	var docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	var activeTabBorder = tabBorderWithBottom("┘", " ", "└")
	var inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	var highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	var inactiveTabStyle = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	var activeTabStyle = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	var windowStyle = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
	var renderedTabs []string

	for i, t := range s.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(s.tabs)-1, i == s.activeTab
		style = inactiveTabStyle.Copy()

		if isActive {
			style = activeTabStyle.Copy()
		}

		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(s.tabContent[s.activeTab]))

	return docStyle.Render(doc.String())

	//return sb.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
