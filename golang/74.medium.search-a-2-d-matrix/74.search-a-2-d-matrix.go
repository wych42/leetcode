/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (37.69%)
 * Likes:    3120
 * Dislikes: 194
 * Total Accepted:    436.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 有 n = rows x cols 个数；对这些数做二分查找
	// 此时位置为 i(=(n-1)/2), 需要定位 i 的 row,col
	// 每行有 cols 个数; i 所在的行数即row=i/cols, 列数为col= i-row*cols
	// 0,1,2,3,4
	// 5,6,7,8,9
	// rows = 2;cols=5; i=8; row=8/5=1; col=8-1*5=3
	// 搜索范围变成了 [0:row*col] 或者 [row*col:n]
	cols := len(matrix[0])
	rows := len(matrix)
	start := 0
	end := cols*rows - 1

	getVal := func(idx int) int {
		row := idx / cols
		col := idx - row*cols
		return matrix[row][col]
	}
	// 二分查找
	for start+1 < end {
		mid := start + (end-start)/2
		midValue := getVal(mid)
		//fmt.Println(start, end, mid, midValue)
		if midValue == target {
			return true
		}
		if midValue < target {
			start = mid
			continue
		}
		if midValue > target {
			end = mid
			continue
		}
	}

	// start, end 这两个还没有检查
	if getVal(start) == target || getVal(end) == target {
		return true
	}
	return false
}

// @lc code=end

