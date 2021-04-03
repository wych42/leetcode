/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 *
 * https://leetcode.com/problems/coin-change-2/description/
 *
 * algorithms
 * Medium (51.71%)
 * Likes:    2967
 * Dislikes: 76
 * Total Accepted:    184.4K
 * Total Submissions: 354.6K
 * Testcase Example:  '5\n[1,2,5]'
 *
 * You are given coins of different denominations and a total amount of money.
 * Write a function to compute the number of combinations that make up that
 * amount. You may assume that you have infinite number of each kind of
 * coin.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: amount = 5, coins = [1, 2, 5]
 * Output: 4
 * Explanation: there are four ways to make up the amount:
 * 5=5
 * 5=2+2+1
 * 5=2+1+1+1
 * 5=1+1+1+1+1
 *
 *
 * Example 2:
 *
 *
 * Input: amount = 3, coins = [2]
 * Output: 0
 * Explanation: the amount of 3 cannot be made up just with coins of 2.
 *
 *
 * Example 3:
 *
 *
 * Input: amount = 10, coins = [10]
 * Output: 1
 *
 *
 *
 *
 * Note:
 *
 * You can assume that
 *
 *
 * 0 <= amount <= 5000
 * 1 <= coin <= 5000
 * the number of coins is less than 500
 * the answer is guaranteed to fit into signed 32-bit integer
 *
 *
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1) // 从0开始 [0] [0,1] [0,1,2], [0,1,2,5]
	dp[0] = make([]int, amount+1)     // 从0开始
	dp[0][0] = 1                      // 只有 0 个 coin 时，有一种方法合成 0 amount, 合成其他 amount 的方法都是 0

	for i, coin := range coins {
		row := make([]int, amount+1)
		row[0] = 1 // 凑成0个肯定有一种方法

		// 用当前 coin 凑 j
		for j := 1; j < amount+1; j++ {
			if coin > j {
				// coin 比 j 大，单用 coin 本身是凑不起来的, 用不上 coin
				// 那么能凑到 j 的方法数，只能是 coins[:i] 的数量
				row[j] += dp[i][j]
				continue
			}
			// 否则，除了 coins[:i] 能凑齐来的 j 的数量(不用 coin)
			// 还多一个就是，用了 coin 能凑起来的数量，这个数等于 用 coin 凑 j-coin 的数量
			row[j] += dp[i][j] + row[j-coin]
		}
		dp[i+1] = row
	}
	return dp[len(coins)][amount]
}

// @lc code=end

