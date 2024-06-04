package day1dfsontree

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func isBalanced(tree *Node) bool {
	return TreeHeightDiff(tree) != 1
}

func TreeHeightDiff(tree *Node) int {
	if tree == nil {
		return 0
	}
	left := TreeHeightDiff(tree.left)
	right := TreeHeightDiff(tree.right)
	if left == -1 || right == -1 || math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}
	return int(math.Max(float64(left), float64(right))) + 1
}

func buildTree(nodes []string, pos *int) *Node {
	val := nodes[*pos]
	*pos++
	if val == "x" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	left := buildTree(nodes, pos)
	right := buildTree(nodes, pos)
	return &Node{v, left, right}
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func BalanceOnTree() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	treeIdx := 0
	tree := buildTree(splitWords(scanner.Text()), &treeIdx)
	res := isBalanced(tree)
	fmt.Println(res)
}
