package internal

import "testing"

func TestConnect(t *testing.T) {
	_, err := ConnectToDB(":memory:")

	if err != nil {
		t.Fatalf("failed to connection to DB, %v", err)
	}
}

func TestAdd(t *testing.T) {
	db, err := ConnectToDB(":memory:")

	if err != nil {
		t.Fatalf("failed to connection to DB, %v", err)
	}

	db.AddMark(&Mark{
		Id:   1,
		Name: "Test",
		Link: "Test",
		Tags: "Test",
	})

	marks, err := db.AllMarks()

	if len(marks) < 1 {
		t.Fatalf("failed to add/fetch marks")
	}
}

func TestEdit(t *testing.T) {
	db, err := ConnectToDB(":memory:")

	if err != nil {
		t.Fatalf("failed to connection to DB, %v", err)
	}

	db.AddMark(&Mark{
		Id:   1,
		Name: "Test",
		Link: "Test",
		Tags: "Test",
	})
}
