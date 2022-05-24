/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {

	windows, need := map[byte]int{}, map[byte]int{}

	for i, _ := range t {
		need[t[i]]++
	}

	left, right := 0, 0

	valid := 0

	start, length := 0, math.MaxInt32

	for right < len(s) {

		c := s[right]
		right++

		if _, ok := need[c]; ok {
			windows[c]++
			if need[c] == windows[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if need[d] == windows[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]

}

// @lc code=end

