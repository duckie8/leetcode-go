/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := postorder[len(postorder)-1]
	var nodeVal int
	for k, v := range inorder {
		if v == root {
			nodeVal = k
			break
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(inorder[:nodeVal], postorder[0:nodeVal]),
		Right: buildTree(inorder[nodeVal+1:], postorder[nodeVal:len(postorder)-1]),
	}
}

// @lc code=end

