package pkg

func DailyTemperatures(temperatures []int) []int {
	// 暴力解法
	// n := len(temperatures)
	// res := make([]int, n)

	// for i := 0; i < n; i++ {
	// J:
	// 	for j := i + 1; j < n; j++ {
	// 		if temperatures[j] > temperatures[i] {
	// 			res[i] = j - i
	// 			break J
	// 		}
	// 	}
	// }

	// return res

	// 栈
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIndex] = i - preIndex
		}

		stack = append(stack, i)
	}

	return res
}
