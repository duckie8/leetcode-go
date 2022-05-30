/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int, len(nums2))
	var st []int
	for _, num := range nums2 {
		for len(st) > 0 && num > st[len(st)-1] {
			//单调栈的下一个大的数就是当前数的下一个大值
			mp[st[len(st)-1]] = num
			st = st[:len(st)-1]
		}
		st = append(st, num)
	}

	res := make([]int, 0, len(nums1))

	for _, n := range nums1 {
		if v, ok := mp[n]; !ok {
			res = append(res, -1)
		} else {
			res = append(res, v)
		}
	}
	return res
}

// @lc code=end

