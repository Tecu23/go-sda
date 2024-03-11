package binarytree_test

import (
	"fmt"
	"testing"

	binarytree "github.com/Tecu23/go-sda/binaryTree"
)

func Test(t *testing.T) {
	fmt.Print("hehe")
	root := binarytree.CreateRoot(32)

	fmt.Printf("%v", (*root).Value)
}
