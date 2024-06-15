package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	module1depthfirstsearch "github.com/khoaphungnguyen/algorithm-practice/module1-depthfirstsearch"
	module2binarysearchtree "github.com/khoaphungnguyen/algorithm-practice/module2-binarysearchtree"
	module3backtracking "github.com/khoaphungnguyen/algorithm-practice/module3-backtracking"
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
			module1depthfirstsearch.Main()
		case "2":
			module2binarysearchtree.Main()
		case "3":
			module3backtracking.Main()
		case "d":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid input. Please enter 1, 2, or 'd' to exit.")
		}
	}
}
