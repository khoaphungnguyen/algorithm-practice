package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	day1dfsontree "github.com/khoaphungnguyen/algorithm-practice/day-1-DFSonTree"
	day2bst "github.com/khoaphungnguyen/algorithm-practice/day-2-BST"
	day3backtracking "github.com/khoaphungnguyen/algorithm-practice/day-3-BackTracking"
)

func main() {
	for {
		fmt.Println(`
Please select a topic you want to try:
1. Depth-First Search on Tree
2. Binary Search Tree
3. Backtracking
Press 'd' to exit.`)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		switch input {
		case "1":
			day1dfsontree.Main()
		case "2":
			day2bst.Main()
		case "3":
			day3backtracking.Main()
		case "d":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid input. Please enter 1, 2, or 'd' to exit.")
		}
	}
}
