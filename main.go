package main

import (
	"fmt"

	algo "github.com/wangadong/leetcode-go/algorithms"
)

func main() {
	// fmt.Println(algorithms.IsInterleave("ab", "bc", "abbc"))
	nums := []int{9, 5, 2, 3}
	obj := algo.Constructor(nums)
	param_1 := obj.Reset()
	param_2 := obj.Shuffle()
	fmt.Println(param_1, param_2)
}
