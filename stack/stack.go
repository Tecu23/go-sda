package stack


type Stack [T any] []T

func New[T any]() *Stack[T] {
  return &Stack[T]{}
}

func (stack *Stack[T]) Len() int {return 0}

// func (stack *Stack[T]) Peek() T {}
//
// func (stack *Stack[T]) Pop() T {}
//
// func (stack *Stack[T]) Push(item T) {}
