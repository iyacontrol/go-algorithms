package alg

import (
	"container/heap"
)

func TopKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	// 堆
	h := &CntHeap{}
	heap.Init(h)

	for num, counter := range cnt {
		heap.Push(h, [2]int{num, counter})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 取值
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-1-i] = heap.Pop(h).([2]int)[0]
	}

	return res
}

type CntHeap [][2]int

func (h CntHeap) Len() int {
	return len(h)
}

func (h CntHeap) Less(i, j int) bool {
	if h[i][1] < h[j][1] {
		return true
	}
	return false
}

func (h CntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *CntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
