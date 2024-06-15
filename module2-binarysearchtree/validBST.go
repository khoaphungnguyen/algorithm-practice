package module2binarysearchtree

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func validBst(root *Node) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *Node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.val <= min || root.val >= max {
		return false
	}
	return validate(root.left, min, root.val) && validate(root.right, root.val, max)
}
func ValidBst() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rootIdx := 0
	root := buildTree(splitWords(scanner.Text()), &rootIdx)
	res := validBst(root)
	fmt.Println(res)
}
