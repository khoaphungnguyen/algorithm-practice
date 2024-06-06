package day1dfsontree

import (
	"bufio"
	"fmt"
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
		1. Check if the tree is balanced
		2. Deserialize and serialize a tree
		3. Invert a binary tree
		Press 'd' to exit`)
		// Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Enter tree input for checking balance (e.g., '1 2 4 x 7 x x 5 x x 3 x 6 x x'):")
			BalanceOnTree()
			fmt.Println()
		case "2":
			fmt.Println("Enter tree input for deserializing and serializing (e.g., '1 2 4 x 7 x x 5 x x 3 x 6 x x'):")
			SerDerOnTree()
			fmt.Println()
		case "3":
			fmt.Println("Enter tree input for inverting a tree (e.g., '1 2 4 x x 5 6 x x x 3 x x'):")
			InvertBinaryTree()
			fmt.Println()
		case "d":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			fmt.Println()
		}
	}
}
