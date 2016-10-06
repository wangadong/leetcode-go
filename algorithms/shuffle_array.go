package algorithms

import (
	"math/rand"
	"time"
)

//Use Fisher & Yates Permutation Random Algo
type Solution struct {
	Hash   map[int]int
	length int
	r      *rand.Rand
}

func Constructor(nums []int) Solution {
	solution := Solution{
		Hash: make(map[int]int),
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for i, v := range nums {
		solution.Hash[i] = v
		solution.length++
	}
	return solution
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	original := make([]int, this.length)
	for i := 0; i < this.length; i++ {
		original[i] = this.Hash[i]
	}
	return original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	shuffled := make([]int, this.length)

	for i, v := range this.r.Perm(this.length) {
		shuffled[i] = this.Hash[v]
	}
	return shuffled
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
