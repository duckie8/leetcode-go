/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 */

// @lc code=start
type Solution struct {
	nums int
	idx  map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	mapping := make(map[int]int)
	bsize := n - len(blacklist)
	for _, v := range blacklist {
		mapping[v] = 1
	}

	last := n - 1
	for _, v := range blacklist {
		if v >= bsize {
			continue
		}

		for {
			_, ok := mapping[last]
			if ok {
				last--
			} else {
				break
			}
		}
		mapping[v] = last
		last--
	}

	return Solution{nums: bsize, idx: mapping}

}

func (this *Solution) Pick() int {
	//因为可以返回的数目是非白名单的数目
	index := rand.Intn(this.nums)
	//如果在前面的index是黑名单，就转到映射的
	if v, ok := this.idx[index]; ok {
		return v
	}

	return index

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end

