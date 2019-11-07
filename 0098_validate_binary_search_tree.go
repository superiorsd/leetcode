package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = append(stack[:len(stack)-1])
		if pre != nil && pre.Val >= root.Val {
			return false
		}

		pre = root
		root = root.Right
	}

	return true
}
