package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(FileType)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i.String())

	fn := itemStyle.Render

	if i.Marked {
		fn = func(s...string) string {
			return  markedItemStyle.Render("• " + strings.Join(s," "))
		}
	}
	// Heirarchy is important 
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	if index == m.Index() && i.Marked {
		fn = func(s ...string) string {
			return markedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

