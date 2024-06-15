package module2binarysearchtree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lcaOnBst(root *Node, p int, q int) int {
	if root.val > p && root.val > q {
		return lcaOnBst(root.left, p, q)
	} else if root.val < p && root.val < q {
		return lcaOnBst(root.right, q, p)
	}
	return root.val
}

// Find the Lowest Common Ancestor of a Binary Search Tree
func LcaOnBST() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bstIdx := 0
	bst := buildTree(splitWords(scanner.Text()), &bstIdx)
	scanner.Scan()
	p, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	res := lcaOnBst(bst, p, q)
	fmt.Println(res)
}
