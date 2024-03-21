/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Medium (42.85%)
 * Likes:    6766
 * Dislikes: 190
 * Total Accepted:    696.3K
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security systems
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given an integer array nums representing the amount of money of each house,
 * return the maximum amount of money you can rob tonight without alerting the
 * police.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 *
 *
 */

// @lc code=start
func rob(nums []int) int {
	// 子问题：走到 nums[i] 的时候，最大值是多少
	rob1, rob2 := 0, 0
	// 假设有一下下面的 dp，其实只要记住走到前两个屋子时的最大值，就可以计算出当前屋子的最大值
	// input: [1, 2, 3, 1]
	// dp:    [0, 0, 0, 0]
	// dp:    [1, 0, 0, 0]
	// dp:    [1, 2, 0, 0]
	// dp:    [1, 2, 1+3, 0]
	// dp:    [1, 3, 1+3, 1+3]
	// [1, rob1, rob2, 3,...]
	for _, num := range nums {
		// 对每个 num，可以选择 num+rob1, 或者 rob2
		curr := max(num+rob1, rob2)
		// 计算出来当前的最大值，把 rob1, rob2 往前移动
		rob1, rob2 = rob2, curr
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

