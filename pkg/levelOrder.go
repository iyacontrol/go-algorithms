package pkg

func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	// i表示层号
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})

		nxt := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]

			res[i] = append(res[i], node.Val)

			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}

			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}

		}

		queue = nxt
	}

	return res
}
