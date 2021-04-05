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
	m := make(map[string]bool)
	m[""] = true
	var ans string
	for i, r := range s {
		m[string(r)] = true // 单字符是回文
		if ans == "" {
			ans = string(r)
		}
		for j := 0; j < i; j++ {
			str := s[j : i+1] // include j
			subStr := s[j+1 : i]
			if s[j] == byte(r) && m[subStr] { // 首尾相等，检查内部字符串是不是回文
				if len(str) > len(ans) {
					ans = str
				}
				m[str] = true
			} else {
				m[str] = false
			}
		}
	}
	return ans
}

// @lc code=end

