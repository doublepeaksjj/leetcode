package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvertTree(t *testing.T) {
	require.Equal(t, "4, 2, 1, 3, 7, 6, 9", sprintTree(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}))
	require.Equal(t, "4, 7, 9, 6, 2, 3, 1", sprintTree(invertTree(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	})))
}

func sprintTree(root *TreeNode) string {
	str := ""
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
		if str == "" {
			str += fmt.Sprint(next.Val)
		} else {
			str += fmt.Sprint(", ", next.Val)
		}
		if next.Right != nil {
			push(next.Right)
		}
		if next.Left != nil {
			push(next.Left)
		}
	}
	return str
}
