/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 *
 * https://leetcode.com/problems/rotate-image/description/
 *
 * algorithms
 * Medium (59.83%)
 * Likes:    4692
 * Dislikes: 332
 * Total Accepted:    560.4K
 * Total Submissions: 923.8K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * You are given an n x n 2D matrix representing an image, rotate the image by
 * 90 degrees (clockwise).
 *
 * You have to rotate the image in-place, which means you have to modify the
 * input 2D matrix directly. DO NOT allocate another 2D matrix and do the
 * rotation.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[7,4,1],[8,5,2],[9,6,3]]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 *
 *
 * Example 3:
 *
 *
 * Input: matrix = [[1]]
 * Output: [[1]]
 *
 *
 * Example 4:
 *
 *
 * Input: matrix = [[1,2],[3,4]]
 * Output: [[3,1],[4,2]]
 *
 *
 *
 * Constraints:
 *
 *
 * matrix.length == n
 * matrix[i].length == n
 * 1 <= n <= 20
 * -1000 <= matrix[i][j] <= 1000
 *
 *
 */

// @lc code=start
func rotate(matrix [][]int) {
	/*
		123    741
		456 -> 852
		789    963
		观察发现是这样一个个替换的;替换完一圈之后，旋转内部的 matrix
		1-3-9-7-1
		2-6-8-4-2
		...
		idx
		1->3 matrix[top][left+idx]->matrix[top+idx][right]
		3->9 matrix[top+idx][right]->matrix[bottom][right-idx]
		9->7 matrix[bottom][right-idx]->matrix[bottom-idx][left]
		7->1 matrix[bottom-idx][left]->matrix[top][left+idx]
	*/
	top, left := 0, 0
	right, bottom := len(matrix)-1, len(matrix)-1
	for left < right {
		for idx := 0; idx < right-left; idx++ { // 只走到 倒数第二个
			topLeft := matrix[top][left+idx]
			// bottom-left to top-left
			matrix[top][left+idx] = matrix[bottom-idx][left]
			// bottom-right to bottom-left
			matrix[bottom-idx][left] = matrix[bottom][right-idx]
			// top-right to bottom-right
			matrix[bottom][right-idx] = matrix[top+idx][right]
			// top-left to top-right
			matrix[top+idx][right] = topLeft
		}

		left++
		right--
		top++
		bottom--
	}
	return
}

// @lc code=end