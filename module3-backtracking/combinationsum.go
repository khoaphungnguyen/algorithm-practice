package module3backtracking

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)


func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var currentCombination []int

	sort.Ints(candidates)

	var backtrack func(target int, start int)
	backtrack = func(target int, start int) {
		//Base case if target is zero
		if target == 0 {
			// make a copy of currentCombination and add it to the result
			combination := make([]int, len(currentCombination))
			copy(combination, currentCombination)
			res = append(res, combination)
			return
		}

		// Iterate over the candidates start with the start index
		for i := start; i < len(candidates); i++ {
			// if the current candidates exceeds the target, skip it
			if candidates[i] > target {
				break
			}
			// Include the current candidate in the combination
			currentCombination = append(currentCombination, candidates[i])
			// Recursive explore further with the recuded target
			backtrack(target-candidates[i], i)
			// Backtrack: remove the last included candidate to explore other combinations
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
	backtrack(target, 0)
	return res
}

func arrayItoa(arr []int) []string {
	res := []string{}
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

func CombinationSum() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	candidates := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	res := combinationSum(candidates, target)
	list_lt := func(a, b []int) bool {
		for i := 0; i < int(math.Min(float64(len(a)), float64(len(b)))); i++ {
			if a[i] < b[i] {
				return true
			} else if a[i] > b[i] {
				return false
			}
		}
		return len(a) < len(b)
	}
	for _, lst := range res {
		sort.Ints(lst)
	}
	sort.Slice(res, func(i, j int) bool { return list_lt(res[i], res[j]) })
	for _, lst := range res {
		fmt.Println(strings.Join(arrayItoa(lst), " "))
	}
}
