package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var (
	warnStatusStyle = lipgloss.NewStyle().Foreground(PrimaryGrayColor).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#D120AF")).Width(25).Align(lipgloss.Center)
	infoStatusStyle = lipgloss.NewStyle().Foreground(PrimaryGrayColor).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#2dd4bf")).Width(25).Align(lipgloss.Center)
)

type ShowToastMsg struct {
	Message string
	Toast   ToastType
}

const (
	ToastInfo = iota
	ToastWarn
)

type ToastType = int

func ShowToast(message string, toast ToastType) tea.Cmd {
	return func() tea.Msg {
		return ShowToastMsg{Message: message, Toast: toast}
	}
}

func NewToast() ToastModel {
	return ToastModel{}
}

type ToastModel struct {
	message string
	toast   ToastType
}

func (m ToastModel) Init() tea.Cmd {
	return nil
}

func (m ToastModel) View() string {
	if len(m.message) > 0 && m.toast == ToastInfo {
		return infoStatusStyle.Render(m.message)
	} else if len(m.message) > 0 && m.toast == ToastWarn {
		return warnStatusStyle.Render(m.message)
	}
	return ""
}

func (m ToastModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case ShowToastMsg:
		m.message = msg.Message
		m.toast = msg.Toast
		cmd = clearStatus()
	case clearStatusMsg:
		m.message = ""
	}

	return m, cmd
}

type clearStatusMsg struct{}

func clearStatus() tea.Cmd {
	return tea.Tick(time.Second*5, func(_ time.Time) tea.Msg {
		return clearStatusMsg{}
	})
}
