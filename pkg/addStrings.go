package pkg

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

// 示例 1：

// 输入：num1 = "11", num2 = "123"
// 输出："134"

func AddStrings(num1 string, num2 string) string {
	res := ""

	l1, l2 := len(num1), len(num2)
	p, q := l1-1, l2-1

	v1, v2, carry := 0, 0, 0
	for p >= 0 || q >= 0 {
		if p >= 0 {
			v1 = int(num1[p] - '0')
			p--
		} else {
			v1 = 0
		}

		if q >= 0 {
			v2 = int(num2[q] - '0')
			q--
		} else {
			v2 = 0
		}

		sum := v1 + v2 + carry

		carry = sum / 10

		res = string(sum%10+'0') + res

	}

	if carry > 0 {
		res = string(1+'0') + res
	}

	return res
}
