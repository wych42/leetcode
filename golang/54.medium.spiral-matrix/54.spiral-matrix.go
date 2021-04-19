/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (35.93%)
 * Likes:    3717
 * Dislikes: 660
 * Total Accepted:    488.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	ans := make([]int, 0, rows*cols)
	left := 0
	right := cols
	top := 0
	bottom := rows
	for left < right && top < bottom {
		for i := left; i < right; i++ { // top row
			ans = append(ans, matrix[top][i])
		}
		top++
		for i := top; i < bottom; i++ { // right column
			ans = append(ans, matrix[i][right-1])
		}
		right--

		// important
		if !(left < right && top < bottom) {
			break
		}

		for i := right - 1; i >= left; i-- { // bottom row
			ans = append(ans, matrix[bottom-1][i])
		}
		bottom--
		for i := bottom - 1; i >= top; i-- { // left column
			ans = append(ans, matrix[i][left])
		}
		left++
	}
	return ans
}

// @lc code=end

