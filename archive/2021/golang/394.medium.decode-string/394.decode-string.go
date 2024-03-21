/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 *
 * https://leetcode.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (52.63%)
 * Likes:    4606
 * Dislikes: 224
 * Total Accepted:    301.6K
 * Total Submissions: 573K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * Given an encoded string, return its decoded string.
 * 
 * The encoding rule is: k[encoded_string], where the encoded_string inside the
 * square brackets is being repeated exactly k times. Note that k is guaranteed
 * to be a positive integer.
 * 
 * You may assume that the input string is always valid; No extra white spaces,
 * square brackets are well-formed, etc.
 * 
 * Furthermore, you may assume that the original data does not contain any
 * digits and that digits are only for those repeat numbers, k. For example,
 * there won't be input like 3a or 2[4].
 * 
 * 
 * Example 1:
 * Input: s = "3[a]2[bc]"
 * Output: "aaabcbc"
 * Example 2:
 * Input: s = "3[a2[c]]"
 * Output: "accaccacc"
 * Example 3:
 * Input: s = "2[abc]3[cd]ef"
 * Output: "abcabccdcdcdef"
 * Example 4:
 * Input: s = "abc3[cd]xyz"
 * Output: "abccdcdcdxyz"
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 30
 * s consists of lowercase English letters, digits, and square brackets
 * '[]'.
 * s is guaranteed to be a valid input.
 * All the integers in s are in the range [1, 300].
 * 
 * 
 */

package main
// import "fmt"

/*
1. 用两个 strings.Builer 保存临时的倍数(count)和字符串(segs)
在遇到 [ 时，把倍数 Builder 清空，入队，此时对应的字符串为 ""
遇到数字或者 ] 时，字符串的 Builder 清空, 入队。
2. 遇到 ] 时，向前寻找到匹配的 [(“”字符串为标记),以及对应的倍数,之间的字符串拼接，重复。
*/
// @lc code=start
import "strings"
import "strconv"

func decodeString(s string) string {
	segs := make([]string, 0)
	counts := make([]int, 0)
	b := strings.Builder{}
	d := strings.Builder{}
	for _, char := range s {
		switch c:=char; {
		case c <= '9' && c >='0':
			if b.Len() > 0 {
				segs = append(segs, b.String())
				counts = append(counts, 1)
				b.Reset()
			}
			d.WriteRune(c)
		case c == '[':
			n, _ := strconv.Atoi(d.String())
			d.Reset()
			counts = append(counts, n)
			segs = append(segs, "")
		case c == ']':
			if b.Len() > 0 {
				segs = append(segs, b.String())
				counts = append(counts, 1)
				b.Reset()
			}
			idx := 0
			for l:=len(segs)-1;l>0;l--{
				if segs[l] == "" {
					idx=l
					break
				}
			}
			segs[idx]=strings.Repeat(strings.Join(segs[idx+1:],""), counts[idx])
			segs = segs[:idx+1]
			counts = counts[:idx+1]
		default:
			b.WriteRune(c)
		}
	}
	return strings.Join(segs, "")+b.String()
}
// @lc code=end

func main() {
	fmt.Println(decodeString("a2[b3[cd]e]f"))
}

