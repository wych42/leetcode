/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 *
 * https://leetcode.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (23.65%)
 * Likes:    1600
 * Dislikes: 3226
 * Total Accepted:    520.5K
 * Total Submissions: 2.1M
 * Testcase Example:  '"the sky is blue"'
 *
 * Given an input string s, reverse the order of the words.
 *
 * A word is defined as a sequence of non-space characters. The words in s will
 * be separated by at least one space.
 *
 * Return a string of the words in reverse order concatenated by a single
 * space.
 *
 * Note that s may contain leading or trailing spaces or multiple spaces
 * between two words. The returned string should only have a single space
 * separating the words. Do not include any extra spaces.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "the sky is blue"
 * Output: "blue is sky the"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "  hello world  "
 * Output: "world hello"
 * Explanation: Your reversed string should not contain leading or trailing
 * spaces.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "a good   example"
 * Output: "example good a"
 * Explanation: You need to reduce multiple spaces between two words to a
 * single space in the reversed string.
 *
 *
 * Example 4:
 *
 *
 * Input: s = "  Bob    Loves  Alice   "
 * Output: "Alice Loves Bob"
 *
 *
 * Example 5:
 *
 *
 * Input: s = "Alice does not even like bob"
 * Output: "bob like even not does Alice"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s contains English letters (upper-case and lower-case), digits, and spaces '
 * '.
 * There is at least one word in s.
 *
 *
 *
 * Follow up: Could you solve it in-place with O(1) extra space?
 *
 */

// @lc code=start
import "strings"

func reverseWords(s string) string {
	/*
		最简单的解法：strings.Split(s, " "), reverse, strings.Join
		O(1) 解法:
		- 从后往前遍历
		- 跳过所有的 " "
		- 碰到字符时，开始找到下一个空格的 idx, 中间的就是单词
		- 把单词加到结果里
	*/
	b := strings.Builder{}
	for i := len(s) - 1; i >= 0; {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		if i < 0 {
			break // 说明一直到 0 都是空白
		}
		// s[i] != ' '; 即单词到 i 结束
		end := i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		// s[i] == ' '; 即单词从 i+1 开始
		b.WriteString(s[i+1 : end+1])
		b.WriteString(" ")
	}
	ans := b.String()
	return ans[:len(ans)-1]
}

// @lc code=end

