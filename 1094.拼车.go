/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	for _, v := range trips {
		nums[v[1]] += v[0]
		nums[v[2]] -= v[0]

	}

	var preSum int = 0
	for i := 0; i < 1001; i++ {
		preSum += nums[i]
		if preSum > capacity {
			return false
		}
	}
	return true
}

// @lc code=end

