/*
 * @lc app=leetcode.cn id=1706 lang=golang
 *
 * [1706] 球会落何处
 *
 * https://leetcode-cn.com/problems/where-will-the-ball-fall/description/
 *
 * algorithms
 * Medium (68.82%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    25.6K
 * Total Submissions: 37.1K
 * Testcase Example:  '[[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]'
 *
 * 用一个大小为 m x n 的二维网格 grid 表示一个箱子。你有 n 颗球。箱子的顶部和底部都是开着的。
 *
 * 箱子中的每个单元格都有一个对角线挡板，跨过单元格的两个角，可以将球导向左侧或者右侧。
 *
 *
 * 将球导向右侧的挡板跨过左上角和右下角，在网格中用 1 表示。
 * 将球导向左侧的挡板跨过右上角和左下角，在网格中用 -1 表示。
 *
 *
 * 在箱子每一列的顶端各放一颗球。每颗球都可能卡在箱子里或从底部掉出来。如果球恰好卡在两块挡板之间的 "V"
 * 形图案，或者被一块挡导向到箱子的任意一侧边上，就会卡住。
 *
 * 返回一个大小为 n 的数组 answer ，其中 answer[i] 是球放在顶部的第 i 列后从底部掉出来的那一列对应的下标，如果球卡在盒子里，则返回
 * -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid =
 * [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]
 * 输出：[1,-1,-1,-1,-1]
 * 解释：示例如图：
 * b0 球开始放在第 0 列上，最终从箱子底部第 1 列掉出。
 * b1 球开始放在第 1 列上，会卡在第 2、3 列和第 1 行之间的 "V" 形里。
 * b2 球开始放在第 2 列上，会卡在第 2、3 列和第 0 行之间的 "V" 形里。
 * b3 球开始放在第 3 列上，会卡在第 2、3 列和第 0 行之间的 "V" 形里。
 * b4 球开始放在第 4 列上，会卡在第 2、3 列和第 1 行之间的 "V" 形里。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[-1]]
 * 输出：[-1]
 * 解释：球被卡在箱子左侧边上。
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid =
 * [[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]]
 * 输出：[0,1,2,3,4,-1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * grid[i][j] 为 1 或 -1
 *
 *
 */

// @lc code=start
func findBall(grid [][]int) []int {
	// 题目约束 height, width >= 1, 不做判断
	// height := len(grid)
	width := len(grid[0])

	result := make([]int, width)
	// 初始位置
	for i := 0; i < width; i++ {
		result[i] = i
	}
	// 处理第几行
	for _, row := range grid {
		// 当前第几个球在哪个位置上
		for idx, pos := range result {
			// 如果确定不能往下走了，就跳过
			if pos == -1 {
				continue
			}

			// 小球所在位置的数值
			value := row[pos]

			if value == 1 { // 向右导球, 看右边的值
				if pos == width-1 || row[pos+1] == -1 {
					result[idx] = -1
					continue
				}
				result[idx] += 1
			}
			if value == -1 { // 向左导球，看左边的值
				if pos == 0 || row[pos-1] == 1 {
					result[idx] = -1
					continue
				}
				result[idx] -= 1

			}

		}
	}
	return result
}

// @lc code=end

