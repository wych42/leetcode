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
		单独的空间用来记录哪些行、列要设置成0
		可以利用 matrix 的第一行和第一列来记录这个数据
		再单独判断第一行和第一列
	*/
	rows := len(matrix)
	cols := len(matrix[0])

	firstRowZero := false
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				// 如果 row,col 是0，对应的第一行的这一列也是0
				matrix[0][col] = 0
				if row > 0 { // 跳过第一行，需要保留第一行第一列的数字，否则会影响到后面判断第一列是否为0
					matrix[row][0] = 0
				} else { // 记录第一行是不是 0
					firstRowZero = true
				}
			}
		}
	}
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			// 对于每一个cell，如果它对应的第一行是0或者第一列是0，那么当前cell也是0
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
	}
	// 如果第一个元素是0，那么第一列都是0
	if matrix[0][0] == 0 {
		for row := 0; row < rows; row++ {
			matrix[row][0] = 0
		}
	}
	if firstRowZero {
		for col := 0; col < cols; col++ {
			matrix[0][col] = 0
		}
	}
	return
}

// @lc code=end

