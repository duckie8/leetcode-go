/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		preSum: make([]int, len(nums)+1),
	}
	na.preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		na.preSum[i+1] = na.preSum[i] + nums[i]
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

