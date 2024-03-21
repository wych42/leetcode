/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (48.62%)
 * Likes:    6315
 * Dislikes: 200
 * Total Accepted:    931.8K
 * Total Submissions: 1.9M
 * Testcase Example:  '2'
 *
 * You are climbing a staircase. It takes n steps to reach the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 45
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	// 对第i级台阶，上一步可以选择 i-1;i-2
	// dp[0] = 1
	// dp[1] = dp[0] + dp[-1] = 1
	// dp[2] = dp[1] + dp[0] = 2
	// dp[3] = dp[2] + dp[1] = 3
	// dp[4] = dp[3] + dp[2]
	// 好像是个斐波那契数列？
	prev, curr, next := 0, 0, 1
	for i := 1; i <= n; i++ {
		prev = curr
		curr = next
		next = prev + curr
	}
	return next
}

// @lc code=end

