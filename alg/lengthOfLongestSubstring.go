package alg

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func LengthOfLongestSubstring(s string) int {
	res := 0
	left, right := 0, 0
	cnt := [128]int{}

	for right < len(s) {
		cnt[s[right]]++
		for cnt[s[right]] > 1 {
			cnt[s[left]]--
			left++
		}

		res = max(res, right-left+1)

		right++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
