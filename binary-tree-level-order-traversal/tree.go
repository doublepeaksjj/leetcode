package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	result := [][]int{}
	nextQueue := []*TreeNode{}

	push := func(q *[]*TreeNode, node *TreeNode) {
		*q = append(*q, node)
	}
	pop := func(q *[]*TreeNode) *TreeNode {
		node := (*q)[0]
		*q = (*q)[1:]
		return node
	}

	if root != nil {
		push(&nextQueue, root)
	}
	for len(nextQueue) > 0 {
		result = append(result, []int{})
		queue := nextQueue
		nextQueue = []*TreeNode{}
		for len(queue) > 0 {
			next := pop(&queue)
			level := len(result) - 1
			result[level] = append(result[level], next.Val)
			if next.Left != nil {
				push(&nextQueue, next.Left)
			}
			if next.Right != nil {
				push(&nextQueue, next.Right)
			}
		}
	}
	return result
}
