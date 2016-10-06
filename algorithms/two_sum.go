package algorithms

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if indice1, ok := numMap[target-num]; ok {
			return []int{indice1, i}
		}
		numMap[num] = i
	}
	return []int{}
}
