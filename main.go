package main

import (
	"fmt"

	pkg "github.com/iyacontrol/go-algorithms/pkg"
)

func main() {
	// 矩阵置零
	fmt.Println("矩阵置零")
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	pkg.SetZeroes(matrix)
	fmt.Println(matrix)

	// 两数之和
	fmt.Println("两数之和")
	nums := []int{2, 7, 11, 15}
	fmt.Println(pkg.TwoSum(nums, 9))

	// 搜索旋转有序数组
	fmt.Println("搜索旋转有序数组")
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(pkg.Search(nums, 0))

	// 两数相加
	// l1 = [2,4,3], l2 = [5,6,4]
	fmt.Println("两数相加")
	l1 := &pkg.ListNode{
		2,
		&pkg.ListNode{
			4,
			&pkg.ListNode{
				3,
				nil,
			},
		},
	}

	l2 := &pkg.ListNode{
		5,
		&pkg.ListNode{
			6,
			&pkg.ListNode{
				4,
				nil,
			},
		},
	}

	l3 := pkg.AddTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

	// 无重复字符的最长子串
	fmt.Println("无重复字符的最长子串")
	s := "abcabcbb"
	fmt.Println(pkg.LengthOfLongestSubstring(s))

	fmt.Println("三数之和")
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(pkg.ThreeSum(nums))

	fmt.Println("盛最多水的容器")
	nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(pkg.MaxArea(nums))

	fmt.Println("最长递增子序列")
	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(pkg.LengthOfLIS(nums))

	fmt.Println("打家劫舍")
	nums = []int{1, 2, 3, 1}
	fmt.Println(pkg.Rob(nums))

	fmt.Println("删除链表的倒数第 N 个结点")
	l := &pkg.ListNode{
		1,
		&pkg.ListNode{
			2,
			&pkg.ListNode{
				3,
				&pkg.ListNode{
					4,
					&pkg.ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	l = pkg.RemoveNthFromEnd(l, 2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("最大子数组和")
	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(pkg.MaxSubArray(nums))

	fmt.Println("反转链表")
	l = &pkg.ListNode{
		1,
		&pkg.ListNode{
			2,
			&pkg.ListNode{
				3,
				&pkg.ListNode{
					4,
					&pkg.ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	l = pkg.ReverseList(l)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
