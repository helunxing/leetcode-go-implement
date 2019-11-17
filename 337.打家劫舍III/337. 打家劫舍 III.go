package leetcode

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)

	dfs = func(root *TreeNode) (a, b int) {
		if root == nil {
			return 0, 0
		}
		la, lb := dfs(root.Left)
		ra, rb := dfs(root.Right)
		a = lb + rb + root.Val
		b = max(la, lb) + max(ra, rb)
		return
	}

	return max(dfs(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
