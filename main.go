// main
package main

import (
	"fmt"

	"github.com/Tecu23/go-sda/array"
  "github.com/Tecu23/go-sda/stack"
)
func main() {
  testStack()
}

func testQuickSort() {

	tmp := []int{12, 11, 10, 9, 8, 7}

	array.QuickSort(&tmp, 0, len(tmp)-1)

	fmt.Println(tmp)
}

func testStack() {

  stk := stack.New[int](10)

  fmt.Print(stk)

  fmt.Println(stk.IsEmpty())

	tmp := []int{12, 11, 10, 9, 8, 7}
  for _, s := range tmp {
    stk.Push(&s)
  }

  fmt.Print(stk.Peek())

  stk.Print()
}
