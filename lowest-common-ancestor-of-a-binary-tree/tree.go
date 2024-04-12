package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PathTreeNode struct {
	Node *TreeNode
	Path []*TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	stack := []*PathTreeNode{}
	push := func(n *TreeNode, p []*TreeNode) {
		stack = append(stack, &PathTreeNode{Node: n, Path: p})
	}
	pop := func() *PathTreeNode {
		ret := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return ret
	}
	lca := &TreeNode{}
	if root != nil {
		push(root, []*TreeNode{root})
		lca = root
	}
	var pPath []*TreeNode
	var qPath []*TreeNode
	for len(stack) > 0 {
		next := pop()
		// Search the tree until we know the paths for both p and q
		if next.Node == p {
			pPath = next.Path
		} else if next.Node == q {
			qPath = next.Path
		}
		if pPath != nil && qPath != nil {
			break
		}
		if next.Node.Left != nil {
			push(next.Node.Left, newPath(next.Path, next.Node.Left))
		}
		if next.Node.Right != nil {
			push(next.Node.Right, newPath(next.Path, next.Node.Right))
		}
	}

	if pPath == nil || qPath == nil {
		return root
	}

	// LCA is the point at which pPath and qPath diverge
	for i := 0; i < len(pPath) && i < len(qPath) && pPath[i] == qPath[i]; i++ {
		lca = pPath[i]
	}
	return lca
}

func newPath(p []*TreeNode, n *TreeNode) []*TreeNode {
	newPath := make([]*TreeNode, 0, len(p)+1)
	newPath = append(newPath, p...)
	return append(newPath, n)
}
