// Package stack implements a stack
package stack

import (
	"errors"
	"fmt"
)

// Stack is the structure of the Stack
type Stack [T any] struct {
  arr []T
  top int
}

// New = Stack Constructor Function
func New[T any](cap int) *Stack[T] {
  arr := make([]T, 0, cap)
  stack := Stack[T]{top: -1, arr: arr}

  return &stack
}

// IsEmpty returns whether the stack is empty or not
func (stack *Stack[T]) IsEmpty() bool {
  return stack.top == -1
}
// Len retrieved the length of the stack
func (stack *Stack[T]) Len() int {
  return len(stack.arr) 
}

// Peek retrieved the last item in the stack but does not remove it
func (stack *Stack[T]) Peek() T {
  return stack.arr[stack.top]
}

// Pop retrieves the last item and removes from the stack
func (stack *Stack[T]) Pop() (*T, error) {

  if stack.IsEmpty() {
    return nil, errors.New("stack underflow! cannot pop element")
  }

  item := stack.arr[stack.top]
  stack.top--

  return &item, nil
}

// Push adds an itement to the stack
func (stack *Stack[T]) Push(item *T) {

  stack.arr = append(stack.arr, *item)
  stack.top++
}

// Print displays the stack in order
func (stack *Stack[T]) Print() {
  
  for _, v := range(stack.arr) {
    fmt.Println(v)
  }
}
