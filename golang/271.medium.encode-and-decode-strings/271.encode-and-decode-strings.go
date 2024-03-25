/*
 * @lc app=leetcode.cn id=271 lang=golang
 *
 * [271] Encode and Decode Strings
 *
 * https://leetcode.cn/problems/encode-and-decode-strings/description/
 *
 * algorithms
 * Medium (58.62%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 10.5K
 * Testcase Example:  '["Hello","World"]'
 *
 * Design an algorithm to encode a list of strings to a string. The encoded
 * string is then sent over the network and is decoded back to the original
 * list of strings.
 *
 * Machine 1 (sender) has the function:
 *
 *
 * string encode(vector<string> strs) {
 * ⁠ // ... your code
 * ⁠ return encoded_string;
 * }
 * Machine 2 (receiver) has the function:
 *
 *
 * vector<string> decode(string s) {
 * ⁠ //... your code
 * ⁠ return strs;
 * }
 *
 *
 * So Machine 1 does:
 *
 *
 * string encoded_string = encode(strs);
 *
 *
 * and Machine 2 does:
 *
 *
 * vector<string> strs2 = decode(encoded_string);
 *
 *
 * strs2 in Machine 2 should be the same as strs in Machine 1.
 *
 * Implement the encode and decode methods.
 *
 * You are not allowed to solve the problem using any serialize methods (such
 * as eval).
 *
 *
 * Example 1:
 *
 *
 * Input: dummy_input = ["Hello","World"]
 * Output: ["Hello","World"]
 * Explanation:
 * Machine 1:
 * Codec encoder = new Codec();
 * String msg = encoder.encode(strs);
 * Machine 1 ---msg---> Machine 2
 *
 * Machine 2:
 * Codec decoder = new Codec();
 * String[] strs = decoder.decode(msg);
 *
 *
 * Example 2:
 *
 *
 * Input: dummy_input = [""]
 * Output: [""]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] contains any possible characters out of 256 valid ASCII
 * characters.
 *
 *
 *
 * Follow up: Could you write a generalized algorithm to work on any possible
 * set of characters?
 *
 */
package main

// @lc code=start
type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	values := make([]rune, 0)
	for _, str := range strs {
		values = append(values, rune(len(str)))
		values = append(values, []rune(str)...)
	}
	return string(values)
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	result := make([]string, 0)

	runes := []rune(strs)
	for i := 0; i < len(runes); {
		// fmt.Println(i, runes[i], size)
		size := int(runes[i])
		if size == 0 {
			result = append(result, "")
			i++
			continue
		}
		result = append(result, string(runes[i+1:i+1+size]))
		i = i + size + 1
		// reset size
		size = 0
	}
	return result
}

// @lc code=end
