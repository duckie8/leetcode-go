/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {

	l := len(matrix)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		i, j := 0, len(row)-1
		for i < j {
			// swap(row[i], row[j]);
			row[i], row[j] = row[j], row[i]
			i++
			j--
		}
	}

}

// @lc code=end

