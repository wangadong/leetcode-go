package algorithms

// https://leetcode.com/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	startEven := even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = startEven
	return head
}
