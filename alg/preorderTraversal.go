package alg

func PreorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return res
}
