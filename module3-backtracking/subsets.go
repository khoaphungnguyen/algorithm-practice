package module3backtracking

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)
func subsets(nums []int) [][]int {
	var result [][]int
	var currentSubset []int

	var backtrack func(index int)
	backtrack = func(index int) {
		// Add a copy of currrentSubset
		temp := make([]int, len(currentSubset))
		copy(temp, currentSubset)
		result = append(result, temp)

		for i := index; i < len(nums); i++ {
			currentSubset = append(currentSubset, nums[i])
			backtrack(i + 1)
			currentSubset = currentSubset[:len(currentSubset)-1]
		}
	}
	backtrack(0)
	return result
}

func Sunset() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := arrayAtoi(splitWords(scanner.Text()))
	res := subsets(nums)
	var strSubsets []string
	for _, subset := range res {
		sort.Ints(subset)
		var subsetStr = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(subset)), " "), "[]")
		strSubsets = append(strSubsets, subsetStr)
	}
	sort.Strings(strSubsets)
	for _, row := range strSubsets {
		fmt.Println(row)
	}
}
