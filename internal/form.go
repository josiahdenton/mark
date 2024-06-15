package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	formStyle      = lipgloss.NewStyle().PaddingTop(5)
	formTitleStyle = lipgloss.NewStyle().Foreground(AccentColor).Bold(true)
	formLabelStyle = lipgloss.NewStyle().Foreground(AccentColor)
	boxFormStyle   = lipgloss.NewStyle().Padding(2).Width(100).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#3a3b5b"))
)

type FormIndex int

const (
	Name = iota
	Link
	Tags
)

// results
type CloseFormMsg struct{}

func closeForm() tea.Cmd {
	return func() tea.Msg {
		return CloseFormMsg{}
	}
}

type MarkModifiedMsg struct {
	mark *Mark
}

func markModified(m *Mark) tea.Cmd {
	return func() tea.Msg {
		return MarkModifiedMsg{m}
	}
}

type MarkCreatedMsg struct {
	mark *Mark
}

func markCreated(m *Mark) tea.Cmd {
	return func() tea.Msg {
		return MarkCreatedMsg{m}
	}
}

// actions
type EditMarkMsg struct {
	mark *Mark
}

func editMark(mark *Mark) tea.Cmd {
	return func() tea.Msg {
		return EditMarkMsg{
			mark: mark,
		}
	}
}

type AddMarkMsg struct {
	mark *Mark
}

func addMark(mark *Mark) tea.Cmd {
	return func() tea.Msg {
		return AddMarkMsg{
			mark: mark,
		}
	}
}

func NewForm() *FormModel {
	name := textinput.New()
	name.Focus()
	name.Width = 60
	name.CharLimit = 60
	name.Prompt = "Name: "
	name.PromptStyle = formLabelStyle
	name.Placeholder = "..."

	name.Validate = func(s string) error {
		if len(strings.Trim(s, " \n")) < 1 {
			return fmt.Errorf("name missing")
		}
		return nil
	}

	link := textinput.New()
	link.Blur()
	link.Width = 60
	link.CharLimit = 300
	link.Prompt = "Link: "
	link.PromptStyle = formLabelStyle
	link.Placeholder = "..."

	link.Validate = func(s string) error {
		if len(strings.Trim(s, " \n")) < 1 {
			return fmt.Errorf("link missing")
		}
		return nil
	}

	tags := textinput.New()
	tags.Blur()
	tags.Width = 60
	tags.CharLimit = 300
	tags.Prompt = "Tags: "
	tags.PromptStyle = formLabelStyle
	tags.Placeholder = "..."

	return &FormModel{
		name:        name,
		link:        link,
		tags:        tags,
		activeInput: Name,
		keys:        DefaultKeyMapForm(),
	}
}

type FormModel struct {
	name        textinput.Model
	link        textinput.Model
	tags        textinput.Model
	m           Mark
	activeInput FormIndex
	keys        KeyMapForm
}

func (f *FormModel) Init() tea.Cmd {
	return nil
}

func (f *FormModel) View() string {
	var b strings.Builder
	b.WriteString(formTitleStyle.Render("Add Mark"))
	b.WriteString("\n\n")
	b.WriteString(f.name.View())
	b.WriteString("\n\n")
	b.WriteString(f.link.View())
	b.WriteString("\n\n")
	b.WriteString(f.tags.View())
	return boxFormStyle.Render(b.String())
}

func (f *FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// -- init --
	switch msg := msg.(type) {
	case EditMarkMsg:
		f.m.Id = msg.mark.Id
		f.m.Name = msg.mark.Name
		f.m.Link = msg.mark.Link
		f.m.Tags = msg.mark.Tags
	}

	// -- key messags --
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, f.keys.Close):
			cmds = append(cmds, closeForm())
		case key.Matches(msg, f.keys.NextInput):
			f.focusNextField()
			return f, tea.Batch(cmds...)
		case key.Matches(msg, f.keys.Submit):
			if f.activeInput < Tags {
				f.focusNextField()
				return f, tea.Batch(cmds...)
			}

			// assign form values and signal parent component
			f.m.Name = f.name.Value()
			f.m.Link = f.link.Value()
			f.m.Tags = f.link.Value()

			// send off
			if f.m.Id != 0 { // existing mark
				cmds = append(cmds, markModified(&f.m), closeForm())
			} else { // new mark
				cmds = append(cmds, markCreated(&f.m), closeForm())
			}
		}

	}

	// form field input
	// -- name --
	f.name, cmd = f.name.Update(msg)
	cmds = append(cmds, cmd)
	// -- link --
	f.link, cmd = f.link.Update(msg)
	cmds = append(cmds, cmd)
	// -- tags --
	f.tags, cmd = f.tags.Update(msg)
	cmds = append(cmds, cmd)

	return f, tea.Batch(cmds...)
}

func (f *FormModel) focusNextField() {
	switch f.activeInput {
	case Name:
		f.activeInput = Link
		f.name.Blur()
		f.link.Focus()
	case Link:
		f.activeInput = Tags
		f.link.Blur()
		f.tags.Focus()
	case Tags:
		f.activeInput = Name
		f.tags.Blur()
		f.name.Focus()
	}
}

func validateForm(errs ...error) tea.Cmd {
	for _, err := range errs {
		if err != nil {
			return ShowToast(fmt.Sprintf("%v", err), ToastWarn)
		}
	}
	return nil
}
