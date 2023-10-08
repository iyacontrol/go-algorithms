package alg

import "container/heap"

func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	if k > n1*n2 {
		k = n1 * n2
	}

	h := &PairsH{}
	heap.Init(h)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			heap.Push(h, [2]int{nums1[i], nums2[j]})

			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}

	res := [][]int{}

	for i := 0; i < k; i++ {
		x := heap.Pop(h).([2]int)
		res = append(res, []int{x[0], x[1]})
	}

	return res
}

type PairsH [][2]int

func (h PairsH) Len() int {
	return len(h)
}

func (h PairsH) Less(i, j int) bool {
	if h[i][0]+h[i][1] > h[j][0]+h[j][1] {
		return true
	}

	return false
}

func (h PairsH) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairsH) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *PairsH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x

}
