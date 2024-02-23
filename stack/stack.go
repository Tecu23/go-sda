package stack


type Stack [T any] struct {
  arr []T
  top int
  cap int
}

func New[T any](cap int) *Stack[T] {
  arr := make([]T, cap)

  stack := Stack[T]{top: 0, cap: cap, arr: arr}

  return &stack
}

func (stack *Stack[T]) Len() int {return 0}

// func (stack *Stack[T]) Peek() T {}
//
// func (stack *Stack[T]) Pop() T {}
//
// func (stack *Stack[T]) Push(item T) {}
