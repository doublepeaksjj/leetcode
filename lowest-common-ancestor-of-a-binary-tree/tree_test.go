package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvertTree(t *testing.T) {
	p := &TreeNode{Val: 6}
	q := &TreeNode{Val: 9}
	require.Equal(t, 7, lowestCommonAncestor(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  p,
			Right: q,
		},
	}, p, q).Val)

	p = &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	}
	q = &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 8},
	}
	require.Equal(t, 3, lowestCommonAncestor(&TreeNode{
		Val:   3,
		Left:  p,
		Right: q,
	}, p, q).Val)

	q = &TreeNode{Val: 4}
	p = &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 7},
			Right: q,
		},
	}
	require.Equal(t, 5, lowestCommonAncestor(&TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}, p, q).Val)

	p = &TreeNode{Val: 9}
	q = &TreeNode{Val: 11}
	require.Equal(t, 10, lowestCommonAncestor(&TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val:   10,
				Left:  p,
				Right: q,
			},
		},
	}, p, q).Val)
}
