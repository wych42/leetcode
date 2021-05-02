/*
 * @lc app=leetcode id=554 lang=golang
 *
 * [554] Brick Wall
 *
 * https://leetcode.com/problems/brick-wall/description/
 *
 * algorithms
 * Medium (50.79%)
 * Likes:    1463
 * Dislikes: 72
 * Total Accepted:    87.1K
 * Total Submissions: 168.7K
 * Testcase Example:  '[[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]'
 *
 * There is a rectangular brick wall in front of you with n rows of bricks. The
 * i^th row has some number of bricks each of the same height (i.e., one unit)
 * but they can be of different widths. The total width of each row is the
 * same.
 *
 * Draw a vertical line from the top to the bottom and cross the least bricks.
 * If your line goes through the edge of a brick, then the brick is not
 * considered as crossed. You cannot draw a line just along one of the two
 * vertical edges of the wall, in which case the line will obviously cross no
 * bricks.
 *
 * Given the 2D array wall that contains the information about the wall, return
 * the minimum number of crossed bricks after drawing such a vertical line.
 *
 *
 * Example 1:
 *
 *
 * Input: wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: wall = [[1],[1],[1]]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * n == wall.length
 * 1 <= n <= 10^4
 * 1 <= wall[i].length <= 10^4
 * 1 <= sum(wall[i].length) <= 2 * 10^4
 * sum(wall[i]) is the same for each row i.
 * 1 <= wall[i][j] <= 2^31 - 1
 *
 *
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	/*
		Input: wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
		Output: 2
		记录每一行的缝隙位置：
		0,1,-,3,-,5,6
		0,-,-,3,4,-,6
		0,1,-,-,4,-,6
		0,-,2,-,-,-,6
		0,-,-,3,4,-,6
		0,1,-,-,4,5,6
		0和6不算，剩下的 - 最少的就是结果
	*/
	gaps := make(map[int]int) // 记录每一列有缝隙的行数
	for _, row := range wall {
		idx := 0
		for i := 0; i < len(row)-1; i++ {
			idx += row[i]
			gaps[idx]++
		}
	}
	rowCount := len(wall)
	ans := rowCount
	for _, v := range gaps {
		if diff := rowCount - v; diff < ans {
			ans = diff
		}
	}
	return ans
}

// @lc code=end

