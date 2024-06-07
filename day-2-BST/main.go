package day2bst

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
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

func formatIntBinaryTree(node *Node, out []string) []string {
	if node == nil {
		out = append(out, "x")
	} else {
		out = append(out, strconv.Itoa(node.val))
		out = formatIntBinaryTree(node.left, out)
		out = formatIntBinaryTree(node.right, out)
	}
	return out
}

func Main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(`Please select an option:
		1. Check if the tree is valid BST
		2. Insert a node into a BST correctly
		3. Find the Lowest Common Ancestor of BST
		Press 'd' to exit`)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Enter tree input for checking a BST tree (e.g., '6 4 3 x x 8 x x 8 x x'):")
			ValidBst()
		case "2":
			fmt.Println("Enter tree and val input for inserting a node into a BST tree (e.g., '8 4 2 1 x x 3 x x 6 x x 12 10 x x 14 x 15 x x -- 7'):")
			InsertBst()
		case "3":
			fmt.Println("Enter tree, val1, and val2 input for finding LCA of BST (e.g., '6 2 0 x x 4 3 x x 5 x x 8 7 x x 9 x x--2--4-->2'):")
			LcaOnBST()
		case "d":
			fmt.Println("Exiting the BST program.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			fmt.Println()
		}
	}
}

/*
	Letcode's similar question:
	98. Medium - https://leetcode.com/problems/validate-binary-search-tree/description
	791. Medium - https://leetcode.com/problems/insert-into-a-binary-search-tree/description/
	235.Medium - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
*/
