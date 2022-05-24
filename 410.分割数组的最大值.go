/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	var check func(x int) int
	check = func(x int) int {
		d := 0
		i := 0
		for i < len(nums) {
			cap := x
			for i < len(nums) {
				if cap < nums[i] {
					break
				} else {
					cap -= nums[i]
				}
				i++
			}
			d++
		}
		return d
	}

	l, r := 0, 1
	for _, v := range nums {
		if v > l {
			l = v
		}
		r += v

	}
	for l < r {
		mid := l + (r-l)/2
		if check(mid) <= m {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end

