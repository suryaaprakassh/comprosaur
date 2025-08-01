package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 40

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#feb129"))
	markedItemStyle  = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#3399ff"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)


type model struct {
	cwd    *Cwd
	status string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	s := fmt.Sprintf("DIR: %s\n\n", m.cwd.path)
	s += m.cwd.Children.View()
	s += fmt.Sprintf("\n\nSTATUS: %s\n", m.status)
	return s
}

func (m model) handleErrorCall(fn func() error ) {
		if err := fn();err != nil {
			m.status = err.Error()
		}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.status = ""
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.cwd.Children.SetWidth(msg.Width)
		m.cwd.Children.SetHeight(msg.Height - 6)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m,tea.Quit
		case "l":
			m.handleErrorCall(m.cwd.moveForward)
		case "h":
			m.handleErrorCall(m.cwd.moveBack)
		case "m":
			m.handleErrorCall(m.cwd.markItem)
		}
	}
	var cmd tea.Cmd
	m.cwd.Children, cmd = m.cwd.Children.Update(msg)
	return m , cmd
}

func initialModel() model {
	cwd, err := NewCwd()
	if err != nil {
		log.Fatal(err)
	}

	return model{
		cwd:    cwd,
	}
}
