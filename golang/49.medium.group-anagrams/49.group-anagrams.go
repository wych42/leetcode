package main

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

import "strconv"

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)
	for _, str := range strs {
		// for each str, find its singature
		signature := make(map[rune]int) // only contains lowercase letter
		for _, r := range str {
			signature[r]++
		}
		sign := ""
		for i := 'a'; i <= 'z'; i++ {
			sign += string(i) + strconv.Itoa(signature[i])
		}

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

// @lc code=end
