/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	res := traverse(root)
	return res

}

func traverse(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return max(traverse(root.Left)+1, traverse(root.Right)+1)
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// @lc code=end

