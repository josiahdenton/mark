package internal

import "github.com/charmbracelet/bubbles/key"

type KeyMapList struct {
	Up     key.Binding
	Down   key.Binding
	Add    key.Binding
	Edit   key.Binding
	Delete key.Binding
	Quit   key.Binding
}

func DefaultKeyMapList() KeyMapList {
	return KeyMapList{
		Up:     key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("k", "move up")),
		Down:   key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("j", "move down")),
		Add:    key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add mark")),
		Edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit mark")),
		Delete: key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete mark")),
		Quit:   key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+q", "quit")),
	}
}

type KeyMapForm struct {
	NextInput key.Binding
	Submit    key.Binding
	Close     key.Binding
}

func DefaultKeyMapForm() KeyMapForm {
	return KeyMapForm{
		NextInput: key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next input")),
		Submit:    key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
		Close:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "close")),
	}
}
