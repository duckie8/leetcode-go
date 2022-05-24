/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hs := map[int]int{}
	for i, v := range nums {
		if p, ok := hs[target-v]; ok {
			return []int{p, i}
		}
		hs[v] = i
	}
	return nil
}

// @lc code=end

