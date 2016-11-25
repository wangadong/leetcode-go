package main

import "fmt"
import "github.com/wangadong/leetcode-go/algorithms"

func main() {
	// // Interleaving String
	// fmt.Println(algorithms.IsInterleave("ab", "bc", "abbc"))

	// // Shuffle Array
	// nums := []int{9, 5, 2, 3}
	// obj := algo.Constructor(nums)
	// param_1 := obj.Reset()
	// param_2 := obj.Shuffle()
	// fmt.Println(param_1, param_2)

	// // Two Sum
	// fmt.Println(algo.TwoSum([]int{3, 2, 4}, 6))

	// // Two Sum Sorted
	// fmt.Println(algo.TwoSum([]int{2, 3, 4}, 6))

	// Length of Longest Substring
	// fmt.Println(algorithms.LengthOfLongestSubstring("dvdf"))

	// Number of Islands
	fmt.Println(algorithms.NumIslands([][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}}))
}
