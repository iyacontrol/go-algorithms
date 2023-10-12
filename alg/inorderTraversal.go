package alg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	// res := []int{}

	// inorder(root, &res)

	// return res

	res := []int{}

	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)

		root = root.Right
	}

	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	// 中序
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)

	// 前序
	// *res = append(*res, node.Val)
	// inorder(node.Left, res)
	// inorder(node.Right, res)

	// 后序
	inorder(node.Left, res)
	inorder(node.Right, res)
	*res = append(*res, node.Val)
}
