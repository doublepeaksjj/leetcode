package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	push := func(n *TreeNode) {
		stack = append(stack, n)
	}
	pop := func() *TreeNode {
		ret := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return ret
	}
	if root != nil {
		push(root)
	}
	for len(stack) > 0 {
		next := pop()
		next.Left, next.Right = next.Right, next.Left
		if next.Left != nil {
			push(next.Left)
		}
		if next.Right != nil {
			push(next.Right)
		}
	}
	return root
}
