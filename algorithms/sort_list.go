package algorithms

import "sort"

// https://leetcode.com/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	intSlice := sort.IntSlice{head.Val}

	for head.Next != nil {
		head = head.Next
		intSlice = append(intSlice, head.Val)
	}
	intSlice.Sort()
	head = &ListNode{}
	next := head
	if len(intSlice) > 0 {
		next.Val = intSlice[0]
	} else {
		return head
	}
	for _, val := range intSlice[1:] {
		next.Next = &ListNode{}
		next = next.Next
		next.Val = val
	}
	return head
}
