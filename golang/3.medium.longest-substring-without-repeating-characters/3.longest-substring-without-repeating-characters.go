/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (31.39%)
 * Likes:    13772
 * Dislikes: 713
 * Total Accepted:    2.1M
 * Total Submissions: 6.7M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string s, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * Notice that the answer must be a substring, "pwke" is a subsequence and not
 * a substring.
 *
 *
 * Example 4:
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
 * 0 <= s.length <= 5 * 10^4
 * s consists of English letters, digits, symbols and spaces.
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 用一个 map 存访问过的元素的坐标
	// 用 slow, fast 两者指针作为滑动窗口
	// 没有发现重复时，fast向前
	// 发现重复时: fast-slow 为当前窗口找到的不重复字串, 与 max 比较
	// slow 需要移动到重复字符的下一个元素（但是不能向后移动) 即 idx+1 > slow

	runes := []rune(s)
	if len(runes) == 0 || len(runes) == 1 {
		return len(runes)
	}
	max := 0
	m := make(map[rune]int) // 元素上次出现的坐标
	slow, fast := 0, 0
	for ; fast < len(runes); fast++ {
		// 没有访问过改元素，记录坐标
		r := runes[fast]
		idx, ok := m[r]
		if ok {
			// 访问过了，比较 count 和 max
			if fast-slow > max {
				max = fast - slow
			}
			// 从重复元素的下一个开始计算
			if idx+1 > slow {
				slow = idx + 1
			}
		}
		m[r] = fast
	}
	if fast-slow > max {
		max = fast - slow
	}
	return max
}

// @lc code=end

