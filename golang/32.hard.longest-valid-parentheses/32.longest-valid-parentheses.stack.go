/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (29.26%)
 * Likes:    5042
 * Dislikes: 182
 * Total Accepted:    367.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()".
 *
 *
 * Example 2:
 *
 *
 * Input: s = ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()".
 *
 *
 * Example 3:
 *
 *
 * Input: s = ""
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 3 * 10^4
 * s[i] is '(', or ')'.
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	/*
		最长的距离实际是：在 ) 出现，并且匹配时，最近一个没有匹配的 ( 的距离
		如果 ( 都被匹配了，需要一个标记来区分, 防止匹配到上一个连续, 比如 ()(()
		这个标记就用栈底元素记录最近一个落单的 ), 因为落单的 ) 中断了有效序列
	*/
	ans := 0
	stack := make([]int, 0, len(s))
	stack = append(stack, -1) // 假装上一个没匹配
	for i, c := range s {
		switch c {
		case '(':
			stack = append(stack, i)
		case ')':
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

