// https://leetcode.com/problems/remove-element
func removeElement(nums []int, val int) int {
    nextIndex := 0
    for _, num := range nums {
        if num != val {
            nums[nextIndex] = num
            nextIndex++
            continue
        }
    }
    return nextIndex
}
