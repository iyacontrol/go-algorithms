package pkg

func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	mid := n / 2

	return mergeArray(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func mergeArray(left, right []int) []int {
	n1, n2 := len(left), len(right)
	res := make([]int, 0)

	p, q := 0, 0

	for p < n1 && q < n2 {
		if left[p] < right[q] {
			res = append(res, left[p])
			p++
		} else {
			res = append(res, right[q])
			q++
		}
	}

	if p < n1 {
		res = append(res, left[p:]...)
	}

	if q < n2 {
		res = append(res, right[q:]...)
	}

	return res
}
