package module2binarysearchtree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertBst(root *Node, val int) *Node {
	if root == nil {
		return &Node{val: val}
	}
	if val < root.val {
		root.left = insertBst(root.left, val)
	} else {
		root.right = insertBst(root.right, val)
	}
	return root
}

func InsertBst() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rootIdx := 0
	tree := buildTree(splitWords(scanner.Text()), &rootIdx)
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	res := insertBst(tree, val)
	out := make([]string, 0)
	out = formatIntBinaryTree(res, out)
	fmt.Println(strings.Join(out, " "))
}
