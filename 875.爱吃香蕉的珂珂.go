/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 10000000000 + 1
	var sumHour func(piles []int, t int) int
	sumHour = func(piles []int, x int) int {
		hour := 0
		for _, v := range piles {
			hour += v / x
			if v%x > 0 {
				hour++
			}
		}
		return hour
	}

	for left < right {
		mid := left + (right-left)/2
		if sumHour(piles, mid) == h {
			right = mid
		}
		if sumHour(piles, mid) < h {
			right = mid
		}
		if sumHour(piles, mid) > h {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

