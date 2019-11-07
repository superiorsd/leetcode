package main

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root == nil {
		return res
	}

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = append(stack[:len(stack)-1])

		if node.Right != nil {
			node = node.Right
			for node != nil {
				stack = append(stack, node)
				node = node.Left
			}
		}
	}

	return res
}
