package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type NULLPOINTER struct {
	Err     error
	message string
}

func (root *Node) traverse_tree() (int, error) {
	if root == nil {
		return -1, fmt.Errorf("nil node")
	}
	maxVal := root.val
	if root.left != nil {
		if left, err := root.left.traverse_tree(); err == nil {
			maxVal = max(maxVal, left)
		}

	}

	if root.right != nil {
		if right, err := root.right.traverse_tree(); err == nil {
			maxVal = max(maxVal, right)
		}
	}

	return maxVal, nil
}

func createTree() *Node {
	root := Node{
		val: 10,
	}
	root.left = &Node{val: 10,
		left: &Node{val: 2,
			right: nil, left: nil},
		right: &Node{val: 11,
			left: nil, right: nil}}

	root.right = &Node{val: 2,
		left: &Node{val: 5,
			right: nil, left: nil},
		right: &Node{val: 20,
			left: nil, right: nil}}

	return &root
}

func main() {
	root := createTree()
	maxVal, err := root.traverse_tree()
	if err == nil {
		fmt.Printf("max value is %3d\n", maxVal)
	}
}
