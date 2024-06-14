package day3backtracking

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val      int
	children []Node
}

func buildTree(nodes []string, pos *int) Node {
	val, _ := strconv.Atoi(nodes[*pos])
	*pos++
	num, _ := strconv.Atoi(nodes[*pos])
	*pos++
	children := []Node{}
	for i := 0; i < num; i++ {
		children = append(children, buildTree(nodes, pos))
	}
	return Node{val, children}
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func Main() {
	reader := bufio.NewReader(os.Stdin)

	{
		fmt.Println(`Please select an option:
		1. Find all the root-to-leaf paths of a ternary tree 
		2. Find a list of strings in lexicographical order.
		3. Find all possible letter combinations the phone
		4. Partition A String Into Palindromes
		5. Generate All Valid Parentheses
		6. General All Permutations
		7. Work Break
		8. Number of Ways to Decode a Message
		9. Minimum Number of Coins to Make Up a Given Value
		Press 'd' to exit`)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			fmt.Println("Enter tree input for find all paths of ternary tree (e.g., '1 3 2 1 5 0 3 0 4 0'):")
			TernaryTreePaths()
		case "2":
			fmt.Println("Enter a number for fiding a list of strings in lexicographical order(e.g., '2'):")
			LetterCombination()
		case "3":
			fmt.Println("Enter phone number that contains digits 2-9 to find all possible letter combinations the phone number(e.g., '23'):")
			LetterCombinationsOfPhoneNumber()
		case "4":
			fmt.Println("Enter a string s for partitioning s such that every substring of the partition is a palindrome(e.g., 'aba'):")
			Partition()
		case "5":
			fmt.Println("Enter an integer n for generating all strings with n matching parentheses(e.g., '2'):")
			GenerateParentheses()
		case "6":
			fmt.Println("Enter a string of unique letters for finding all of its distinct permutations.(e.g., 'abc'):")
			Permutations()
		case "7":
			fmt.Println("Enter a string and a list of words for determining if the string can be constructed from concatenating words from the list of words.(e.g., 'workbreak' 'word break'):")
			WorkBreak()
		case "8":
			fmt.Println("Enter a non-empty string of digits for showing how many ways are there to decode it(e.g., '18' --> '2' ways):")
			DecodeWays()
		case "9":
			fmt.Println("Enter coins of different denominations and an amount for returning the fewest number of coins that you need to make up that amount(e.g., '[1, 2, 5] 11' --> '3' coins):")
			CoinChange()
		default:
			fmt.Println("Invalid option. Please try again.")
			fmt.Println()
		}
	}
}

/*
	Letcode's similar question:
	17. Medium - https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
	131. Medium - https://leetcode.com/problems/palindrome-partitioning/description/
	22. Medium - https://leetcode.com/problems/generate-parentheses/description/
	46. Medium - https://leetcode.com/problems/permutations/description/
	139. Medium - https://leetcode.com/problems/word-break/description/
	91. Medium - https://leetcode.com/problems/decode-ways/description/
	322. Medium - https://leetcode.com/problems/coin-change/description/

*/
