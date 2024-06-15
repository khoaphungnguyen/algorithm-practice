package module3backtracking

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		if val, ok := memo[amount]; ok {
			return val
		}

		minCoins := math.MaxInt32

		for _, coin := range coins {
			res := dfs(amount - coin)
			if res >= 0 && res < minCoins {
				minCoins = res + 1
			}
		}

		// Store the result in the menoization map
		if minCoins == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = minCoins
		}

		return memo[amount]

	}
	return dfs(amount)
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func CoinChange() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	coins := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())
	res := coinChange(coins, amount)
	fmt.Println(res)
}