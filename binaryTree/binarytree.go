package binarytree

type node[T comparable] struct {
	Value T
	Left  *node[T]
	Right *node[T]
}

type BinaryTree[T comparable] *node[T]

func CreateRoot[T comparable](item T) *BinaryTree[T] {
	node := &node[T]{
		Value: item,
		Left:  nil,
		Right: nil,
	}

	bt := BinaryTree[T](node)

	return &bt
}
