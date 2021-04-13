/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (33.64%)
 * Likes:    5538
 * Dislikes: 977
 * Total Accepted:    468.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array nums, find the smallest missing positive
 * integer.
 *
 *
 * Example 1:
 * Input: nums = [1,2,0]
 * Output: 3
 * Example 2:
 * Input: nums = [3,4,-1,1]
 * Output: 2
 * Example 3:
 * Input: nums = [7,8,9,11,12]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 300
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you implement an algorithm that runs in O(n) time and uses
 * constant extra space?
 *
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	/*
		把数组当作hash，下标就是hash func
		把每一个数放到与它的值相同的hash位置上
		遍历数组，找到第一个与下标值不同的
	*/
	for i := range nums {
		// num > 0 && num <= len(nums)
		// 应该把它放到 i-1 的位置上
		// 如果 i-1 的位置上的数，跟当前数一样，也不用交换
		for nums[i] > 0 && nums[i] <= len(nums) {
			if nums[i] == nums[nums[i]-1] { // nums[i] 对应的index上的数
				break
			}
			idx := nums[i] - 1
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}

	}
	return len(nums) + 1
}

// @lc code=end

