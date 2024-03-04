package linkedList

type node[T any] struct {
	value T
	next  *node[T]
}

type SinglyLinkedList[T any] struct {
	head   *node[T]
	length int
}

func New[T any]() *SinglyLinkedList[T] {
	return nil
}

func (list *SinglyLinkedList[T]) Add(item T) {}

func (list *SinglyLinkedList[T]) AddAtIndex(idx int) {}

func (list *SinglyLinkedList[T]) AddAtStart() {}

func (list *SinglyLinkedList[T]) GetFirst() *T { return nil }

func (list *SinglyLinkedList[T]) GetLast() *T { return nil }

func (list *SinglyLinkedList[T]) Remove() *T { return nil }

func (list *SinglyLinkedList[T]) RemoveAt(idx int) *T { return nil }

func (list *SinglyLinkedList[T]) PrintList() {}
