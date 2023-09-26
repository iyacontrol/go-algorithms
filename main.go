package main

import (
	"fmt"
	"runtime"
	"time"

	alg "github.com/iyacontrol/go-algorithms/alg"
	"github.com/iyacontrol/go-algorithms/project"
)

type Task struct {
	taskID int
}

func (t *Task) Do() {
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d is being executed\n", t.taskID)
}

func main() {

	runtime.GOMAXPROCS(1)

	// 矩阵置零
	fmt.Println("矩阵置零")
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	alg.SetZeroes(matrix)
	fmt.Println(matrix)

	// 两数之和
	fmt.Println("两数之和")
	nums := []int{2, 7, 11, 15}
	fmt.Println(alg.TwoSum(nums, 9))

	// 搜索旋转有序数组
	fmt.Println("搜索旋转有序数组")
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(alg.Search(nums, 0))

	// 两数相加
	// l1 = [2,4,3], l2 = [5,6,4]
	fmt.Println("两数相加")
	l1 := &alg.ListNode{
		2,
		&alg.ListNode{
			4,
			&alg.ListNode{
				3,
				nil,
			},
		},
	}

	l2 := &alg.ListNode{
		5,
		&alg.ListNode{
			6,
			&alg.ListNode{
				4,
				nil,
			},
		},
	}

	l3 := alg.AddTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

	// 无重复字符的最长子串
	fmt.Println("无重复字符的最长子串")
	s := "abcabcbb"
	fmt.Println(alg.LengthOfLongestSubstring(s))

	fmt.Println("三数之和")
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(alg.ThreeSum(nums))

	fmt.Println("盛最多水的容器")
	nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(alg.MaxArea(nums))

	fmt.Println("最长递增子序列")
	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(alg.LengthOfLIS(nums))

	fmt.Println("打家劫舍")
	nums = []int{1, 2, 3, 1}
	fmt.Println(alg.Rob(nums))

	fmt.Println("删除链表的倒数第 N 个结点")
	l := &alg.ListNode{
		1,
		&alg.ListNode{
			2,
			&alg.ListNode{
				3,
				&alg.ListNode{
					4,
					&alg.ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	l = alg.RemoveNthFromEnd(l, 2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("最大子数组和")
	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(alg.MaxSubArray(nums))

	fmt.Println("反转链表")
	l = &alg.ListNode{
		1,
		&alg.ListNode{
			2,
			&alg.ListNode{
				3,
				&alg.ListNode{
					4,
					&alg.ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	l = alg.ReverseList(l)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("删除有序数组中的重复项")
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(alg.RemoveDuplicates(nums))

	fmt.Println("只出现一次的数字")
	nums = []int{4, 1, 2, 1, 2}
	fmt.Println(alg.SingleNumber(nums))

	fmt.Println("爬楼梯")
	fmt.Println(alg.ClimbStairs(3))

	fmt.Println("买卖股票的最佳时机")
	nums = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(alg.MaxProfit(nums))

	fmt.Println("最长公共前缀")
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(alg.LongestCommonPrefix(strs))

	fmt.Println("移动零")
	nums = []int{0, 1, 0, 3, 12}
	alg.MoveZeroes(nums)
	fmt.Println(nums)

	fmt.Println("合并两个有序链表")
	l1 = &alg.ListNode{
		1,
		&alg.ListNode{
			2,
			&alg.ListNode{
				4,
				nil,
			},
		},
	}

	l2 = &alg.ListNode{
		1,
		&alg.ListNode{
			3,
			&alg.ListNode{
				4,
				nil,
			},
		},
	}

	l = alg.MergeTwoLists(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("和为 K 的子数组")
	nums = []int{1, 2, 3}
	fmt.Println(alg.SubarraySum(nums, 3))

	fmt.Println("搜索插入位置")
	nums = []int{1, 3, 5, 7}
	fmt.Println(alg.SearchInsert(nums, 6))

	fmt.Println("字符串相加")
	num1 := "11"
	num2 := "123"

	fmt.Println(alg.AddStrings(num1, num2))

	fmt.Println("两两交换链表中的节点")
	l1 = &alg.ListNode{
		1,
		&alg.ListNode{
			2,
			&alg.ListNode{
				3,
				&alg.ListNode{
					4,
					nil,
				},
			},
		},
	}

	l = alg.SwapPairs(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("螺旋矩阵")
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(alg.SpiralOrder(matrix))

	fmt.Println("排序链表")

	l1 = &alg.ListNode{
		4,
		&alg.ListNode{
			2,
			&alg.ListNode{
				1,
				&alg.ListNode{
					3,
					nil,
				},
			},
		},
	}

	l = alg.SortList(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("字母异位词分组")
	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(alg.GroupAnagrams(strs))

	fmt.Println("最长连续序列")
	nums = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(alg.LongestConsecutive(nums))

	fmt.Println("找到字符串中所有字母异位词")
	s = "cbaebabacd"
	p := "abc"
	fmt.Println(alg.FindAnagrams(s, p))

	fmt.Println("合并区间")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println(alg.Merge(intervals))

	fmt.Println("快排")
	nums = []int{100, 4, 200, 1, 3, 2}
	alg.QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)

	fmt.Println("归并")
	nums = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(alg.MergeSort(nums))

	fmt.Println("判断子序列")
	s = "ahbgdc"
	p = "abc"
	fmt.Println(alg.IsSubsequence(p, s))

	fmt.Println("除自身以外数组的乘积")
	nums = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(alg.ProductExceptSelf(nums))

	fmt.Println("轮转数组")
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	alg.Rotate(nums, 6)
	fmt.Println(nums)

	fmt.Println("前 K 个高频元素")
	nums = []int{1, 1, 1, 2, 3, 3, 4, 4, 4}
	fmt.Println(alg.TopKFrequent(nums, 2))

	fmt.Println("数组中的第K个最大元素")
	nums = []int{3, 2, 1, 5, 6, 4}
	fmt.Println(alg.FindKthLargest(nums, 2))

	fmt.Println("字符串解码")
	s = "3[a2[c]]"
	fmt.Println(alg.DecodeString(s))

	fmt.Println("每日温度")
	nums = []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(alg.DailyTemperatures(nums))

	fmt.Println("蛇形矩阵")
	for _, m := range alg.SpiderMatrix(9, 5) {
		fmt.Println(m)
	}

	fmt.Println("长度最小的子数组")
	nums = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(alg.MinSubArrayLen(7, nums))

	fmt.Println("旋转图像")
	matrix = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	alg.RotateMatrix(matrix)
	fmt.Println(matrix)

	fmt.Println("组合之和")
	nums = []int{2, 3, 5}
	fmt.Println(alg.CombinationSum(nums, 8))

	fmt.Println("组合之和2")
	nums = []int{2, 3, 5}
	fmt.Println(alg.CombinationSum2(nums, 8))

	fmt.Println("括号生成")
	fmt.Println(alg.GenerateParenthesis(2))

	fmt.Println("全排列")
	nums = []int{1, 2, 3}
	fmt.Println(alg.Permute(nums))

	fmt.Println("子集")
	nums = []int{1, 2, 3}
	fmt.Println(alg.Subsets(nums))

	fmt.Println("下一个排列")
	nums = []int{1, 2, 3}
	alg.NextPermutation(nums)
	fmt.Println(nums)

	fmt.Println("电话号码的字母组合")
	fmt.Println(alg.LetterCombinations("23"))

	fmt.Println("旋转列表")
	l = &alg.ListNode{
		1,
		&alg.ListNode{
			2,
			&alg.ListNode{
				3,
				&alg.ListNode{
					4,
					&alg.ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	l = alg.RotateRight(l, 1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	fmt.Println("从前序与中序遍历序列构造二叉树")
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := alg.BuildTree(preorder, inorder)
	fmt.Println(tree)

	fmt.Println("从后序与中序遍历序列构造二叉树")
	postorder := []int{9, 15, 7, 20, 3}
	inorder = []int{9, 3, 15, 20, 7}

	tree = alg.PostBuildTree(inorder, postorder)
	fmt.Println(tree)

	fmt.Println("最小路径和")
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(alg.MinPathSum(grid))

	fmt.Println("颜色分类")
	nums = []int{2, 0, 2, 1, 1, 0}
	alg.SortColors(nums)
	fmt.Println(nums)

	fmt.Println("N字形变换")
	s = "PAYPALISHIRING"
	fmt.Println(alg.Convert(s, 3))

	fmt.Println("最长回文子串")
	s = "babad"
	fmt.Println(alg.LongestPalindrome(s))

	fmt.Println("路径之和")
	tree = &alg.TreeNode{
		Val: 5,
		Left: &alg.TreeNode{
			Val: 4,
			Left: &alg.TreeNode{
				Val: 11,
				Left: &alg.TreeNode{
					Val: 7,
				},
				Right: &alg.TreeNode{
					Val: 2,
				},
			},
		},
		Right: &alg.TreeNode{
			Val: 8,
			Left: &alg.TreeNode{
				Val: 13,
			},
			Right: &alg.TreeNode{
				Val: 4,
				Left: &alg.TreeNode{
					Val: 5,
				},
				Right: &alg.TreeNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(alg.PathSum(tree, 22))

	fmt.Println("插入区间")
	intervals = [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(alg.Insert(intervals, newInterval))

	fmt.Println("4数之和")
	nums = []int{1, 0, -1, 0, -2, 2}
	fmt.Println(alg.FourSum(nums, 0))

	fmt.Println("字符串相乘")
	fmt.Println(alg.Multiply("1234", "567"))

	return

	fmt.Println("任务管理器")
	numWorkers := 5
	scheduler := project.New(numWorkers)

	for i := 0; i < 200; i++ {
		task := &Task{
			taskID: i,
		}

		scheduler.AddTask(task)

	}

	time.Sleep(250 * time.Second)

	scheduler.Stop()

	fmt.Println("交替打印数字")
	chan1, chan2, chan3 := make(chan int, 1), make(chan int, 1), make(chan int, 1)

	chan1 <- 1
	go func() {
		for i := 1; i <= 10; i++ {
			<-chan1
			fmt.Println(1)
			chan2 <- 2
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			<-chan2
			fmt.Println(2)
			chan3 <- 3
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			<-chan3
			fmt.Println(3)
			if i <= 9 {
				chan1 <- 1
			}
		}
	}()

}
