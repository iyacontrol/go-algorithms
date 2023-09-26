package alg

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

func IsSubsequence(s string, t string) bool {
	lenS, lenT := len(s), len(t)

	ps, pt := 0, 0

	for pt < lenT && ps < lenS {
		if s[ps] == t[pt] {
			ps++
		}

		pt++
	}

	return ps == lenS

}
