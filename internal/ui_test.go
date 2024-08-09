package internal

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	addPressed = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'a'},
	}

	editPressed = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'e'},
	}

	deletePressed = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'d'},
	}
)

func TestAddMarkFlow(t *testing.T) {
	//m := New()
	//m.Update(addPressed)
}
