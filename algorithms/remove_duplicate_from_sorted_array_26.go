// https://leetcode.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
    var lastNum, lastIdx int
    for i, num := range nums {
        if i == 0 {
            lastNum = num
            lastIdx = 1
            continue
        }
        if num != lastNum {
            lastNum = num
            nums[lastIdx] = num
            lastIdx++
        }
    }
    return lastIdx
}

func removeDuplicates1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    i := 0
    for j := 1; j< len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }
    return i+1
}
