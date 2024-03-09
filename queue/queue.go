package queue

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) Enqueue(item T) {
	node := &node[T]{
		value: item,
		next:  nil,
	}

	q.length++

	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() *T {
	if q.head == nil {
		return nil
	}

	q.length--

	head := q.head
	q.head = q.head.next

	if q.IsEmpty() {
		q.tail = nil
	}

	head.next = nil

	return &head.value
}

func (q *Queue[T]) Peek() *T {
	if q.head == nil {
		return nil
	}

	return &q.head.value
}
