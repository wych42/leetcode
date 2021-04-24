/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 *
 * https://leetcode.com/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (47.06%)
 * Likes:    2223
 * Dislikes: 251
 * Total Accepted:    171K
 * Total Submissions: 363.3K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * Given an array of distinct integers nums and a target integer target, return
 * the number of possible combinations that add up to target.
 *
 * The answer is guaranteed to fit in a 32-bit integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3], target = 4
 * Output: 7
 * Explanation:
 * The possible combination ways are:
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * Note that different sequences are counted as different combinations.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [9], target = 3
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * 1 <= nums[i] <= 1000
 * All the elements of nums are unique.
 * 1 <= target <= 1000
 *
 *
 *
 * Follow up: What if negative numbers are allowed in the given array? How does
 * it change the problem? What limitation we need to add to the question to
 * allow negative numbers?
 *
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	/*
		跟换零钱2的差异在于，同一个组合的不同排列，是要重复计数的。
		问题拆解:
		要凑 target 值，从 0 到 len(nums)-1 -> i
		- 如果选定了 x=nums[i], 产生的组合数为 target=target-x
		相加即为结果:
		[1,2,3], 4 为例:
		target=0, 只有一种可能，即什么都不选
		target=1
			选 1，即 target=0, 有1种可能
		target=2
			选 1，即 target=1, 有1种可能
			选 2，即 target=0, 有1种可能
			共4种可能
		target=3
			选 1，即 target=2, 有2种可能
			选 2，即 target=1, 有1种可能
			选 3，即 target=0, 有1种可能
			共4种可能
		target=3
			选 1，即 target=3, 有4种可能
			选 2，即 target=2, 有2种可能
			选 3，即 target=1, 有1种可能
			选 4，即 target=0, 有1种可能
			共7种可能
		时间复杂度为 O(target*len(nums))
	*/
	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for _, num := range nums {
			if num > t { //
				continue
			}
			dp[t] += dp[t-num]
		}
	}
	return dp[target]
}

// @lc code=end

