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
	// top-down/dfs 求解
	// 把所有可能的路径算一遍
	// 超时: [1,2,5] 100
	// amount - (coin for coin in coins) 看结果是不是 0

	// 用于缓存，amount = X 的时候，需要多少个，一个需要 amount+1 个元素存储
	minCoins := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		minCoins[i] = amount + 1
	}
	minCoins[0] = 0

	var dfs func(int) int
	dfs = func(a int) int {
		if a < 0 {
			return a
		}

		for _, c := range coins {
			// 每一个 coin, 都走 dfs 算一遍
			// 比如 [1, 2, 5] 11; c=1; 需要计算 dfs(10); 就需要分别计算 dfs(9), dfs(8), dfs(5)
			res := dfs(a - c) // res + 1 就是这个路径下需要的最小数量
			if res >= 0 {
				minCoins[a] = min(minCoins[a], res+1)
			}
		}
		if minCoins[a] < amount+1 {
			return minCoins[a]
		} else {
			return -1
		}
	}
	return dfs(amount)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

