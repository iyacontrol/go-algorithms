package alg

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := slow.Next
	slow.Next = nil

	return merge(SortList(head), SortList(tail))

}

func merge(head, tail *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}

	pre := dummy

	for head != nil && tail != nil {
		if head.Val < tail.Val {
			pre.Next = head
			head = head.Next
		} else {
			pre.Next = tail
			tail = tail.Next
		}

		pre = pre.Next
	}

	if head != nil {
		pre.Next = head
	}

	if tail != nil {
		pre.Next = tail
	}

	return dummy.Next

}
