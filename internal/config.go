package internal

import "github.com/charmbracelet/bubbles/key"

type KeyMapList struct {
	Up     key.Binding
	Down   key.Binding
	Add    key.Binding
	Open   key.Binding
	Copy   key.Binding
	Edit   key.Binding
	Delete key.Binding
	Undo   key.Binding
	Help   key.Binding
	Quit   key.Binding
}

func DefaultKeyMapList() KeyMapList {
	return KeyMapList{
		Up:     key.NewBinding(key.WithKeys("k", "up", "ctrl+p"), key.WithHelp("k/↑/ctrl+p", "move up")),
		Down:   key.NewBinding(key.WithKeys("j", "down", "ctrl+n"), key.WithHelp("j/↓/ctrl+n", "move down")),
		Add:    key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add mark")),
		Open:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "open mark")),
		Copy:   key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "copy mark")),
		Edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit mark")),
		Delete: key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete mark")),
		Undo:   key.NewBinding(key.WithKeys("u"), key.WithHelp("u", "undo delete")),
		Help:   key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "toggle help")),
		Quit:   key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
	}
}

func (k KeyMapList) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMapList) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Open, k.Copy},
		{k.Add, k.Edit, k.Delete, k.Undo},
		{k.Help, k.Quit},
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
