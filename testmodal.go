package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var dialogBoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#874BFD")).Padding(1, 0).BorderTop(true).BorderLeft(true).BorderRight(true).BorderBottom(true)


type TestModal struct{
}

func (t TestModal) Init() tea.Cmd{
	return nil
}

func (t TestModal) View() string{
	return dialogBoxStyle.Render("Test Modal!")
}

func (t TestModal) Update(msg tea.Msg) (tea.Model,tea.Cmd) {
	return t,nil
}

func NewTestModal() *TestModal{
	return &TestModal{}
}
