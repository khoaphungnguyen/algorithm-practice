package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func generateParentheses(n int) []string {
	if n == 0 {
		return []string{}
	}
	var res []string
	var dfs func(path string, open, close int)
	dfs = func(path string, open, close int) {
		if len(path) == n*2 {
			res = append(res, path)
			return
		}
		if open < n {
			dfs(path+"(", open+1, close)
		}
		if close < open {
			dfs(path+")", open, close+1)
		}
	}
	dfs("", 0, 0)
	return res
}

func GenerateParentheses() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	res := generateParentheses(n)
	sort.Strings(res)
	for _, row := range res {
		fmt.Println(row)
	}
}
