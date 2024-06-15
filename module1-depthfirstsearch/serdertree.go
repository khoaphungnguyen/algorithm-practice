package module1depthfirstsearch

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func serialize(root *Node) string {
	var res []string
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			res = append(res, "x")
			return
		}
		res = append(res, strconv.Itoa(root.val))
		dfs(root.left)
		dfs(root.right)
	}
	dfs(root)
	return strings.Join(res, " ")
}

func deserialize(root string) *Node {
	nodes := strings.Split(root, " ")
	index := 0
	var dfs func() *Node
	dfs = func() *Node {
		if index >= len(nodes) {
			return nil
		}
		val := nodes[index]
		index++
		if val == "x" {
			return nil
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Error parsing string to int:", err)
			return nil
		}
		newNode := &Node{val: num}
		newNode.left = dfs()
		newNode.right = dfs()
		return newNode
	}
	return dfs()
}

func SerDerOnTree() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bstIdx := 0
	root := buildTree(splitWords(scanner.Text()), &bstIdx)
	res := deserialize(serialize(root))
	out := make([]string, 0)
	out = formatIntBinaryTree(res, out)
	fmt.Println(strings.Join(out, " "))
}
