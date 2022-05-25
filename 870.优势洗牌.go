/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 */

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)

	type Node struct {
		idx   int
		value int
	}

	var nodes = make([]Node, 0, len(nums2))
	for i, v := range nums2 {
		nodes = append(nodes, Node{i, v})
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].value < nodes[j].value
	})
	var res = make([]int, len(nums1))

	idx_start, idx_end := 0, len(res)-1
	for i := 0; i < len(nums1); i++ {
		if nums1[i] <= nodes[idx_start].value {
			res[nodes[idx_end].idx] = nums1[i]
			idx_end--
		} else {
			res[nodes[idx_start].idx] = nums1[i]
			idx_start++
		}
	}
	return res

}

// @lc code=end

