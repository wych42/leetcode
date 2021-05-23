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
		一行一行扫描，新开数组记录当前行的结果；
		发现零时记录列值，并把当前行设置为0
	*/
	zeroColMap := make(map[int]bool)
	for row := range matrix {
		resRow := make([]int, len(matrix[row]))
		isZeroRow := false
		for col, num := range matrix[row] {
			if num == 0 {
				zeroColMap[col] = true
				isZeroRow = true
				for r := 0; r < row; r++ {
					matrix[r][col] = 0
				}
			} else {
				if zeroColMap[col] {
					resRow[col] = 0
				} else {
					resRow[col] = num
				}
			}
		}
		if isZeroRow {
			matrix[row] = make([]int, len(matrix[row]))
		} else {
			matrix[row] = resRow
		}
	}
	return
}

// @lc code=end

