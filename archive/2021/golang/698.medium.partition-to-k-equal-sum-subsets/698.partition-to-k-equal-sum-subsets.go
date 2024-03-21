/*
 * @lc app=leetcode id=698 lang=golang
 *
 * [698] Partition to K Equal Sum Subsets
 *
 * https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/
 *
 * algorithms
 * Medium (45.72%)
 * Likes:    2958
 * Dislikes: 196
 * Total Accepted:    127.3K
 * Total Submissions: 283.6K
 * Testcase Example:  '[4,3,2,3,5,2,1]\n4'
 *
 * Given an integer array nums and an integer k, return true if it is possible
 * to divide this array into k non-empty subsets whose sums are all equal.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [4,3,2,3,5,2,1], k = 4
 * Output: true
 * Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3),
 * (2,3) with equal sums.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3,4], k = 3
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= nums.length <= 16
 * 1 <= nums[i] <= 10^4
 * The frequency of each element is in the range [1, 4].
 *
 *
 */

// @lc code=start

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	/*
		k 个非空子集，每一个的和还要相等，那么每一个的和应该是 avg = sum(nums)/k

		nums 里都是正数，可以先对 nums 排序，如果一个元素 > avg, 那么肯定就分不了了

		其余的情况用回溯解决
	*/
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return backtrack(nums, 0, make([]int, k), sum/k)
}

func backtrack(nums []int, idx int, groups []int, limit int) bool {
	// caputre: 都分完了
	if idx == len(nums) {
		return true
	}
	num := nums[idx]
	// constrain: 这个数比 limit 大，那么肯定无法分配了
	if num > limit {
		return false
	}
	// 尝试把 num 分到每一个组里
	for i := range groups {
		if groups[i]+num <= limit {
			// choose: 选择 groups[i]
			groups[i] += num
			if backtrack(nums, idx+1, groups, limit) {
				// 成功分配，直接返回
				return true
			}
			// 此路不通
			groups[i] -= num

			// 把 num 分配给任何一个 >=0 的组，都不会成功
			if groups[i] == 0 {
				break
			}
		}
	}
	return false
}

// @lc code=end

