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
	dp := make([]int, amount+1) // 从0开始
	dp[0] = 1                   // 只有 0 个 coin 时，有一种方法合成 0 amount

	// 子问题：用 coins[0:i] 凑 dp 的组合数
	// [1,2,5] 5
	// dp     0,1,2,3,4,5
	// [0]   [1,0,0,0,0,0]
	// 可以选则用 1 或者不用 1；
	// 用 1 的话，如果 1 > j, 那么凑不成，即 0 ; 如果 1 <= j, 那么 [0,1] 的组合数量是 [0, j-1] 的数量
	// 不用 1 的话，就是 [0] 的组合数量
	// 两个相加，就得到了 [0,1] 凑成 j 的组合数
	// [0,1] [1+0,0+dp[0],0+dp[1]...]
	// 过程跟二维数组一样，区别是，只保留最新的结果就行
	for i, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

// @lc code=end

