/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {

	windows, need := map[byte]int{}, map[byte]int{}

	for i, _ := range p {
		need[p[i]]++
	}

	left, right := 0, 0

	valid := 0

	indexs := []int{}

	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			windows[c]++
			if need[c] == windows[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				indexs = append(indexs, left)
			}

			d := s[left]
			left++

			if _, ok := windows[d]; ok {
				if need[d] == windows[d] {
					valid--
				}
				windows[d]--
			}
		}
	}
	return indexs
}

// @lc code=end

