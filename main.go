package main

import (
	"fmt"

	"github.com/Tecu23/go-sda/array"
)

func main() {

	tmp := []int{12, 11, 10, 9, 8, 7}

	array.QuickSort(&tmp, 0, len(tmp)-1)

	fmt.Println(tmp)
}
