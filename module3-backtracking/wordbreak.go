package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
)

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	memo := make(map[int]bool)
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(s) {
			return true
		}
		if val, exists := memo[index]; exists {
			return val
		}

		for i := index + 1; i <= len(s); i++ {
			if wordSet[s[index:i]] && dfs(i) {
				memo[index] = true
				return true
			}
		}

		memo[index] = false
		return false
	}
	return dfs(0)
}

func WorkBreak() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	words := splitWords(scanner.Text())
	res := wordBreak(s, words)
	fmt.Println(res)
}
