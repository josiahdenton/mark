package internal

import (
	// "fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	centerStyle = lipgloss.NewStyle().Align(lipgloss.Center)
	titleStyle  = lipgloss.NewStyle().MarginTop(2).Width(100).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#3a3b5b")).Foreground(PrimaryColor).Bold(true).Align(lipgloss.Center)
	// windowStyle             = lipgloss.NewStyle().PaddingLeft(4)
	formKeyStyle            = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Bold(true)
	listTitleStyle          = lipgloss.NewStyle().Foreground(AccentColor).Bold(true)
	helpKeyStyle            = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Bold(true)
	helpKeyDescriptionStyle = lipgloss.NewStyle().Foreground(SecondaryGrayColor)
	boxStyle                = lipgloss.NewStyle().PaddingLeft(1).Width(100).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#3a3b5b"))
)

type Repository interface {
	AllMarks() []Mark
	EditMark(m *Mark) error
	AddMark(m *Mark)
	DeleteMark(id int) error
}

func New() *Model {
	// TODO: use the db value
	// r := ConnectToDB()
	r := ConnectToInMemory()

	// setup the list after connecting to DB
	l := marksAsList(r.AllMarks())

	return &Model{
		repository: r,
		keys:       DefaultKeyMapList(),
		marks:      l,
		form:       NewForm(),
		toast:      NewToast(),
	}
}

type Model struct {
	repository Repository
	keys       KeyMapList
	marks      list.Model
	form       tea.Model
	toast      tea.Model
	showForm   bool
	width      int
	height     int
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m *Model) View() string {
	// fmt.Println("m.showForm = %v", m.showForm)
	var builder strings.Builder
	builder.WriteString(titleStyle.Render("Mark"))
	builder.WriteString("\n")
	if m.showForm {
		builder.WriteString(m.form.View())
	} else {
		builder.WriteString(boxStyle.Render(m.marks.View()))
	}
	// s := boxStyle.Render(builder.String())
	// return windowStyle.Width(m.width).Height(m.height).Render(s)
	return centerStyle.Width(m.width).Height(m.height).Render(builder.String())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	// always enable quit from anywhere
	switch msg := msg.(type) {
	case RefreshMarksMsg:
		marks := transformToItems(m.repository.AllMarks())
		m.marks.SetItems(marks)
		return m, tea.Batch(cmds...)
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, tea.Batch(cmds...)
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}
	}

	// global listeners
	m.toast, cmd = m.toast.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.showForm { // don't allow key input with form
			break
		}

		switch {
		case key.Matches(msg, m.keys.Add):
			m.showForm = true
			return m, tea.Batch(cmds...)
		case key.Matches(msg, m.keys.Edit):
			selected := m.marks.SelectedItem().(*Mark)
			cmds = append(cmds, editMark(selected)) // add form listens for this message
			m.showForm = true
			return m, tea.Batch(cmds...)
		case key.Matches(msg, m.keys.Delete):
			selected := m.marks.SelectedItem().(*Mark)
			cmds = append(cmds, deleteMark(selected.Id))
			return m, tea.Batch(cmds...)
		}
	}

	switch msg := msg.(type) {
	case CloseFormMsg:
		m.showForm = false
		// fmt.Println("CLOSING FORM")
	case MarkCreatedMsg:
		m.repository.AddMark(msg.mark)
		return m, tea.Batch(append(cmds, refreshMarks())...)
	case MarkModifiedMsg:
		m.repository.EditMark(msg.mark)
		return m, tea.Batch(append(cmds, refreshMarks())...)
	case DeleteMarkMsg:
		m.repository.DeleteMark(msg.id)
		return m, tea.Batch(append(cmds, refreshMarks())...)
	}

	// child components BUG: this never stops showing form because the check to close is never hit...
	if m.showForm {
		// TODO: implement a Blur/Focus for form...
		// also, if I want an explicit type I can just
		// explicitly type cast this to (tea.Model) as it does fulfill that interface
		m.form, cmd = m.form.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	if !m.showForm {
		m.marks, cmd = m.marks.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func marksAsList(marks []Mark) list.Model {
	items := transformToItems(marks)
	l := list.New(items, resourceDelegate{}, 30, 15)
	l.Styles.Title = listTitleStyle
	l.Title = ""
	l.SetShowStatusBar(false)
	l.DisableQuitKeybindings()
	l.SetShowHelp(false)
	return l
}

func transformToItems(marks []Mark) []list.Item {
	items := make([]list.Item, len(marks))
	for i, mark := range marks {
		item := &mark
		items[i] = item
	}
	return items
}
