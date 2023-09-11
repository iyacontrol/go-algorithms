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

	fmt.Println("删除有序数组中的重复项")
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(pkg.RemoveDuplicates(nums))

	fmt.Println("只出现一次的数字")
	nums = []int{4, 1, 2, 1, 2}
	fmt.Println(pkg.SingleNumber(nums))

	fmt.Println("爬楼梯")
	fmt.Println(pkg.ClimbStairs(3))

	fmt.Println("买卖股票的最佳时机")
	nums = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(pkg.MaxProfit(nums))

	fmt.Println("最长公共前缀")
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(pkg.LongestCommonPrefix(strs))

	fmt.Println("移动零")
	nums = []int{0, 1, 0, 3, 12}
	pkg.MoveZeroes(nums)
	fmt.Println(nums)

	fmt.Println("合并两个有序链表")
	l1 = &pkg.ListNode{
		1,
		&pkg.ListNode{
			2,
			&pkg.ListNode{
				4,
				nil,
			},
		},
	}

	l2 = &pkg.ListNode{
		1,
		&pkg.ListNode{
			3,
			&pkg.ListNode{
				4,
				nil,
			},
		},
	}

	l = pkg.MergeTwoLists(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("和为 K 的子数组")
	nums = []int{1, 2, 3}
	fmt.Println(pkg.SubarraySum(nums, 3))

	fmt.Println("搜索插入位置")
	nums = []int{1, 3, 5, 7}
	fmt.Println(pkg.SearchInsert(nums, 6))

	fmt.Println("字符串相加")
	num1 := "11"
	num2 := "123"

	fmt.Println(pkg.AddStrings(num1, num2))

	fmt.Println("两两交换链表中的节点")
	l1 = &pkg.ListNode{
		1,
		&pkg.ListNode{
			2,
			&pkg.ListNode{
				3,
				&pkg.ListNode{
					4,
					nil,
				},
			},
		},
	}

	l = pkg.SwapPairs(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("螺旋矩阵")
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(pkg.SpiralOrder(matrix))

	fmt.Println("排序链表")

	l1 = &pkg.ListNode{
		4,
		&pkg.ListNode{
			2,
			&pkg.ListNode{
				1,
				&pkg.ListNode{
					3,
					nil,
				},
			},
		},
	}

	l = pkg.SortList(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("字母异位词分组")
	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(pkg.GroupAnagrams(strs))

	fmt.Println("最长连续序列")
	nums = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(pkg.LongestConsecutive(nums))

	fmt.Println("找到字符串中所有字母异位词")
	s = "cbaebabacd"
	p := "abc"
	fmt.Println(pkg.FindAnagrams(s, p))

	fmt.Println("合并区间")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println(pkg.Merge(intervals))

	fmt.Println("快排")
	nums = []int{100, 4, 200, 1, 3, 2}
	pkg.QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)

	fmt.Println("归并")
	nums = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(pkg.MergeSort(nums))

}
