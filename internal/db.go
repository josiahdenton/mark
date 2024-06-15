package internal

import (
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

// NOTE: can use the following to grab latest ID to increment in DB
// mark_id value 0 should mean a new ID is needed
// if the SELECT returns nothing, set mark_id as 1  (we want the smallest id to be 1)
// SELECT * 
// FROM    TABLE
// WHERE   ID = (SELECT MAX(ID)  FROM TABLE);

func ConnectToDB() *MarkDatabase {
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail;
	db.MustExec(schema)
	return &MarkDatabase{
		db: db,
	}
}

type MarkDatabase struct {
	db *sqlx.DB
}

func (m *MarkDatabase) AllMarks() []Mark {
	marks := []Mark{}
	err := m.db.Select(&marks, "SELECT * FROM marks ORDER BY name ASC")
	if err != nil {
		log.Fatalf("AllMarks failed: %v", err)
	}
	return marks
}

func (m *MarkDatabase) EditMark(mark *Mark) {
	_, err := m.db.Exec("UPDATE marks SET name=$1, link=$2, tags=$3 WHERE mark_id=$4", mark.Name, mark.Link, mark.Tags, mark.Id)
	if err != nil {
		log.Fatalf("EditMark failed: %v", err)
	}
}

func (m *MarkDatabase) AddMark(mark *Mark) {
	// mark_id should be auto-assigned a value on insert
	_, err := m.db.Exec("INSERT INTO marks (name, link, tags) VALUES ($1, $2, $3)", mark.Name, mark.Link, mark.Tags)
	if err != nil {
		log.Fatalf("EditMark failed: %v", err)
	}
}

func (m *MarkDatabase) DeleteMark(id int) {
	_, err := m.db.Exec("DELETE FROM marks WHERE mark_id=$1", id)
	if err != nil {
		log.Fatalf("EditMark failed: %v", err)
	}
}
