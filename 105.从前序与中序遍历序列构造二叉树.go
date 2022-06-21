/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root int
	for k, v := range inorder {
		if v == preorder[0] {
			root = k
			break
		}

	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}

// @lc code=end

