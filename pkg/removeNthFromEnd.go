package pkg

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		-1,
		head,
	}

	slow, fast := dummy, dummy

	for i := 0; i < n; i++ {
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
