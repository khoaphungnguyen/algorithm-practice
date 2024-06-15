package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func decodeWays(s string) int {
	if len(s) == 0 {
		return 0
	}

	memo := make(map[int]int)

	var dfs func(index int) int
	dfs = func(index int) int {
		if index == len(s) {
			return 1
		}

		if s[index] == '0' {
			return 0
		}

		// Check memoization map
		if val, ok := memo[index]; ok {
			return val
		}

		// Decode single character
		result := dfs(index + 1)

		// Decode two characters
		if index+1 < len(s) {
			twoDigit, _ := strconv.Atoi(s[index : index+2])
			if twoDigit >= 10 && twoDigit <= 26 {
				result += dfs(index + 2)
			}
		}

		// Store result in memoization map
		memo[index] = result
		return result
	}
	return dfs(0)
}

func DecodeWays() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	digits := scanner.Text()
	res := decodeWays(digits)
	fmt.Println(res)
}
