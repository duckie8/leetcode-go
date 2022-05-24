/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 */

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {

	nums := make([]int, n)
	for _, v := range bookings {
		nums[v[0]-1] += v[2]
		if v[1] < n {
			nums[v[1]] -= v[2]
		}
	}

	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums

}

// @lc code=end

