/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (43.98%)
 * Likes:    6860
 * Dislikes: 155
 * Total Accepted:    522.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an integer array nums, return the length of the longest strictly
 * increasing subsequence.
 *
 * A subsequence is a sequence that can be derived from an array by deleting
 * some or no elements without changing the order of the remaining elements.
 * For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,0,3,2,3]
 * Output: 4
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [7,7,7,7,7,7,7]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up:
 *
 *
 * Could you come up with the O(n^2) solution?
 * Could you improve it to O(n log(n)) time complexity?
 *
 *
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// 子问题：到 nums[:i] 的最长递增子序列长度
	// dp[i] = max(dp[j]) + 1; 0<j<=i-1 if dp[i] > dp[j]

	if len(nums) == 1 {
		return len(nums)
	}
	// 初始的时候每个节点的最长递增子序列长度都是 1
	lis := make([]int, len(nums))
	var result int
	for i, num := range nums {
		lis[i] = 1
		// 到 i 的时候，比较 j = [0:i-1]
		// 如果 nums[j] >= nums[i], 那么 i 对递增子序列没有帮助，跳过
		// 如果 nums[j] < nums[i], 那么 i 对递增子序列有用， lis[j]+1 是这种条件下的值，与 lis[i] 取大的那个
		for j := 0; j < i; j++ {
			if num > nums[j] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
		// 更新结果
		result = max(result, lis[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

