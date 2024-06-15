package internal

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type resourceDelegate struct{}

func (d resourceDelegate) Height() int  { return 1 }
func (d resourceDelegate) Spacing() int { return 0 }
func (d resourceDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}
func (d resourceDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	resource, ok := item.(*Mark)
	if !ok {
		return
	}
	fmt.Fprintf(w, renderMark(resource, index == m.Index()))
}

var (
	activeResourceStyle  = lipgloss.NewStyle().Foreground(SecondaryColor).Width(40).PaddingRight(2)
	defaultResourceStyle = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Width(40).PaddingRight(2)
	cursorStyle          = lipgloss.NewStyle().Foreground(PrimaryColor)
	tagsStyle            = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Width(40).Align(lipgloss.Right)
	selectedMetaStyle    = tagsStyle.Copy().Foreground(SecondaryColor)
	alignStyle           = lipgloss.NewStyle().PaddingLeft(1)
)

func renderMark(mark *Mark, selected bool) string {
	cursor := " "
	style := defaultResourceStyle
	tagsStyle := tagsStyle
	if selected {
		cursor = ">"
		style = activeResourceStyle
		tagsStyle = selectedMetaStyle
	}
	name := style.Render(mark.Name)
	tags := tagsStyle.Render(mark.Tags)
	s := lipgloss.JoinHorizontal(lipgloss.Left, name, tags)
	return lipgloss.JoinHorizontal(lipgloss.Left, cursorStyle.Render(cursor), alignStyle.Render(s))
}
