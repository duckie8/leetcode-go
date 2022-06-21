/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = preorder[0]
	if len(postorder) == 1 {
		return root
	}

	mid := 0
	for i, v := range postorder {
		if v == preorder[1] {
			mid = i
			break
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  constructFromPrePost(preorder[1:mid+2], postorder[:mid+1]),
		Right: constructFromPrePost(preorder[mid+2:], postorder[mid+1:len(postorder)-1]),
	}
}

// @lc code=end

