package LinkedList

import (
	"testing"
)

func createDLL() *DoubleLinkedList {
	first := &ListEntry{1, "1", nil, nil}
	second := &ListEntry{2, "2", nil, first}
	first.next = second
	third := &ListEntry{3, "3", nil, second}
	second.next = third
	dList := &DoubleLinkedList{first, third}
	return dList
}

func TestLegth(t *testing.T) {
	dList := createDLL()
	got := dList.Lenght()
	expected := 3
	if got != expected {
		t.Errorf("Length mismatch: got %d, want %d", got, expected)
	}
}

func TestToString(t *testing.T) {
	dList := createDLL()
	got := dList.ToString()
	expect := "[{key : 1,1}, {key : 2,2}, {key : 3,3}]"

	if got != expect {
		t.Errorf("this is the string ser. : %s", got)
	}
}

func TestAppend(t *testing.T) {
	dList := createDLL()
	dList.Append("4", 4)
	if dList.tail.key != "4" {
		t.Errorf("this is the string ser. : %s", got)
	}
	length := 
	if dList.Lenght() != 4 {
		t.Errorf("not the right length after appending: %s",)
	}
}
