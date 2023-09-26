package alg

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		nxt := head.Next
		head.Next = pre

		pre = head
		head = nxt

	}

	return pre
}
