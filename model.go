package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/suryaaprakassh/comprosaur/shared"
	"github.com/suryaaprakassh/comprosaur/explorer"
	"github.com/suryaaprakassh/comprosaur/logger"
)


type model struct {
	ctx shared.CTX
	
	explorer tea.Model
	popup tea.Model

	state shared.ModalState
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) SetModal(modal tea.Model) {
	m.popup = modal
}

func (m model) View() string {
	return m.explorer.View()	
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
		newModel , cmd  := m.explorer.Update(msg)
		m.explorer = newModel
		return m, cmd
}

func (m *model) ChangeState(state shared.ModalState) {
	m.state = state
}

func initialModel() model {
	path, err := os.Getwd()

	if err != nil {
		panic("ERROR: " + err.Error())
	}

	logger := logger.New()
	ctx := shared.New(path, logger)

	cwd, err := explorer.NewCwd(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return model{
		ctx: ctx,

		explorer: explorer.InitialModel(cwd,ctx),
		popup: NewTestModal(),

		state: shared.Explorer,
	}
}
