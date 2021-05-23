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
		扫描两次
	*/
	zeroRowMap := make(map[int]bool)
	zeroColMap := make(map[int]bool)
	for row := range matrix {
		for col, num := range matrix[row] {
			if num == 0 {
				zeroRowMap[row] = true
				zeroColMap[col] = true
			}
		}
	}
	for row := range matrix {
		if zeroRowMap[row] {
			for i := range matrix[row] {
				matrix[row][i] = 0
			}
			continue
		}
		for col := range zeroColMap {
			matrix[row][col] = 0
		}
	}
	return
}

// @lc code=end

