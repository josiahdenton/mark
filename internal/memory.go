package internal

import (
	"errors"
	"slices"
)

var (
	InMemNotFoundError = errors.New("failed to find mark")
)

func ConnectToInMemory() *InMemory {
	return &InMemory{
		marks: []Mark{
			{
				Id:   999,
				Name: "Google",
				Link: "https://www.google.com",
				Tags: "search-engine",
			},
		},
	}
}

type InMemory struct {
	marks     []Mark
	lastestId int
}

func (im *InMemory) AllMarks() []Mark {
	return im.marks
}

func (im *InMemory) EditMark(m *Mark) error {
	// replace the Mark
	for i, mark := range im.marks {
		if mark.Id == m.Id {
			im.marks[i] = *m
			return nil
		}
	}
	return InMemNotFoundError
}
func (im *InMemory) AddMark(m *Mark) {
    im.lastestId += 1
    m.Id = im.lastestId
	im.marks = append(im.marks, *m)
}
func (im *InMemory) DeleteMark(id int) error {
	for i, mark := range im.marks {
		if mark.Id == id {
			im.marks = slices.Delete(im.marks, i, i+1)
			return nil
		}
	}
	return InMemNotFoundError
}
