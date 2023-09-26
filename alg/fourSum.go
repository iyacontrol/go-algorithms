package alg

import "sort"

func FourSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
					left++

				} else if sum > target {
					right--
				} else {
					left++
				}
			}

		}

	}

	return res
}
