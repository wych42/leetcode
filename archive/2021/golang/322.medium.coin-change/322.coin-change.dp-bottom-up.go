/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (37.16%)
 * Likes:    6515
 * Dislikes: 190
 * Total Accepted:    611.8K
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given an integer array coins representing coins of different
 * denominations and an integer amount representing a total amount of money.
 *
 * Return the fewest number of coins that you need to make up that amount. If
 * that amount of money cannot be made up by any combination of the coins,
 * return -1.
 *
 * You may assume that you have an infinite number of each kind of coin.
 *
 *
 * Example 1:
 *
 *
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Example 3:
 *
 *
 * Input: coins = [1], amount = 0
 * Output: 0
 *
 *
 * Example 4:
 *
 *
 * Input: coins = [1], amount = 1
 * Output: 1
 *
 *
 * Example 5:
 *
 *
 * Input: coins = [1], amount = 2
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// bottom-up
	// dp[0], dp[1] 挨个算到 dp[amount]

	// 初始化 dp, 一共 amount+1 个元素, 从 0 开始计算
	// 每个初始值是 amount+1/math.MaxInt
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0 // 需要 0 个才能凑成 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			// 如果 i - c < 0; 说明用 c 凑不成 i
			if i-c >= 0 {
				// 取 1+dp[i-c] 和 dp[i] 的最小值
				dp[i] = min(dp[i], 1+dp[i-c])
			}
		}
	}
	if dp[amount] != amount+1 { // 更新过了
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

