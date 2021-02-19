/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 *
 * https://leetcode.com/problems/binary-search/description/
 *
 * algorithms
 * Easy (48.14%)
 * Likes:    1153
 * Dislikes: 57
 * Total Accepted:    232.8K
 * Total Submissions: 430.9K
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * Given a sorted (in ascending order) integer array nums of n elements and a
 * target value, write a function to search target in nums. If target exists,
 * then return its index, otherwise return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,0,3,5,9,12], target = 9
 * Output: 4
 * Explanation: 9 exists in nums and its index is 4
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-1,0,3,5,9,12], target = 2
 * Output: -1
 * Explanation: 2 does not exist in nums so return -1
 *
 *
 *
 *
 * Note:
 *
 *
 * You may assume that all elements in nums are unique.
 * n will be in the range [1, 10000].
 * The value of each element in nums will be in the range [-9999, 9999].
 *
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	start := 0
	stop := len(nums) - 1
	for start+1 < stop {
		mid := start + (stop-start)/2
		switch v := nums[mid]; {
		case v == target:
			return mid
		case v > target:
			stop = mid
		case v < target:
			start = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[stop] == target {
		return stop
	}
	return -1
}

// @lc code=end

