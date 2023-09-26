package alg

import (
	"container/heap"
)

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4

func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	h := &IntArryHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		heap.Push(h, nums[i])
	}

	for i := 0; i < n-k; i++ {
		heap.Pop(h)
		// fmt.Println(heap.Pop(h).(int))
	}

	return heap.Pop(h).(int)
}

type IntArryHeap []int

func (h IntArryHeap) Len() int {
	return len(h)
}

func (h IntArryHeap) Less(i, j int) bool {
	if h[i] < h[j] {
		return true
	}
	return false
}

func (h IntArryHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntArryHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntArryHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
