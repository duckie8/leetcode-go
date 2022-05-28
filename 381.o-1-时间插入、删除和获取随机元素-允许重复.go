/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */

// @lc code=start
type RandomizedCollection struct {
	idx map[int]map[int]struct{}
	nums []int
}


func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx:map[int]map[int]struct{},
	}
}


func (this *RandomizedCollection) Insert(val int) bool {

}


func (this *RandomizedCollection) Remove(val int) bool {

}


func (this *RandomizedCollection) GetRandom() int {

}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

