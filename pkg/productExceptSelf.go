package pkg

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	// l 表示索引为i左侧数据的乘积
	// r 表示索引为i右侧数据的乘积
	l, r, res := make([]int, n), make([]int, n), make([]int, n)

	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = nums[i-1] * l[i-1]
	}

	r[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = l[i] * r[i]
	}

	return res
}
