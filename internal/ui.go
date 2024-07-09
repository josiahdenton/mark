package internal

import (
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	contentStyle            = lipgloss.NewStyle().MarginLeft(2)
	formKeyStyle            = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Bold(true)
	listTitleStyle          = lipgloss.NewStyle().Foreground(AccentColor).Bold(true)
	helpKeyStyle            = lipgloss.NewStyle().Foreground(SecondaryGrayColor).Bold(true)
	helpKeyDescriptionStyle = lipgloss.NewStyle().Foreground(SecondaryGrayColor)
)

type Repository interface {
	AllMarks() ([]Mark, error)
	EditMark(m *Mark) error
	AddMark(m *Mark) error
	DeleteMark(id int) (*Mark, error)
}

func New(dbPath string) *Model {
	r, err := ConnectToDB(dbPath)
	if err != nil {
		log.Fatalf("failed to open DB %v", err)
	}

	// setup the list after connecting to DB
	marks, err := r.AllMarks()
	if err != nil {
		log.Fatalf("failed to fetch all marks %v", err)
	}

	l := asList(marks)

	return &Model{
		repository: r,
		keys:       DefaultKeyMapList(),
		marks:      l,
		form:       NewForm(),
		toast:      NewToast(),
	}
}

type Model struct {
	repository   Repository
	keys         KeyMapList
	marks        list.Model
	form         tea.Model
	toast        tea.Model
	showForm     bool
	width        int
	height       int
	deletedMarks []*Mark
}

func (m *Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *Model) View() string {
	var builder strings.Builder
	if m.showForm {
		builder.WriteString(m.form.View())
	} else {
		builder.WriteString(m.marks.View())
	}
	builder.WriteString("\n")
	builder.WriteString(m.toast.View())
	return contentStyle.Render(builder.String())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	// global listeners
	m.toast, cmd = m.toast.Update(msg)
	cmds = append(cmds, cmd)

	// always enable quit from anywhere
	switch msg := msg.(type) {
	case RefreshMarksMsg:
		marks, err := m.repository.AllMarks()
		if err != nil {
			cmds = append(cmds, ShowToast("failed to get all marks", ToastWarn))
			return m, tea.Batch(cmds...)
		}
		items := transformToItems(marks)
		m.marks.SetItems(items)
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.showForm || m.marks.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.Undo):
			if len(m.deletedMarks) == 0 {
				return m, tea.Batch(append(cmds, ShowToast("no deleted marks left to undo", ToastInfo))...)
			}
			lastRemoved := m.deletedMarks[len(m.deletedMarks)-1]
			m.deletedMarks = m.deletedMarks[:len(m.deletedMarks)-1]
			m.repository.AddMark(lastRemoved) // this keeps the same id...
			return m, tea.Batch(append(cmds, ShowToast("re-added deleted mark!", ToastInfo), refreshMarks())...)
		case key.Matches(msg, m.keys.Add):
			m.showForm = true
			return m, tea.Batch(cmds...)
		case key.Matches(msg, m.keys.Edit):
			if selected := m.marks.SelectedItem(); selected != nil {
				mark := selected.(*Mark)
				cmds = append(cmds, editMark(mark)) // add form listens for this message
				m.showForm = true
				return m, tea.Batch(cmds...)
			}
		case key.Matches(msg, m.keys.Delete):
			if selected := m.marks.SelectedItem(); selected != nil {
				mark := selected.(*Mark)
				cmds = append(cmds, deleteMark(mark.Id), ShowToast("deleted mark!", ToastInfo))
				return m, tea.Batch(cmds...)
			}
		case key.Matches(msg, m.keys.Open):
			if selected := m.marks.SelectedItem(); selected != nil {
				mark := selected.(*Mark)
				mark.Open()
				cmds := append(cmds, ShowToast("Opened!", ToastInfo))
				return m, tea.Batch(cmds...)
			}
		case key.Matches(msg, m.keys.Copy):
			if selected := m.marks.SelectedItem(); selected != nil {
				mark := selected.(*Mark)
				err := clipboard.WriteAll(mark.Link)
				if err != nil {
					log.Printf("failed to copy to clipboard: %v", err)
					return m, tea.Batch(append(cmds, ShowToast("failed to copy to clipboard", ToastWarn))...)
				}
				return m, tea.Batch(append(cmds, ShowToast("copied to clipboard!", ToastInfo))...)
			}
		}
	}

	switch msg := msg.(type) {
	case CloseFormMsg:
		log.Printf("m.showForm = %v", m.showForm)
		m.showForm = false
	case MarkCreatedMsg:
		log.Printf("mark created %+v", msg.mark)
		err := m.repository.AddMark(&msg.mark)
		if err != nil {
			log.Printf("%v", err)
			return m, tea.Batch(append(cmds, ShowToast("failed to add mark", ToastWarn))...)
		}
		return m, tea.Batch(append(cmds, refreshMarks())...)
	case MarkModifiedMsg:
		log.Printf("modifying mark %+v", msg.mark)
		err := m.repository.EditMark(&msg.mark)
		if err != nil {
			log.Printf("%v", err)
			return m, tea.Batch(append(cmds, ShowToast("failed to edit mark", ToastWarn))...)
		}
		return m, tea.Batch(append(cmds, refreshMarks())...)
	case DeleteMarkMsg:
		mark, err := m.repository.DeleteMark(msg.id)
		m.deletedMarks = append(m.deletedMarks, mark)
		if err != nil {
			log.Printf("%v", err)
			return m, tea.Batch(append(cmds, ShowToast("failed to delete mark", ToastWarn))...)
		}
		return m, tea.Batch(append(cmds, refreshMarks())...)
	}

	if m.showForm {
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

func asList(marks []Mark) list.Model {
	items := transformToItems(marks)
	l := list.New(items, delegate{}, 30, 15)
	l.Styles.Title = listTitleStyle
	l.Title = ""
	l.SetShowStatusBar(false)
	l.DisableQuitKeybindings()
	l.SetShowHelp(false)
    l.SetShowPagination(false)
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
