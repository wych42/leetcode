/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 *
 * https://leetcode.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (44.65%)
 * Likes:    3494
 * Dislikes: 370
 * Total Accepted:    432.2K
 * Total Submissions: 963.6K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given an m x n matrix. If an element is 0, set its entire row and column to
 * 0. Do it in-place.
 *
 * Follow up:
 *
 *
 * A straight forward solution using O(mn) space is probably a bad idea.
 * A simple improvement uses O(m + n) space, but still not the best
 * solution.
 * Could you devise a constant space solution?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * Output: [[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	/*
		不用新开辟空间，实时更新当前行的值，并且记住哪些列该设置为0
	*/
	zeroColMap := make(map[int]bool)
	for row := range matrix {
		isZeroRow := false
		for col, num := range matrix[row] {
			if num == 0 {
				zeroColMap[col] = true
				isZeroRow = true
				// 把当前列，在row之前的行都设置为0
				for r := 0; r < row; r++ {
					matrix[r][col] = 0
				}
			} else {
				if zeroColMap[col] {
					matrix[row][col] = 0
				} else {
					matrix[row][col] = num
				}
			}
		}
		if isZeroRow {
			for i := range matrix[row] {
				matrix[row][i] = 0
			}
		}
	}
	return
}

// @lc code=end

