/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (49.01%)
 * Likes:    8164
 * Dislikes: 241
 * Total Accepted:    1M
 * Total Submissions: 2.1M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given an m x n 2D binary grid grid which represents a map of '1's (land) and
 * '0's (water), return the number of islands.
 *
 * An island is surrounded by water and is formed by connecting adjacent lands
 * horizontally or vertically. You may assume all four edges of the grid are
 * all surrounded by water.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 300
 * grid[i][j] is '0' or '1'.
 *
 *
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	/*
		选中一个格子；如果格子==1（是个岛屿），就开始出发找四周
		标记是岛屿的格子，直到结束

		再选一个是岛屿的格子，跳过已经访问过的，标记
	*/
	rows := len(grid)
	cols := len(grid[0])
	ans := 0

	var dfs func(row, col int)
	dfs = func(row, col int) {
		// constrain
		if row >= rows || col >= cols || row < 0 || col < 0 {
			return
		}
		if grid[row][col] != '1' {
			return
		}
		grid[row][col] = 2
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
		return
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				dfs(row, col)
				ans++
			}
		}
	}
	return ans
}

// @lc code=end

