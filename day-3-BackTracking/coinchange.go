package day3backtracking

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

//  You are given coins of different denominations and
//  an amount. Write a function to compute the fewest
//  number of coins that you need to make up that amount.
//  If that amount of money cannot be made up by
//  any combination of the coins, return -1.

// Example 1:

// Input: coins = [1, 2, 5], amount = 11

// Output: 3
// Explanation:

// 11 = 5 + 5 + 1, 3 total coins
// Example 2:

// Input: coins = [3], amount = 1

// Output: -1
