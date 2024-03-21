/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (30.47%)
 * Likes:    10243
 * Dislikes: 669
 * Total Accepted:    1.3M
 * Total Submissions: 4.1M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, return the longest palindromic substring in s.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "cbbd"
 * Output: "bb"
 *
 *
 * Example 3:
 *
 *
 * Input: s = "a"
 * Output: "a"
 *
 *
 * Example 4:
 *
 *
 * Input: s = "ac"
 * Output: "a"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consist of only digits and English letters (lower-case and/or upper-case),
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	/*
		从子问题出发：s[i:j] 是回文的前提是 s[i+1:j-1] 是回文，并且 s[i]==s[j]
	*/
	var ans string
	var max int
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for l := 1; l <= len(s); l++ { // 遍历所有长度的字串
		for start := 0; start < len(s); start++ { // 遍历开始的下标
			end := start + l - 1
			if end >= len(s) { // 越界了
				break
			}

			if l == 1 { // 长度是1，是回文
				dp[start][end] = true
			} else if l == 2 {
				dp[start][end] = s[start] == s[end]
			} else {
				dp[start][end] = s[start] == s[end] && dp[start+1][end-1]
			}
			if dp[start][end] && l > max {
				ans = s[start : end+1]
				max = l
			}
		}
	}
	return ans
}

// @lc code=end

