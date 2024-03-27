/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.63%)
 * Likes:    3541
 * Dislikes: 0
 * Total Accepted:    816.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '3'
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// we have n open'(' and n close')'
	// add ( if open < n
	// add ) if close < open
	// do this start 0 and recursively
	result := make([]string, 0)
	temp := ""

	generate(n, temp, 0, 0, &result)
	return result
}

func generate(n int, temp string, openCnt, closeCnt int, result *[]string) {
	if openCnt == n && closeCnt == n {
		*result = append(*result, temp)
		return
	}
	if openCnt < n {
		// add ( to temp str
		temp += "("
		// calculate next
		generate(n, temp, openCnt+1, closeCnt, result)
		// remove appended (, continue next
		temp = temp[:len(temp)-1]
	}
	if closeCnt < openCnt {
		// add ( to temp str
		temp += ")"
		// calculate next
		generate(n, temp, openCnt, closeCnt+1, result)
		// remove appended (, continue next
		temp = temp[:len(temp)-1]
	}
}

// @lc code=end

