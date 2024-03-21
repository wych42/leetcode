/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.cn/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (66.36%)
 * Likes:    902
 * Dislikes: 0
 * Total Accepted:    749.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t, return true if t is an anagram of s, and false
 * otherwise.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * Example 2:
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 *
 *
 *
 * Follow up: What if the inputs contain Unicode characters? How would you
 * adapt your solution to such a case?
 *
 * construct a map `m` contains each char in `s`, with count of each char
 * iterate each char in t, substract count, return false if count < 0
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		if cnt, ok := m[r]; ok {
			cnt--
			if cnt < 0 {
				return false
			}
			m[r] = cnt
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	m := make(map[rune]int, len(s))
// 	for _, r := range s {
// 		m[r]++
// 	}
// 	n := make(map[rune]int, len(t))
// 	for _, r := range t {
// 		n[r]++
// 	}
// 	for k, v := range m {
// 		if v != n[k] {
// 			return false
// 		}
// 	}
// 	return true
// }