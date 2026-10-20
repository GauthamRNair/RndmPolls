package main

import (
	"fmt"
	"log"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/ssh"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func ()

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithHostKeyPath(".ssh/wish_ed25519"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					fmt.Fprintf(sess, "TESTTTS %s\n", sess)
					next(sess)
				}
			},
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