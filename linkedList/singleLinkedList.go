package linkedList

import "fmt"

type node[T any] struct {
	value T
	next  *node[T]
}

type SinglyLinkedList[T any] struct {
	head   *node[T]
	length int
}

func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head:   nil,
		length: 0,
	}
}

func (list *SinglyLinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *SinglyLinkedList[T]) Add(item T) {
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

func (list *SinglyLinkedList[T]) AddAtIndex(idx int, item T) {}

func (list *SinglyLinkedList[T]) AddAtStart(item T) {
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
