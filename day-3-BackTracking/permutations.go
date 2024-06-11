package day3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func permutations(letters string) []string {
	if len(letters) == 0 {
		return []string{}
	}
	var res []string
	visited := make([]bool, len(letters))

	var dfs func(path []byte)
	dfs = func(path []byte) {
		if len(path) == len(letters) {
			res = append(res, string(path))
			return
		}
		for i := 0; i < len(letters); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && letters[i] == letters[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			dfs(append(path, letters[i]))
			visited[i] = false
		}
	}

	dfs([]byte{})
	return res
}

func Permutations() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	letters := scanner.Text()
	res := permutations(letters)
	sort.Strings(res)
	for _, row := range res {
		fmt.Println(row)
	}
}
