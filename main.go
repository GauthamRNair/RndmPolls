package main

import (
	"fmt"
	"log"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/ssh"

	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
)

type model struct{
	user string
	choices []string
	cursor int
	picked string
	width, height int

	title lipgloss.Style
	selected lipgloss.Style
	normal lipgloss.Style
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "q", "ctrl+c":
					return m, tea.Quit
				case "up":
					if m.cursor > 0 {
						m.cursor--
					}
				case "down":
					if m.cursor < len(m.choices)-1 {
						m.cursor++
					}
				case "enter":
					m.picked = m.choices[m.cursor]
			}
		case tea.WindowSizeMsg:
			m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m model) View() string {
	s := m.title.Render(fmt.Sprintf("Hello %s!", m.user)) + "\n\n"

	for i, c := range m.choices {
		if i == m.cursor {
			s += "> " + m.selected.Render(c) + "\n"
		}else {
			s += m.normal.Render("  " + c) + "\n"
		}
	}

	if m.picked != "" {
		s += m.normal.Render("\nYou picked: " + m.picked) + "\n"
	}

	s += m.normal.Render("\n↑/↓ move · q quit")

	// lipgloss auto handles the new sizes
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	r := bubbletea.MakeRenderer(s)

	m := model{
		user: s.User(),
		choices: []string{"Coffee", "Tea", "Water"},
		title: r.NewStyle().Bold(true).Foreground(lipgloss.Color("205")),
		selected: r.NewStyle().Foreground(lipgloss.Color("212")),
		normal: r.NewStyle().Faint(true),
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithHostKeyPath(".ssh/wish_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			// These ones are just for logging/safety
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on :2222")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}