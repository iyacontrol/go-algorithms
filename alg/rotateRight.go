package alg

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

func RotateRight(head *ListNode, k int) *ListNode {
	// 优先返回特殊情况
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	// 计算长度
	n := 1
	p := head
	for p.Next != nil {
		p = p.Next
		n++
	}

	add := n - k%n

	if add == n {
		return head
	}

	p.Next = head
	for add > 0 {
		p = p.Next
		add--
	}

	res := p.Next
	p.Next = nil

	return res

}
