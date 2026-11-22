package main

import (
	"fmt"
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"

	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"

	"database/sql"
	_ "embed"
	"strings"

	_ "modernc.org/sqlite"
)

//go:embed banner.txt
var banner string

var db *sql.DB

var (
	colPink    = lipgloss.AdaptiveColor{Light: "#ea76cb", Dark: "#f5c2e7"}
	colMauve   = lipgloss.AdaptiveColor{Light: "#8839ef", Dark: "#cba6f7"}
	colText    = lipgloss.AdaptiveColor{Light: "#4c4f69", Dark: "#cdd6f4"}
	colSubtext = lipgloss.AdaptiveColor{Light: "#6c6f85", Dark: "#a6adc8"}
	colOverlay = lipgloss.AdaptiveColor{Light: "#9ca0b0", Dark: "#6c7086"}
	colPeach   = lipgloss.AdaptiveColor{Light: "#fe640b", Dark: "#fab387"}
	colGreen   = lipgloss.AdaptiveColor{Light: "#40a02b", Dark: "#a6e3a1"}
)

type model struct {
	user          string
	poll          Poll
	cursor        int
	picked        string
	choice        int
	width, height int
	voted         bool

	banner     lipgloss.Style
	intro      lipgloss.Style
	subtitle   lipgloss.Style
	selected   lipgloss.Style
	unselected lipgloss.Style
	normal     lipgloss.Style
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
			if m.cursor < len(m.poll.Choices)-1 {
				m.cursor++
			}
		case "enter":
			if m.voted {
				break
			}

			m.picked = m.poll.Choices[m.cursor]
			m.choice = m.cursor
			log.Println("insert", getDayNum(), m.user)
			_, err := db.Exec(`INSERT INTO votes (day, voter, choice) VALUES (?, ?, ?)`, getDayNum(), m.user, m.choice)

			if err != nil {
				log.Println("vote insert error: ", err)
			}

			m.voted = true
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m model) View() string {
	if !m.voted {
		return ViewMainPage(m)
	} else {
		return ViewResults(m)
	}
}

func ViewMainPage(m model) string {
	s := m.banner.Render(banner) + "\n"
	s += m.intro.Render(fmt.Sprintf("> Welcome, %s", strings.ToUpper(m.user))) + "\n"
	s += m.subtitle.Render("- A new question every day. Vote, then see the results!") + "\n"
	s += "──────────────────────────────────────────────────────────" + "\n\n"

	s += fmt.Sprintf("Today's Poll: \"%s\"\n", m.poll.Question)

	for i, c := range m.poll.Choices {
		if i == m.cursor {
			s += m.selected.Render("> "+c) + "\n"
		} else {
			s += "  " + m.unselected.Render(c) + "\n"
		}
	}

	s += "\n"

	if m.picked != "" {
		s += m.normal.Render("You picked: "+m.picked) + "\n\n"
	}

	s += m.normal.Render("Help || q quit · h poll history · ↑/↓ move · enter select")

	// left align the whole block
	block := lipgloss.NewStyle().Align(lipgloss.Left).Render(s)

	// lipgloss auto handles the new sizes
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, block)
}

func ViewResults(m model) string {
	s := m.banner.Render(banner) + "\n"
	s += m.intro.Render(fmt.Sprintf("> Well here are today's results, %s", strings.ToUpper(m.user))) + "\n"
	s += m.subtitle.Render("- You can't go back to change your vote!") + "\n"
	s += "──────────────────────────────────────────────────────────" + "\n\n"

	s += fmt.Sprintf("Today's Poll: \"%s\"\n", m.poll.Question)

	var numVotes int
	var totalVotes = getTotalVotes()
	for i, c := range m.poll.Choices {
		err := db.QueryRow(
			"SELECT COUNT(*) FROM votes WHERE day = ? AND choice = ? GROUP BY choice",
			getDayNum(), i).Scan(&numVotes)

		if err == sql.ErrNoRows {
			numVotes = 0
		} else if err != nil {
			log.Println("vote results lookup failed: ", err)
		}

		percentVotes := (numVotes * 100) / totalVotes

		if i == m.cursor {
			s += m.selected.Render(" "+strconv.Itoa(percentVotes)+"% | "+c+" | "+strconv.Itoa(numVotes)) + "\n"
		} else {
			s += m.unselected.Render(" "+strconv.Itoa(percentVotes)+"% | "+c) + "\n"
		}

	}

	s += "\n"

	if m.picked != "" {
		s += m.normal.Render("Your pick was: "+m.picked) + "\n\n"
	}

	s += m.normal.Render("Help || q quit · ↑/↓ move")

	// left align the whole block
	block := lipgloss.NewStyle().Align(lipgloss.Left).Render(s)

	// lipgloss auto handles the new sizes
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, block)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	r := bubbletea.MakeRenderer(s)

	m := model{
		user: s.User(),
		poll: todayPoll(),
		banner: r.NewStyle().Bold(true).
			Foreground(colPink).
			Padding(1, 13),
		intro:      r.NewStyle().Bold(true).Foreground(colText),
		subtitle:   r.NewStyle().Bold(true).Foreground(colSubtext),
		selected:   r.NewStyle().Foreground(colMauve),
		unselected: r.NewStyle().Foreground(colOverlay),
		normal:     r.NewStyle().Faint(true).Foreground(colOverlay),
	}

	// previous vote?
	if choice, ok := hasExistingVote(m.user); ok {
		m.choice = choice
		m.picked = m.poll.Choices[choice]
		m.voted = true
	}

	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

func getTotalVotes() int {
	var totalVotes int

	err := db.QueryRow(
		"SELECT COUNT(*) FROM votes WHERE day = ?",
		getDayNum()).Scan(&totalVotes)

	if err != nil {
		log.Println("total vote lookup failed: ", err)
		return 0
	}

	return totalVotes
}

func hasExistingVote(user string) (int, bool) {
	var choice int

	log.Println("lookup", getDayNum(), user)

	err := db.QueryRow(
		`SELECT choice FROM votes WHERE day = ? AND voter = ?`,
		getDayNum(), user).Scan(&choice)

	// No prior vote found
	if err == sql.ErrNoRows {
		return 0, false
	}

	if err != nil {
		log.Println("vote lookup failed: ", err)
		return 0, false
	}

	return choice, true
}

func main() {
	var err error
	db, err = sql.Open("sqlite", "polls.db")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS votes (
			day INTEGER NOT NULL,
			voter TEXT NOT NULL,
			choice INTEGER NOT NULL,
			UNIQUE(day, voter)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

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
