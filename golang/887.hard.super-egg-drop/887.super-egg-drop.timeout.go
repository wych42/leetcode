/*
 * @lc app=leetcode id=887 lang=golang
 *
 * [887] Super Egg Drop
 *
 * https://leetcode.com/problems/super-egg-drop/description/
 *
 * algorithms
 * Hard (27.02%)
 * Likes:    1322
 * Dislikes: 96
 * Total Accepted:    28.6K
 * Total Submissions: 106K
 * Testcase Example:  '1\n2'
 *
 * You are given k identical eggs and you have access to a building with n
 * floors labeled from 1 to n.
 *
 * You know that there exists a floor f where 0 <= f <= n such that any egg
 * dropped at a floor higher than f will break, and any egg dropped at or below
 * floor f will not break.
 *
 * Each move, you may take an unbroken egg and drop it from any floor x (where
 * 1 <= x <= n). If the egg breaks, you can no longer use it. However, if the
 * egg does not break, you may reuse it in future moves.
 *
 * Return the minimum number of moves that you need to determine with certainty
 * what the value of f is.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 1, n = 2
 * Output: 2
 * Explanation:
 * Drop the egg from floor 1. If it breaks, we know that f = 0.
 * Otherwise, drop the egg from floor 2. If it breaks, we know that f = 1.
 * If it does not break, then we know f = 2.
 * Hence, we need at minimum 2 moves to determine with certainty what the value
 * of f is.
 *
 *
 * Example 2:
 *
 *
 * Input: k = 2, n = 6
 * Output: 3
 *
 *
 * Example 3:
 *
 *
 * Input: k = 3, n = 14
 * Output: 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= 100
 * 1 <= n <= 10^4
 *
 *
 */

// @lc code=start
/*
这个解法是会超时的：o(kn^2)
输入 2,6, dp table
[0 1 2 3 4 5 6]
[0 1 2 2 3 3 3]
*/
const maxInt = 1 << 31

func superEggDrop(k int, n int) int {
	/*
		如果从 x 层开始测试，问题就成了:
		- x 层没破 求解(k, n-x) + 1
		- x 层破了 求解(k-1, x-1) + 1
		这两个数里的最大值就是从 x 层开始测试需要的操作次数

		对 1<=x<=n，计算每一个的最大值，列表的最小值，就是答案
	*/
	dp := make([][]int, k+1) // 包含 0 个鸡蛋情况
	dp[0] = make([]int, n+1) // 0 个鸡蛋测所有的都是 2^31(不可能做到)
	for i := 1; i <= n; i++ {
		dp[0][i] = maxInt
	}
	for i := 1; i <= k; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0 // 第0层需要0次
		for j := 1; j <= n; j++ {
			// 计算第 j 层操作需要的次数
			min := maxInt
			// 假设从 x 层开始仍
			for x := 1; x <= j; x++ {
				//         在x层没碎的;  在x层碎了,蛋减1，楼层减1
				val := max(dp[i][j-x], dp[i-1][x-1])
				//	fmt.Println("input", i, j, "on floor", x, val)
				if val < min {
					min = val
				}
			}
			//fmt.Println("input", i, j, "val", min)
			dp[i][j] = min + 1
		}
		//for _, row := range dp {
		//	fmt.Println(row)
		//}
	}
	return dp[k][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

