/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (57.17%)
 * Likes:    1789
 * Dislikes: 3184
 * Total Accepted:    492.8K
 * Total Submissions: 859.2K
 * Testcase Example:  '3'
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 *
 *
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * For example, 2 is written as II in Roman numeral, just two one's added
 * together. 12 is written as XII, which is simply X + II. The number 27 is
 * written as XXVII, which is XX + V + II.
 *
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 *
 *
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 *
 *
 * Given an integer, convert it to a roman numeral.
 *
 *
 * Example 1:
 *
 *
 * Input: num = 3
 * Output: "III"
 *
 *
 * Example 2:
 *
 *
 * Input: num = 4
 * Output: "IV"
 *
 *
 * Example 3:
 *
 *
 * Input: num = 9
 * Output: "IX"
 *
 *
 * Example 4:
 *
 *
 * Input: num = 58
 * Output: "LVIII"
 * Explanation: L = 50, V = 5, III = 3.
 *
 *
 * Example 5:
 *
 *
 * Input: num = 1994
 * Output: "MCMXCIV"
 * Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= num <= 3999
 *
 *
 */

// @lc code=start
/*
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000

 4, IV, 9, IX
 40, XL, 90 XC
 400 CD, 900 CM

3 -> III
4 -> IV
9 -> IX
58 -> 50 L; 8-> 5 V, 3 -> III -> LVIII
1994 -> 1000 -> M -> 994
	 	-> 900 CM -> 94
		-> 90 XC -> 4
		-> 4 IV
	MCMXCIV

1 <= num <= 3999
把千、百、十、个位分别计算拼到一起。
*/

import "strings"

var m = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	ans := strings.Builder{}
	quo := num / 1000
	reminder := num % 1000
	if quo != 0 {
		ans.WriteString(strings.Repeat("M", quo))
	}
	ans.WriteString(reminder, 100)
	return ans.String()
}

func convert(num int, base int) string {
	if num == 0 || base == 0 {
		return ""
	}
	ans := strings.Builder{}
	quo := num / base
	reminder := num % base
	if quo != 0 {
		if str, ok := m[quo*base]; ok {
			ans.WriteString(str)
		} else {
			if quo < 4 {
				ans.WriteString(strings.Repeat(m[base], quo))
			} else {
				ans.WriteString(m[5*base])
				ans.WriteString(strings.Repeat(m[base], quo-5))
			}
		}
	}
	ans.WriteString(convert(reminder, base/10))
	return ans.String()
}

// @lc code=end

