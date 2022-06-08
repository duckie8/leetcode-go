/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {

	l := len(nums)
	ans := make([]int, l)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < 2*l-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%l] {
			ans[stack[len(stack)-1]] = nums[i%l]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%l)
	}
	return ans

}

// @lc code=end

