package module1depthfirstsearch

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func invertBinaryTree(tree *Node) *Node {
	if tree == nil {
		return nil
	}
	tree.left, tree.right = tree.right, tree.left
	// Recursively invert the left and right subtrees
	invertBinaryTree(tree.left)
	invertBinaryTree(tree.right)
	return tree
}

func InvertBinaryTree() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	treeIdx := 0
	tree := buildTree(splitWords(scanner.Text()), &treeIdx)
	res := invertBinaryTree(tree)
	out := make([]string, 0)
	out = formatIntBinaryTree(res, out)
	fmt.Println(strings.Join(out, " "))
}
