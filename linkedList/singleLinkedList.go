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
	tail   *node[T]
	length int
}

// New creates a new single linked list
func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head:   nil,
		tail:   nil,
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

	tmp := node[T]{
		value: item,
		next:  nil,
	}

	if list.tail == nil {
		list.head = &tmp
		list.tail = &tmp

		return
	}

	list.tail.next = &tmp
	list.tail = &tmp
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

	curr := list.getAt(idx - 1)

	tmp := node[T]{
		value: item,
		next:  curr.next,
	}

	curr.next = &tmp
}

// Prepend adds a new element at the start of the list
func (list *SinglyLinkedList[T]) Prepend(item T) {
	list.length++

	tmp := node[T]{
		value: item,
		next:  nil,
	}

	if list.head == nil {
		list.head = &tmp
		list.tail = &tmp
		return
	}

	tmp.next = list.head
	list.head = &tmp
}

// Get returns the value at a specified index
func (list *SinglyLinkedList[T]) Get(idx int) *T {
	return &list.getAt(idx).value
}

// GetFirst return the value of the first element
func (list *SinglyLinkedList[T]) GetFirst() *T {
	return &list.head.value
}

// GetLast returns the values of the last element
func (list *SinglyLinkedList[T]) GetLast() *T {
	return &list.tail.value
}

// Remove removes the first element of the list
func (list *SinglyLinkedList[T]) Remove(n node[T]) *T {
	list.length--

	if list.length == 0 {
		out := list.head.value
		list.head = nil
		list.tail = nil
		return &out
	}

	curr := list.head

	for curr.next != &n && curr.next != nil {
		curr = curr.next
	}

	if curr.next == nil {
		return nil
	}

	curr.next = n.next
	n.next = nil

	return &n.value
}

// RemoveAt removes the element at the specified index and return the value
func (list *SinglyLinkedList[T]) RemoveAt(idx int) *T {
	node := list.getAt(idx)

	if node == nil {
		return nil
	}

	return list.Remove(*node)
}

// PrintList prints the list
func (list *SinglyLinkedList[T]) PrintList() {
	curr := list.head

	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (list *SinglyLinkedList[T]) getAt(idx int) *node[T] {
	if idx < 0 || idx > list.length {
		return nil
	}

	if idx == 0 {
		return list.head
	}

	if idx == list.length {
		return list.tail
	}

	curr := list.head
	for range idx {
		curr = curr.next
	}

	return curr
}
