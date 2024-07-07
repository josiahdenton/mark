package internal

import (
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	moveFocusPressed = tea.KeyMsg{
		Type: tea.KeyTab,
	}

	submitPressed = tea.KeyMsg{
		Type: tea.KeyEnter,
	}

    escPressed = tea.KeyMsg{
        Type: tea.KeyEsc,
    }

	inputPressed = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'d'},
	}
)

func TestMoveFocusOnTab(t *testing.T) {
	f := NewForm()
	f.Update(moveFocusPressed)
	if f.activeInput != Link {
		t.Fatalf("failed: expected '%d', got '%d'", Link, f.activeInput)
	}
}

func TestTextInput(t *testing.T) {
	f := NewForm()
	f.Update(inputPressed)
	if f.name.Value() != "d" {
		t.Fatalf("failed: expected 'd' got '%s'", f.name.Value())
	}
}

// func TestFormSubmit(t *testing.T) {
// 	f := NewForm()
// 	// name
// 	f.Update(inputPressed)
// 	f.Update(moveFocusPressed)
// 	// link
// 	f.Update(inputPressed)
// 	f.Update(moveFocusPressed)
// 	// tags
// 	f.Update(inputPressed)
// 	_, cmd := f.Update(submitPressed)
// 	msg := cmd()
// 	if reflect.TypeOf(msg).Name() != reflect.TypeOf(MarkCreatedMsg{}).Name() {
// 		t.Fatalf("failed: submitPressed did not return ")
// 	}
// }
//
func TestCloseForm(t *testing.T) {
    f := NewForm()
    _, cmd := f.Update(escPressed)
	msg := cmd()
	if reflect.TypeOf(msg).Name() != reflect.TypeOf(CloseFormMsg{}).Name() {
		t.Fatalf("failed: submitPressed did not return ")
	}
}
