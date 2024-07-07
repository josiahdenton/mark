package internal

import (
	"os/exec"
	"runtime"
)

type Mark struct {
	Id   int    `db:"mark_id"`
	Name string `db:"name"`
	Link string `db:"link"`
	Tags string `db:"tags"`
}

func (m *Mark) FilterValue() string {
    return m.Name + m.Tags
}

func (n *Mark) Open() bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], n.Link)...)
	return cmd.Start() == nil
}
