/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (49.97%)
 * Likes:    5891
 * Dislikes: 520
 * Total Accepted:    823.3K
 * Total Submissions: 1.6M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent. Return the answer in
 * any order.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * Example 2:
 *
 *
 * Input: digits = ""
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] is a digit in the range ['2', '9'].
 *
 *
 */

// @lc code=start
var m = map[rune][]rune{
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	ans := make([]string, 0)
	dfs(digits, 0, []rune{}, &ans)
	return ans
}

func dfs(digits string, idx int, path []rune, ans *[]string) {
	if idx == len(digits) {
		*ans = append(*ans, string(path))
		return
	}
	for _, r := range m[rune(digits[idx])] {
		path = append(path, r)
		dfs(digits, idx+1, path, ans)
		path = path[:len(path)-1]
	}
}

// @lc code=end

