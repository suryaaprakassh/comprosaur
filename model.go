package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	cwd    *Cwd
	cursor int
	status string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) handleKeys(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "q":
		return m, tea.Quit
	case "l":
		m.cwd.selectChild(m.cursor)
		m.cursor = 0
	case "h":
		if err := m.cwd.moveBack(); err != nil {
			m.status = err.Error()
		}
	case "j":
		m.cursor += 1
		m.cursor += len(m.cwd.Children)
		m.cursor %= len(m.cwd.Children)
	case "k":
		m.cursor -= 1
		m.cursor += len(m.cwd.Children)
		m.cursor %= len(m.cwd.Children)
	}
	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("DIR: %s\n\n", m.cwd.path)
	for idx, child := range m.cwd.Children {
		if idx == m.cursor {
			s += fmt.Sprintf("%s %s\n", ">", child.String())
		} else {
			s += fmt.Sprintf("%s %s\n", " ", child.String())
		}
	}
	s += fmt.Sprintf("\n\nSTATUS: %s\n", m.status)
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.status = ""
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeys(msg.String())
	}
	return m, nil
}

func initialModel() model {
	cwd, err := NewCwd()
	if err != nil {
		log.Fatal(err)
	}
	return model{
		cursor: 0,
		cwd:    cwd,
	}
}

