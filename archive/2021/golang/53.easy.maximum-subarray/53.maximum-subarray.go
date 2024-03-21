/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (47.70%)
 * Likes:    11466
 * Dislikes: 552
 * Total Accepted:    1.4M
 * Total Submissions: 2.9M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [5,4,-1,7,8]
 * Output: 23
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -10^5 <= nums[i] <= 10^5
 *
 *
 *
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution using the divide and conquer approach, which is more subtle.
 */

// @lc code=start
const min = -1 << 31

func maxSubArray(nums []int) int {
	/*
		注意要求连续
		拆分成子问题：如果直到 nums[:i-1] 的结果 maxi，对于 i:
		有两个选择，直接用 nums[i]; 跟前面保持连接nums[i]+dp[i-1]
		如果直接用比之前的大，就断掉了之前的，从当前位置开始；
		如果加起来比当前数大，说明在目前的情况下还是递增的；
		用一个变量存储当前最大值
	*/
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[0] = min
	ans := min
	for i := 1; i <= len(nums); i++ {
		dp[i] = max(nums[i-1], nums[i-1]+dp[i-1])
		ans = max(ans, dp[i])
	}

	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

