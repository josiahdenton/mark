package internal

import tea "github.com/charmbracelet/bubbletea"

type RefreshMarksMsg struct{}

func refreshMarks() tea.Cmd {
	return func() tea.Msg {
		return RefreshMarksMsg{}
	}
}

type DeleteMarkMsg struct {
	id int
}

func deleteMark(id int) tea.Cmd {
	return func() tea.Msg {
		return DeleteMarkMsg{
			id: id,
		}
	}
}
