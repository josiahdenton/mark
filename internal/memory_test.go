package internal

// import (
// 	"slices"
// 	"testing"
// )

// func TestAddMark(t *testing.T) {
// 	im := ConnectToInMemory()
// 	m := &Mark{
// 		Id:   0,
// 		Name: "",
// 		Link: "",
// 		Tags: "",
// 	}
// 	im.AddMark(m)
// 	marks := im.AllMarks()
// 	if !slices.Contains(marks, *m) {
// 		t.Fatalf("expected %+v in %+v", m, marks)
// 	}
//
// }
//
// func TestEditMark(t *testing.T) {
// 	im := ConnectToInMemory()
// 	m := &Mark{
// 		Id:   0,
// 		Name: "",
// 		Link: "",
// 		Tags: "",
// 	}
// 	im.AddMark(m)
// 	m2 := *m
// 	modifiedName := "Test"
// 	m2.Name = modifiedName
// 	im.EditMark(&m2)
// 	marks := im.AllMarks()
// 	if marks[1].Name != modifiedName {
// 		t.Fatalf("edit mark failed")
// 	}
//
// }
//
// func TestDeleteMark(t *testing.T) {
// 	im := ConnectToInMemory()
// 	m := &Mark{
// 		Id:   0,
// 		Name: "",
// 		Link: "",
// 		Tags: "",
// 	}
// 	im.AddMark(m)
// 	err := im.DeleteMark(m.Id)
// 	if err != nil {
// 		t.Fatalf("failed %v", err)
// 	}
// 	marks := im.AllMarks()
// 	if slices.Contains(marks, *m) {
// 		t.Fatalf("delete failed")
// 	}
// }
//
// // TODO: these are broken now, should fix with the auto-inc ids
// func TestDeleteManyMarks(t *testing.T) {
// 	im := ConnectToInMemory()
// 	m := &Mark{
// 		Id:   0,
// 		Name: "",
// 		Link: "",
// 		Tags: "",
// 	}
// 	m2 := &Mark{
// 		Id: 1,
// 	}
// 	m3 := &Mark{
// 		Id: 2,
// 	}
// 	im.AddMark(m)
// 	im.AddMark(m2)
// 	im.AddMark(m3)
// 	err := im.DeleteMark(m2.Id)
// 	if err != nil {
// 		t.Fatalf("failed %v", err)
// 	}
// 	marks := im.AllMarks()
// 	if slices.Contains(marks, *m2) {
// 		t.Fatalf("delete failed to remove %d", m2.Id)
// 	}
// }
