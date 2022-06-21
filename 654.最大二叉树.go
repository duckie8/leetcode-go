/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(nums []int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		index := 0
		maxv := nums[0]
		root := &TreeNode{}
		for i := range nums {
			if nums[i] > maxv {
				maxv = nums[i]
				index = i
			}
		}

		root.Val = maxv
		root.Left = dfs(nums[:index])
		root.Right = dfs(nums[index+1:])
		return root
	}
	return dfs(nums)
}

// @lc code=end

