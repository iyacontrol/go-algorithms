package alg

import (
	"strconv"
	"strings"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// 示例 2：

// 输入：s = "3[a2[c]]"
// 输出："accaccacc"

func DecodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	str := ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			numStack = append(numStack, num)
			num = 0
			strStack = append(strStack, str)
			str = ""
		} else if char == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			content := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			str = content + strings.Repeat(str, count)

		} else {
			str = str + string(char)
		}
	}

	return str
}
