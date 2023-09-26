package alg

// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = BuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
