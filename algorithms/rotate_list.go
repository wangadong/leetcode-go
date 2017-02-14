package algorithms

func rotateRight(head *ListNode, k int) *ListNode {
	cur := head
	if head == nil {
		return head
	}
	len := 1
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	cur.Next = head
	steps := len - k%len
	for i := 0; i < steps; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil
	return head
}
