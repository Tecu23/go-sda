// Package binarysearchtree is a package for BST
package binarysearchtree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Node type provides a single node
type Node[T constraints.Ordered] struct {
	value  T
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
}

// BinarySearchTree type used for creating and using bst
type BinarySearchTree[T constraints.Ordered] struct {
	root *Node[T]
}

// New creates a new BST
func New[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

// Insert adds a new Node to the BST in the correct position
func (bt *BinarySearchTree[T]) Insert(value T) {
	tmp := &Node[T]{
		value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	curr := bt.root
	var y *Node[T]

	for curr != nil {
		y = curr
		if tmp.value < curr.value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	tmp.parent = y

	if y == nil {
		bt.root = tmp
	} else if tmp.value < y.value {
		y.left = tmp
	} else {
		y.right = tmp
	}
}

// Search finds a value in the BST
func (bt *BinarySearchTree[T]) Search(value T) *Node[T] {
	curr := bt.root

	for curr != nil && curr.value != value {
		if value < curr.value {
			curr = curr.left
		} else {
			curr = curr.left
		}
	}

	return curr
}

// PrintInorder prints the tree in order
func (bt *BinarySearchTree[T]) PrintInorder() {
	inorderNodeWalk(bt.root)
}

// PrintPreorder prints the tree in pre-order
func (bt *BinarySearchTree[T]) PrintPreorder() {
	inorderNodeWalk(bt.root)
}

// PrintPostorder prints the tree in post-order
func (bt *BinarySearchTree[T]) PrintPostorder() {
	inorderNodeWalk(bt.root)
}

// Private Methods
func recursiveSearch[T constraints.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil || node.value == value {
		return node
	}
	if value < node.value {
		return recursiveSearch(node.left, value)
	}

	return recursiveSearch(node.right, value)
}

func inorderNodeWalk[T constraints.Ordered](node *Node[T]) {
	if node != nil {
		inorderNodeWalk((*node).left)
		fmt.Println((*node).value)
		inorderNodeWalk((*node).right)
	}
}

func preorderNodeWalk[T constraints.Ordered](node *Node[T]) {
	if node != nil {
		fmt.Println((*node).value)
		inorderNodeWalk((*node).left)
		inorderNodeWalk((*node).right)
	}
}

func postorderNodeWalk[T constraints.Ordered](node *Node[T]) {
	if node != nil {
		inorderNodeWalk((*node).left)
		inorderNodeWalk((*node).right)
		fmt.Println((*node).value)
	}
}
