package day3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func letterCombination(n int) []string {
	var result []string
	var dfs func(n int, res *[]string, startIndex int, path *[]string)
	dfs = func(n int, res *[]string, startIndex int, path *[]string) {
		if startIndex == n {
			*res = append(*res, strings.Join(*path, ""))
			return
		}
		for _, letter := range []string{"a", "b"} {
			*path = append(*path, letter)
			dfs(n, res, startIndex+1, path)
			*path = (*path)[:len(*path)-1]
		}
	}
	dfs(n, &result, 0, &[]string{})
	return result
}

func LetterCombination() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	res := letterCombination(n)
	sort.Strings(res)
	for _, row := range res {
		fmt.Println(row)
	}
}
