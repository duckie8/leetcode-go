/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {

	windows, need := map[byte]int{}, map[byte]int{}

	for i, _ := range s1 {
		need[s1[i]]++
	}

	left, right := 0, 0

	valid := 0

	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if len(need) == valid {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}
	}
	return false
}

// @lc code=end

