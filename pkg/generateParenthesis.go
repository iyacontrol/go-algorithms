package pkg

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例 1：

// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

func GenerateParenthesis(n int) []string {
	res := []string{}

	var dfs func(left, right int, path string)
	dfs = func(left, right int, path string) {
		// 结束条件
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		if left > 0 {
			dfs(left-1, right, path+"(")
		}

		if left < right {
			dfs(left, right-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}
