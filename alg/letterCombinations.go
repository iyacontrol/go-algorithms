package alg

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

// 示例 1：

// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

var dig2alpha = map[byte]string{
	'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
	'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
}

func LetterCombinations(digits string) []string {
	res := []string{}
	n := len(digits)
	if n < 1 {
		return res
	}

	var dfs func(index int, temp []byte)
	dfs = func(index int, temp []byte) {
		if len(temp) == n {
			newTemp := make([]byte, len(temp))
			copy(newTemp, temp)
			res = append(res, string(newTemp))

			return
		}

		now := dig2alpha[digits[index]]

		for i := 0; i < len(now); i++ {
			temp = append(temp, now[i])
			dfs(index+1, temp)
			temp = temp[:len(temp)-1]
		}

	}

	dfs(0, []byte{})

	return res
}
