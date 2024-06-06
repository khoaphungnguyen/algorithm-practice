package day2bst

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func find(tree *Node, val int) bool {
	if tree == nil {
		return false
	}
	if tree.val == val {
		return true
	} else if tree.val < val {
		return find(tree.right, val)
	} else {
		return find(tree.left, val)
	}
}

func insert(tree *Node, val int) *Node {
	if tree == nil {
		return &Node{val: val}
	}
	if tree.val < val {
		tree.right = insert(tree.right, val)
	} else if tree.val > val {
		tree.left = insert(tree.left, val)
	}

	return tree
}

func Main() {
	fmt.Println("BST")
}
