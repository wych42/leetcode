/*
 * @lc app=leetcode id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/number-of-longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (38.48%)
 * Likes:    2181
 * Dislikes: 118
 * Total Accepted:    76.9K
 * Total Submissions: 198.9K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * Given an integer array nums, return the number of longest increasing
 * subsequences.
 *
 * Notice that the sequence has to be strictly increasing.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,3,5,4,7]
 * Output: 2
 * Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and
 * [1, 3, 5, 7].
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,2,2,2,2]
 * Output: 5
 * Explanation: The length of longest continuous increasing subsequence is 1,
 * and there are 5 subsequences' length is 1, so output 5.
 *
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2000
 * -10^6 <= nums[i] <= 10^6
 *
 *
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	// 子问题：到 nums[:i] 的最长递增子序列长度
	// dp[i] = max(dp[j]) + 1; 0<j<=i-1 if dp[i] > dp[j]

	if len(nums) == 1 {
		return len(nums)
	}
	// 初始的时候每个节点的最长递增子序列长度都是 1
	lis := make([]int, len(nums))
	// 次数
	count := make([]int, len(nums))

	var maxLength int
	for i, num := range nums {
		lis[i] = 1
		count[i] = 1
		// dp[i] = 1
		// dp[1]++ // 1 出现的次数

		// 到 i 的时候，比较 j = [0:i-1]
		// 如果 nums[i] < nums[j], 那么 i 对递增子序列没有帮助，跳过
		// 如果 nums[i] > nums[j], 那么 i 对递增子序列有用， lis[j]+1 是这种条件下的值，与 lis[i] 取大的那个

		// lis[i] 就是局部最大长度
		for j := 0; j < i; j++ {
			if num > nums[j] {
				if lis[j]+1 > lis[i] { // 最大值变化了, 次数重置为改变最大值的那个
					lis[i] = lis[j] + 1
					count[i] = count[j]
				} else if lis[j]+1 == lis[i] {
					// 最大值变化了, 次数加上
					count[i] += count[j]
				}
			}
		}
		// 更新结果
		if lis[i] > maxLength {
			maxLength = lis[i]
		}
	}
	res := 0
	for i, val := range lis {
		if val == maxLength {
			res += count[i]
		}
	}
	return res
}

// @lc code=end

