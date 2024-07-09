package internal

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type delegate struct{}

func (d delegate) Height() int  { return 1 }
func (d delegate) Spacing() int { return 0 }
func (d delegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}
func (d delegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	mark, ok := item.(*Mark)
	if !ok {
		return
	}
	fmt.Fprintf(w, renderMark(mark, index == m.Index()))
}

var (
	activeStyle  = lipgloss.NewStyle().Foreground(SecondaryColor).Width(30).PaddingRight(2)
	defaultStyle = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Width(30).PaddingRight(2)
	cursorStyle          = lipgloss.NewStyle().Foreground(PrimaryColor)
	tagsStyle            = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Width(30).Align(lipgloss.Right)
	selectedMetaStyle    = tagsStyle.Copy().Foreground(SecondaryColor)
	alignStyle           = lipgloss.NewStyle().PaddingLeft(1)
)

func renderMark(mark *Mark, selected bool) string {
	cursor := " "
	style := defaultStyle
	tagsStyle := tagsStyle
	if selected {
		cursor = ">"
		style = activeStyle
		tagsStyle = selectedMetaStyle
	}
	name := style.Render(mark.Name)
	tags := tagsStyle.Render(mark.Tags)
	s := lipgloss.JoinHorizontal(lipgloss.Left, name, tags)
	return lipgloss.JoinHorizontal(lipgloss.Left, cursorStyle.Render(cursor), alignStyle.Render(s))
}
