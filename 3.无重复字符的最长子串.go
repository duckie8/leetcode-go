/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	windows := map[byte]int{}
	left, right := 0, 0
	res := 0

	for right < len(s) {
		c := s[right]
		right++
		windows[c]++
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

