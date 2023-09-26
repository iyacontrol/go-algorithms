package alg

func PathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	var dfs func(node *TreeNode, sum int, temp []int)
	dfs = func(node *TreeNode, sum int, temp []int) {

		if node == nil {
			if sum == targetSum {
				newTemp := make([]int, len(temp))
				copy(newTemp, temp)
				res = append(res, newTemp)
			}
			return
		}

		if node.Left == nil && node.Right == nil {
			temp = append(temp, node.Val)
			dfs(nil, sum+node.Val, temp)
			temp = temp[:len(temp)-1]
		}

		if node.Left != nil {
			temp = append(temp, node.Val)
			dfs(node.Left, sum+node.Val, temp)
			temp = temp[:len(temp)-1]
		}

		if node.Right != nil {
			temp = append(temp, node.Val)
			dfs(node.Right, sum+node.Val, temp)
			temp = temp[:len(temp)-1]
		}

	}

	dfs(root, 0, []int{})

	return res
}
