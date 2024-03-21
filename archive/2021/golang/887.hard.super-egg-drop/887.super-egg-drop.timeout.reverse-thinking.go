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
func superEggDrop(k int, n int) int {
	/*
		1. 步骤不会超过 n：一个鸡蛋挨个仍，最多n次也能测试出来
		2. 问题变成求解 k 个鸡蛋，m 步，能测出多少楼层；m <=n; 层数 >=n 的时候就找到答案
		3. 基础case:
			如果只有一步：不管多少个鸡蛋，都只能测出来一层
			如果只有一个蛋：有多少步，能测出多少层
		4. 在求解 f(k,m) 时，用掉一步,分两种情况:
		    鸡蛋碎了: 步骤，鸡蛋都减1， 那么还可以测 f(k-1, m-1) 层(向下)
			鸡蛋没碎: 步骤减1，鸡蛋不变, 还可以测 f(k, m-1) 层(向上)
			+1 测的那层
			f(k,m) = f(k-1,m-1) + f(k,m-1) + 1
		5. 如果 f(k,m) >= n，m 即答案
	*/
	if n == 1 {
		return 1
	}
	if k == 1 {
		return n
	}
	dp := make([][]int, n+1) // 从步骤数 0 开始
	dp[0] = make([]int, k+1) // 0 步只能测0层
	dp[1] = make([]int, k+1)
	for i := 1; i <= k; i++ { // 1步只能测1层
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i // 1 个鸡蛋 i 步能测 i 层
		for j := 2; j <= k; j++ {
			dp[i][j] = 1 + dp[i-1][j] + dp[i-1][j-1]
			if dp[i][j] >= n {
				return i
			}
		}
	}
	return 0
}

// @lc code=end

