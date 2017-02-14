package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-nodes-in-k-group
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	prev := &ListNode{}

	count := 0
	for count != k {
		if head != nil {
			head = head.Next
			count++
		} else {
			return dummyHead.Next
		}
	}

	if head != nil {
		prev = reverseKGroup(head, k)
	} else {
		prev = nil
	}

	curHead := dummyHead.Next
	for i := k; i > 0; i-- {
		tmp := curHead.Next
		curHead.Next = prev
		prev = curHead
		curHead = tmp
	}

	return prev
}
