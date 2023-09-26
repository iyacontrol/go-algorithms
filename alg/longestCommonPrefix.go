package alg

import "strings"

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"

func LongestCommonPrefix(strs []string) string {
	n := len(strs)
	res := strs[0]

	for i := 0; i < n; i++ {
		for strings.Index(strs[i], res) != 0 {
			if len(res) == 0 {
				return ""
			}
			res = res[:len(res)-1]
		}
	}

	return res
}
