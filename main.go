package main

import (
	"fmt"

	"github.com/Tecu23/go-sda/array"
)

func main() {

	tmp := []int{203, 4, 10, 12, 7, 8}

	array.SimpleSort(&tmp)

	fmt.Println(tmp)
}
