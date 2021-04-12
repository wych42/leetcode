/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (65.21%)
 * Likes:    7532
 * Dislikes: 324
 * Total Accepted:    717.2K
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
	/*
		有n个左边和n个右边；
		先选一个；检查是否符合约束；符合就求子问题
		把选的这个拿出来，选另一个

		约束：生成的字符串里，( 的数量要一直 >= ) 的数量
	*/
	res := make([]string, 0)
	generate(0, 0, n, "", &res)
	return res
}

func generate(left, right, max int, str string, res *[]string) {
	if len(str) == 2*max {
		*res = append(*res, str)
		return
	}
	if left < max { // 此时可以添加 (
		str += "("
		generate(left+1, right, max, str, res)
		str = str[:len(str)-1]
	}
	if right < left { // 此时可以添加 )
		str += ")"
		generate(left, right+1, max, str, res)
		str = str[:len(str)-1]
	}
}

// @lc code=end

