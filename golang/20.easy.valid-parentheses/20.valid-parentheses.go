/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (43.91%)
 * Likes:    4399
 * Dislikes: 0
 * Total Accepted:    1.8M
 * Total Submissions: 4M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: s = "(]"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 *
 *
 */
package main

// @lc code=start
var opens = map[rune]struct{}{
	'(': {},
	'{': {},
	'[': {},
}

var closed = map[rune]struct{}{
	')': {},
	'}': {},
	']': {},
}

func isPair(o, c rune) bool {
	return (o == '(' && c == ')') || (o == '{' && c == '}') || (o == '[' && c == ']')
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, rune := range s {
		if _, ok := opens[rune]; ok {
			stack = append(stack, rune)
			continue
		}
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if isPair(top, rune) {
				continue
			} else {
				return false
			}
		} else {
			return false
		}

	}
	return len(stack) == 0
}

// @lc code=end
