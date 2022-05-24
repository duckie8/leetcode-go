/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < l; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return slow
}

// @lc code=end

