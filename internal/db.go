package internal

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS marks (
    mark_id integer PRIMARY KEY,
    name text,
    link text,
    tags text
);
`

func ConnectToDB(dbName string) (*MarkDatabase, error) {
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s, %w", dbName, err)
	}

	// exec the schema or fail;
	db.MustExec(schema)

	return &MarkDatabase{
		db: db,
	}, nil
}

type MarkDatabase struct {
	db *sqlx.DB
}

// TODO: convert from fatalf

func (m *MarkDatabase) AllMarks() ([]Mark, error) {
	marks := []Mark{}
	err := m.db.Select(&marks, "SELECT * FROM marks ORDER BY mark_id ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to select all marks %w", err)
	}
	return marks, nil
}

func (m *MarkDatabase) maxId() (int, error) {
	marks := []Mark{}
	err := m.db.Select(&marks, "SELECT * FROM marks WHERE mark_id = (SELECT MAX(mark_id) FROM marks)")
	if err != nil {
		return 0, fmt.Errorf("could not get max id %w", err)
	}

	if len(marks) < 1 {
		return 0, nil
	}

	return marks[0].Id, nil
}

func (m *MarkDatabase) EditMark(mark *Mark) error {
	_, err := m.db.Exec("UPDATE marks SET name=$1, link=$2, tags=$3 WHERE mark_id=$4", mark.Name, mark.Link, mark.Tags, mark.Id)
	if err != nil {
		return fmt.Errorf("EditMark failed: %w", err)
	}
	return nil
}

func (m *MarkDatabase) AddMark(mark *Mark) error {
	maxId, err := m.maxId()
	if err != nil {
		log.Fatalf("could not get max id %v", err)
	}
	// our ID should be max + 1
	if mark.Id == 0 {
		mark.Id = maxId + 1
	}
	_, err = m.db.Exec("INSERT INTO marks (mark_id, name, link, tags) VALUES ($1, $2, $3, $4)", mark.Id, mark.Name, mark.Link, mark.Tags)
	if err != nil {
		return fmt.Errorf("failed to edit mark: %w", err)
	}
	return nil
}

func (m *MarkDatabase) DeleteMark(id int) (*Mark, error) {
	mark, err := m.mark(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get mark for delete: %w", err)
	}
	_, err = m.db.Exec("DELETE FROM marks WHERE mark_id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete mark: %w", err)
	}
	return mark, nil
}

func (m *MarkDatabase) mark(id int) (*Mark, error) {
	var marks []Mark
	err := m.db.Select(&marks, "SELECT * FROM marks WHERE mark_id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mark %w", err)
	}

	if len(marks) == 0 {
		return nil, fmt.Errorf("no mark found with id %d", id)
	}

	return &marks[0], nil
}
