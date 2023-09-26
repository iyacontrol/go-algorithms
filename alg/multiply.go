package alg

import "strconv"

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)

	if n1 < n2 {
		return Multiply(num2, num1)
	}

	res := ""

	for p2 := n2 - 1; p2 >= 0; p2-- {
		v2 := int(num2[p2] - '0')
		carry := 0
		curr := ""

		for p1 := n1 - 1; p1 >= 0; p1-- {
			v1 := int(num1[p1] - '0')
			temp := v2*v1 + carry
			curr = strconv.Itoa(temp%10) + curr
			carry = temp / 10
		}

		if carry > 0 {
			curr = strconv.Itoa(carry) + curr
		}

		for i := 0; i < n2-1-p2; i++ {
			curr += "0"
		}

		res = AddStrings(res, curr)

	}

	return res
}
