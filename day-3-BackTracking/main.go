package day3backtracking

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val      int
	children []Node
}

func buildTree(nodes []string, pos *int) Node {
	val, _ := strconv.Atoi(nodes[*pos])
	*pos++
	num, _ := strconv.Atoi(nodes[*pos])
	*pos++
	children := []Node{}
	for i := 0; i < num; i++ {
		children = append(children, buildTree(nodes, pos))
	}
	return Node{val, children}
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func Main() {
	reader := bufio.NewReader(os.Stdin)

	{
		fmt.Println(`Please select an option:
		1. Find all the root-to-leaf paths of a ternary tree 
		Press 'd' to exit`)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			fmt.Println("Enter tree input for finnd all paths of ternary tree (e.g., '1 3 2 1 5 0 3 0 4 0'):")
			TernaryTreePaths()
		default:
			fmt.Println("Invalid option. Please try again.")
			fmt.Println()
		}
	}

}
