package pkg

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}

	pre := dummy

	v1, v2, carry := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
		} else {
			v1 = 0
		}

		if l2 != nil {
			v2 = l2.Val
		} else {
			v2 = 0
		}

		sum := v1 + v2 + carry

		carry = sum / 10

		pre.Next = &ListNode{
			Val: sum % 10,
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		pre = pre.Next

	}

	if carry > 0 {
		pre.Next = &ListNode{
			Val: 1,
		}
	}

	return dummy.Next
}
