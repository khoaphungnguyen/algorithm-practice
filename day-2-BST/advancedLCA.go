package day2bst

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lca(root *Node, node1 *Node, node2 *Node) *Node {
	if root == nil || root == node1 || root == node2 {
		return root
	}
	left := lca(root.left, node1, node2)
	right := lca(root.right, node1, node2)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func findNode(root *Node, target int) *Node {
	if root == nil {
		return nil
	}
	if root.val == target {
		return root
	}
	leftSearch := findNode(root.left, target)
	if leftSearch != nil {
		return leftSearch
	}
	return findNode(root.right, target)
}

func AdvancedLca() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	treeIdx := 0
	root := buildTree(splitWords(scanner.Text()), &treeIdx)
	scanner.Scan()
	target1, _ := strconv.Atoi(scanner.Text())
	node1 := findNode(root, target1)
	scanner.Scan()
	target2, _ := strconv.Atoi(scanner.Text())
	node2 := findNode(root, target2)
	res := lca(root, node1, node2)
	if res == nil {
		fmt.Println("null")
	} else {
		fmt.Println(res.val)
	}

}
