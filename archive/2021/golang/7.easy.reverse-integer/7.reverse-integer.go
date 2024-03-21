/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.90%)
 * Likes:    4546
 * Dislikes: 6962
 * Total Accepted:    1.5M
 * Total Submissions: 5.7M
 * Testcase Example:  '123'
 *
 * Given a signed 32-bit integer x, return x with its digits reversed. If
 * reversing x causes the value to go outside the signed 32-bit integer range
 * [-2^31, 2^31 - 1], then return 0.
 *
 * Assume the environment does not allow you to store 64-bit integers (signed
 * or unsigned).
 *
 *
 * Example 1:
 * Input: x = 123
 * Output: 321
 * Example 2:
 * Input: x = -123
 * Output: -321
 * Example 3:
 * Input: x = 120
 * Output: 21
 * Example 4:
 * Input: x = 0
 * Output: 0
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 */

// @lc code=start

const maxInt = 1 << 31

// const minInt = -1 << 31

func reverse(x int) int {
	// 小的直接返回
	if x <= 9 && x >= -9 {
		return x
	}

	// 只处理正数
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	}

	rev := 0

	for x > 0 {
		// 先检查是否可能超过 maxInt
		mod := x % 10
		// rev * 10 > maxInt || rev*10 + mod > maxInt
		if rev > maxInt/10 || (maxInt-mod)/10 < rev {
			return 0
		}

		rev = rev*10 + mod
		x = x / 10
	}
	return rev * flag
}

// @lc code=end

