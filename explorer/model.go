package explorer 

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/suryaaprakassh/comprosaur/context"
	"github.com/suryaaprakassh/comprosaur/utils"
)

const listHeight = 40

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#feb129"))
	markedItemStyle   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#3399ff"))
	partialItemStyle  = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#ff79c6"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type ModalState int 

type model struct {
	cwd *Cwd
	ctx context.CTX
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	s := strings.Builder{}
	_, err := s.WriteString(fmt.Sprintf("DIR: %s\n\n", m.ctx.GetPath()))
	utils.PanicOnError(err)
	_, err = s.WriteString(m.cwd.Children.View())
	utils.PanicOnError(err)
	_, err = s.WriteString(fmt.Sprintf("\n\nSTATUS: %s\n", m.ctx.GetStatus()))
	utils.PanicOnError(err)

	return s.String()
}

func (m model) handleErrorCall(fn func() error) {
	if err := fn(); err != nil {
		m.ctx.UpdateStatus(err.Error())
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.cwd.Children.SetWidth(msg.Width)
		m.cwd.Children.SetHeight(msg.Height - 6)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "l":
			m.handleErrorCall(m.cwd.moveForward)
		case "h":
			m.handleErrorCall(m.cwd.moveBack)
		case "m":
			m.handleErrorCall(m.cwd.markItem)
		case "c":
			m.handleErrorCall(m.cwd.compressSelected)
		case "e":
			m.handleErrorCall(m.cwd.extractSelected)
		case "d":
			m.handleErrorCall(m.cwd.clearChildren)
		case "a":
			m.handleErrorCall(m.cwd.selectChildren)
		}
	}
	var cmd tea.Cmd
	m.cwd.Children, cmd = m.cwd.Children.Update(msg)
	return m, cmd
}


func InitialModel(cwd *Cwd,ctx context.CTX ) model {
	return model{
		cwd: cwd,
		ctx: ctx,
	}
}
