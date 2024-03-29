/*
 * @lc app=leetcode id=1486 lang=golang
 *
 * [1486] XOR Operation in an Array
 *
 * https://leetcode.com/problems/xor-operation-in-an-array/description/
 *
 * algorithms
 * Easy (84.00%)
 * Likes:    476
 * Dislikes: 216
 * Total Accepted:    87K
 * Total Submissions: 103.5K
 * Testcase Example:  '5\n0'
 *
 * Given an integer n and an integer start.
 *
 * Define an array nums where nums[i] = start + 2*i (0-indexed) and n ==
 * nums.length.
 *
 * Return the bitwise XOR of all elements of nums.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 5, start = 0
 * Output: 8
 * Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^
 * 8) = 8.
 * Where "^" corresponds to bitwise XOR operator.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 4, start = 3
 * Output: 8
 * Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.
 *
 * Example 3:
 *
 *
 * Input: n = 1, start = 7
 * Output: 7
 *
 *
 * Example 4:
 *
 *
 * Input: n = 10, start = 5
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 1000
 * 0 <= start <= 1000
 * n == nums.length
 *
 */

// @lc code=start
func xorOperation(n int, start int) int {
	ans := start
	for i := 1; i < n; i++ {
		ans ^= start + 2*i
	}
	return ans
}

// @lc code=end

