package pkg

func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	//

	beforHalfEnd := endOfBeforeHalf(head)
	afterHalfEnd := ReverseList(beforHalfEnd.Next)

	p1, p2 := head, afterHalfEnd

	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

func endOfBeforeHalf(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
