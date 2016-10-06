package algorithms

func TwoSumII(nums []int, target int) []int {
	length := len(nums)
	index1, index2 := 0, length-1
	for nums[index1]+nums[index2] != target {
		if (nums[index1] + nums[index2]) > target {
			index2--
		} else {
			index1++
		}
	}
	return []int{index1, index2}
}
