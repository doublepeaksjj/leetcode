package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreeLevelOrderTraversal(t *testing.T) {
	require.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}))
	require.Equal(t, [][]int{{3}, {9, 20}, {6, 2, 15, 7}, {1, 4}}, levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  9,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:  20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{
				Val:   7,
				Right: &TreeNode{Val: 4},
			},
		},
	}))
	require.Equal(t, [][]int{{1}}, levelOrder(&TreeNode{Val: 1}))
	require.Equal(t, [][]int{}, levelOrder(nil))
}
