/*
 * @lc app=leetcode id=1704 lang=golang
 *
 * [1704] Determine if String Halves Are Alike
 *
 * https://leetcode.com/problems/determine-if-string-halves-are-alike/description/
 *
 * algorithms
 * Easy (77.45%)
 * Likes:    222
 * Dislikes: 14
 * Total Accepted:    36.2K
 * Total Submissions: 46K
 * Testcase Example:  '"book"'
 *
 * You are given a string s of even length. Split this string into two halves
 * of equal lengths, and let a be the first half and b be the second half.
 *
 * Two strings are alike if they have the same number of vowels ('a', 'e', 'i',
 * 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s contains uppercase and
 * lowercase letters.
 *
 * Return true if a and b are alike. Otherwise, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "book"
 * Output: true
 * Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel.
 * Therefore, they are alike.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "textbook"
 * Output: false
 * Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2.
 * Therefore, they are not alike.
 * Notice that the vowel o is counted twice.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "MerryChristmas"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: s = "AbCdEfGh"
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= s.length <= 1000
 * s.length is even.
 * s consists of uppercase and lowercase letters.
 *
 *
 */

// @lc code=start
var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func halvesAreAlike(s string) bool {
	runes := []rune(s)
	var a, b int
	half := len(runes) / 2
	for i, c := range runes {
		if !vowels[c] {
			continue
		}
		if i < half {
			a++
		} else {
			b++
		}
	}
	return a == b
}

// @lc code=end

