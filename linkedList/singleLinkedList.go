// Package linkedlist implements linked lists data structures
package linkedlist

import "fmt"

type node[T any] struct {
	value T
	next  *node[T]
}

// SinglyLinkedList type
type SinglyLinkedList[T any] struct {
	head   *node[T]
	length int
}

// New creates a new single linked list
func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head:   nil,
		length: 0,
	}
}

// IsEmpty checks if the list is empty
func (list *SinglyLinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

// Append adds a new item at the end of the list
func (list *SinglyLinkedList[T]) Append(item T) {
	list.length++

	if list.head == nil {
		list.head = &node[T]{
			value: item,
			next:  nil,
		}
		return
	}

	curr := list.head

	for curr.next != nil {
		curr = curr.next
	}

	tmp := node[T]{
		value: item,
		next:  nil,
	}

	curr.next = &tmp
}

// InsertAt adds a new element at the provided index
func (list *SinglyLinkedList[T]) InsertAt(idx int, item T) {
	if idx > list.length || idx < 0 {
		return
	}

	if idx == list.length {
		list.Append(item)
	} else if idx == 0 {
		list.Prepend(item)
	}

	list.length++

	curr := list.head

	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	tmp := node[T]{
		value: item,
		next:  curr.next,
	}

	curr.next = &tmp
}

// Prepend adds a new element at the start of the list
func (list *SinglyLinkedList[T]) Prepend(item T) {
	list.length++

	if list.head == nil {
		list.head = &node[T]{
			value: item,
			next:  nil,
		}
		return
	}

	tmp := node[T]{
		value: item,
		next:  list.head,
	}

	list.head = &tmp
}

func (list *SinglyLinkedList[T]) GetFirst() *T { return nil }

func (list *SinglyLinkedList[T]) GetLast() *T { return nil }

func (list *SinglyLinkedList[T]) Remove() *T { return nil }

func (list *SinglyLinkedList[T]) RemoveAt(idx int) *T { return nil }

func (list *SinglyLinkedList[T]) PrintList() {
	curr := list.head

	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}
