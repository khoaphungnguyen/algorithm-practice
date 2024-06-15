package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}
	var res [][]string
	// Memoization table for palindrome checks
	palindromeMemo := make([][]bool, len(s))
	for i := range palindromeMemo {
		palindromeMemo[i] = make([]bool, len(s))
	}
	// Precompute the palindrome table
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || palindromeMemo[i+1][j-1]) {
				palindromeMemo[i][j] = true
			}
		}
	}

	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := index; i < len(s); i++ {

			if palindromeMemo[index][i] {
				dfs(i+1, append(path, s[index:i+1]))
			}
		}
	}
	dfs(0, []string{})
	return res
}

func Partition() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	res := partition(s)
	for _, row := range res {
		fmt.Println(strings.Join(row, " "))
	}
}
