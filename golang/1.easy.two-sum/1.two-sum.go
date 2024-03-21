/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.cn/problems/two-sum/description/
 *
 * algorithms
 * Easy (53.46%)
 * Likes:    18409
 * Dislikes: 0
 * Total Accepted:    5.3M
 * Total Submissions: 9.9M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers nums and an integer target, return indices of the
 * two numbers such that they add up to target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * You can return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,2,4], target = 6
 * Output: [1,2]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [3,3], target = 6
 * Output: [0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * Only one valid answer exists.
 *
 *
 *
 * Follow-up: Can you come up with an algorithm that is less than O(n^2) time
 * complexity?
 * nums[i] + nums[j] = target
 * meaning value of nums[i], target-nums[j] would occur twice
 * using a hashing to check
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// key num, value as index
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		counterNum := target - num
		if _, ok := m[counterNum]; ok {
			return []int{m[counterNum], i}
		}
		m[num] = i
	}
	return []int{}
}

// @lc code=end

