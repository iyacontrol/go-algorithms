package pkg

// 给你一个字符串 s，找到 s 中最长的回文子串。

// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

// 示例 1：

// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：

// 输入：s = "cbbd"
// 输出："bb"
func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	res := string(s[0])
	maxLength := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if s[i] != s[j] {
				dp[j][i] = false
				continue
			} else {
				if i-j < 3 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}

			if dp[j][i] && i-j+1 > maxLength {
				maxLength = i - j + 1
				res = s[j : i+1]
			}

		}
	}

	return res

}
