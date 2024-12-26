package LinkedList

import (
	"fmt"
	"strconv"
	"strings"
)

type deque interface {
	Find(key string) ListEntry
	Clear()
	Get(index int) ListEntry
	Set(index int, val int, key string)
	Append(key string, val int)
	Prepend(key string, val int)
	Delete(key string)
	PopLeft() ListEntry
	PopRight() ListEntry
	ToString() string
	Length() int
}

type ListEntry struct {
	val  int
	key  string
	next *ListEntry
	prev *ListEntry
}

type DoubleLinkedList struct {
	head *ListEntry
	tail *ListEntry
}

func (dList *DoubleLinkedList) Lenght() int {
	if dList.head == nil && dList.tail == nil {
		return 0
	}
	curr := dList.head
	counter := 0
	for curr != nil {
		counter++
		curr = curr.next
	}
	return counter
}

func (dList *DoubleLinkedList) Append(key string, val int) {
	if dList.head == nil {
		dList.head = &ListEntry{val, key, dList.tail, dList.head}
	} else {
		curr := dList.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &ListEntry{val, key, dList.tail, curr.next}
	}
}

func (dList *DoubleLinkedList) ToString() string {
	head := dList.head
	if head == nil {
		return "[]"
	} else {
		var result []string
		curr := dList.head
		for curr != nil {
			result = append(result, "{key : "+curr.key+","+strconv.Itoa(curr.val)+"}")
			curr = curr.next
		}

		return "[" + strings.Join(result, ", ") + "]"
	}
}

func (dList *DoubleLinkedList) Find(key string) *ListEntry {
	curr := dList.head

	if curr == nil {
		return nil
	} else {
		for curr != nil {
			if curr.key == key {
				return curr
			}
		}
		return nil
	}
}

func (dList *DoubleLinkedList) Prepend(key string, val int) {
	head := dList.head
	if head == nil {
		// head == tail if only one element
		head = &ListEntry{val, key, nil, head}
	} else {
		temp := head
		dList.head = &ListEntry{val, key, nil, temp}
	}
}

func (dList *DoubleLinkedList) Get(index int) *ListEntry {
	curr := dList.head
	counter := 0
	for curr != nil {
		if counter == index {
			return curr
		}
	}
	return nil
}

func (dList *DoubleLinkedList) Set(index int, key int, val string) {

}

func main() {
	fmt.Println("works")
	first := ListEntry{1, "1", nil, nil}
	second := ListEntry{2, "2", nil, &first}
	first.next = &second
	l := DoubleLinkedList{&first, &second}
	fmt.Println(l.ToString())
}
