package pkg

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		slow := pre.Next
		fast := pre.Next.Next

		pre.Next = fast
		slow.Next = fast.Next
		fast.Next = slow

		//
		pre = slow
	}

	return dummy.Next
}
