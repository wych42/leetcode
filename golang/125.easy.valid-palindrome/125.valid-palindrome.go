/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.93%)
 * Likes:    735
 * Dislikes: 0
 * Total Accepted:    571.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * A phrase is a palindrome if, after converting all uppercase letters into
 * lowercase letters and removing all non-alphanumeric characters, it reads the
 * same forward and backward. Alphanumeric characters include letters and
 * numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric
 * characters.
 * Since an empty string reads the same forward and backward, it is a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s consists only of printable ASCII characters.
 *
 *
 * use two pointers iterate from start and end, skip non-aplphnumberic char, and compare
 */

package main

// @lc code=start
import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	runes := []rune(s)
	sp := 0
	ep := len(s) - 1
	result := true
	for sp < ep {
		fmt.Println(sp, ep)
		if !isAlphanumeric(unicode.ToLower(runes[sp])) {
			sp++
			continue
		}
		if !isAlphanumeric(unicode.ToLower(runes[ep])) {
			ep--
			continue
		}
		if unicode.ToLower(runes[sp]) == unicode.ToLower(runes[ep]) {
			sp++
			ep--
			continue
		}
		result = false
		break
	}
	return result

}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

// func main() {
// 	isPalindrome("race a car")
// }

// @lc code=end
