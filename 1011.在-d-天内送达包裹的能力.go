/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {

	var check func(x int) int
	check = func(x int) int {
		d := 0
		i := 0
		for i < len(weights) {
			cap := x
			for i < len(weights) {
				if cap < weights[i] {
					break
				} else {
					cap -= weights[i]
				}
				i++
			}
			d++
		}
		return d
	}

	l, r := 0, 1
	for _, v := range weights {
		if v > l {
			l = v
		}
		r += v

	}
	for l < r {
		mid := l + (r-l)/2
		if check(mid) <= days {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l

}

// @lc code=end

