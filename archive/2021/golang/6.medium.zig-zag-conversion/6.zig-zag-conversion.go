/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (37.89%)
 * Likes:    2299
 * Dislikes: 5716
 * Total Accepted:    562.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * Example 3:
 *
 *
 * Input: s = "A", numRows = 1
 * Output: "A"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consists of English letters (lower-case and upper-case), ',' and '.'.
 * 1 <= numRows <= 1000
 *
 *
 */

// @lc code=start
func convert(s string, numRows int) string {
	/*
		尝试找到每一行的数据:
		一开始每一个字符都是向下排列的；
		当行数 == numRows 时，开始向上排列
	*/
	if numRows == 1 {
		return s
	}
	grid := make([][]rune, numRows)
	downwards := true
	row := 0
	for _, r := range s {
		grid[row] = append(grid[row], r)
		if row == numRows-1 {
			downwards = false
		} else if row == 0 {
			downwards = true
		}
		if downwards {
			row++
		} else {
			row--
		}
	}
	ans := ""
	for _, row := range grid {
		ans += string(row)
	}
	return ans
}

// @lc code=end

