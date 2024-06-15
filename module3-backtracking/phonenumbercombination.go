package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func letterCombinationsOfPhoneNumber(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var result []string
	keyboard := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi",
		'5': "jkl", '6': "mno", '7': "pqrs",
		'8': "tuv", '9': "wxyz",
	}

	var dfs func(index int, path string)
	dfs = func(index int, path string) {
		if index == len(digits) {
			result = append(result, path)
			return
		}
		letters := keyboard[digits[index]]
		for _, letter := range letters {
			dfs(index+1, path+string(letter))
		}
	}
	dfs(0, "")
	return result
}

func LetterCombinationsOfPhoneNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	digits := scanner.Text()
	res := letterCombinationsOfPhoneNumber(digits)
	fmt.Println(strings.Join(res, " "))
}
