package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/josiahdenton/mark/internal"
)

func main() {
	if err := Run(); err != nil {
		fmt.Printf("failed to brew tea: %v", err)
		os.Exit(1)
	}
}

const MarkDir = "mark"

func Run() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	markHome := fmt.Sprintf("%s/.%s", home, MarkDir)

	if _, err := os.Stat(markHome); errors.Is(err, os.ErrNotExist) {
        // create the dir if not exists and try again
		err := os.Mkdir(markHome, os.ModePerm)
		if err != nil {
			return err
		}
		return Run()
	} else if err != nil {
		log.Fatalf("failed to setup home dir, reason %v", err)
	} else {
		f, err := os.OpenFile(fmt.Sprintf("%s/log", markHome), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("--------------- Mark ---------------")
		log.Println("starting program")

		p := tea.NewProgram(internal.New(fmt.Sprintf("%s/marks.db", markHome)))
		if _, err := p.Run(); err != nil {
			return err
		}
	}

	return nil
}
