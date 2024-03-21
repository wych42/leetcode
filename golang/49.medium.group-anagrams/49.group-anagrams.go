/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.cn/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (67.93%)
 * Likes:    1855
 * Dislikes: 0
 * Total Accepted:    654.5K
 * Total Submissions: 963.6K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
 *
 *
 */

// @lc code=start

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)
	for _, str := range strs {
		// for each str, check if it's anagram aganist each key in result map
		sign := getAnagramSignature(str)
		if _, ok := result[sign]; ok {
			result[sign] = append(result[sign], str)
		} else {
			result[sign] = []string{str}
		}
	}
	res := make([][]string, 0)
	for _, v := range result {
		res = append(res, v)
	}
	return res
}

func getAnagramSignature(s string) string {
	if len(s) == 0 {
		return s
	}
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}
	sign := ""
	for i := 'a'; i <= 'z'; i++ {
		if _, ok := m[i]; ok {
			sign += fmt.Sprintf("%c%d", i, m[i])
		}
	}
	return sign
}

// @lc code=end

