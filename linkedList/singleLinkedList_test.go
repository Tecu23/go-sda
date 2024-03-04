package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	list := New[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	list.AddAtStart(8)
	list.AddAtStart(7)
	list.AddAtStart(6)

	list.PrintList()
	fmt.Print(list.length)
}
