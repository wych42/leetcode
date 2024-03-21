/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (49.71%)
 * Likes:    3209
 * Dislikes: 1726
 * Total Accepted:    1.2M
 * Total Submissions: 2.4M
 * Testcase Example:  '121'
 *
 * Given an integer x, return true if x is palindrome integer.
 *
 * An integer is a palindrome when it reads the same backward as forward. For
 * example, 121 is palindrome while 123 is not.
 *
 *
 * Example 1:
 *
 *
 * Input: x = 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: x = -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: x = 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Example 4:
 *
 *
 * Input: x = -101
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without converting the integer to a string?
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x <= 9 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	// 1221；每次从x最好摘掉一位，生成新的数
	// 如果新生成的数，>= 当前的数了，就暂停（超过一半了)
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	//fmt.Println(x, reversed)
	return x == reversed || x == reversed/10
}

// @lc code=end

