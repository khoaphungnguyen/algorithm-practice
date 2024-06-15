package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ternaryTreePaths(root Node) []string {
	var paths []string
	if len(root.children) == 0 {
		return []string{strconv.Itoa(root.val)}
	}
	for _, child := range root.children {
		subPaths := ternaryTreePaths(child)
		for _, subPath := range subPaths {
			paths = append(paths, strconv.Itoa(root.val)+"->"+subPath)
		}
	}
	return paths
}

func TernaryTreePaths() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rootIdx := 0
	root := buildTree(splitWords(scanner.Text()), &rootIdx)
	res := ternaryTreePaths(root)
	for _, row := range res {
		fmt.Println(row)
	}
}
