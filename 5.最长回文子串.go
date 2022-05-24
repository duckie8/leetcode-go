/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	var res = ""
	for i := 0; i < len(s); i++ {
		s1 := centerPalindrome(s, i, i)

		s2 := centerPalindrome(s, i, i+1)

		if len(res) > len(s1) {
			res = res

		} else {
			res = s1
		}

		if len(res) > len(s2) {
			res = res

		} else {
			res = s2
		}

	}
	return res

}

func centerPalindrome(s string, i int, j int) string {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return s[i+1 : j]
}

// @lc code=end

