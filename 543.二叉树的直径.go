/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil {
		return maxDia
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lHeight := dfs(node.Left)
		rHeight := dfs(node.Right)
		// maxDia即为该二叉树的最大直径
		maxDia = max(maxDia, lHeight+rHeight)
		// 该节点为根的二叉树的最大高度，也就是1+max(lh+rh)
		return 1 + max(lHeight, rHeight)
	}
	dfs(root)
	return maxDia

}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// @lc code=end

