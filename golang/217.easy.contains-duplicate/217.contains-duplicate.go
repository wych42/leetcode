/*
* easy question, use hashing.
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.cn/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (55.15%)
 * Likes:    1079
 * Dislikes: 0
 * Total Accepted:    849.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an integer array nums, return true if any value appears at least twice
 * in the array, and return false if every element is distinct.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,1]
 * Output: true
 * Example 2:
 * Input: nums = [1,2,3,4]
 * Output: false
 * Example 3:
 * Input: nums = [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
*/

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
			return true
		} else {
			m[num] = 1
		}
	}
	return false
}

// @lc code=end

