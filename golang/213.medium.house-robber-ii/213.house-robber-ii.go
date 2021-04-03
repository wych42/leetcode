/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 *
 * https://leetcode.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (37.52%)
 * Likes:    2730
 * Dislikes: 64
 * Total Accepted:    235.6K
 * Total Submissions: 625.9K
 * Testcase Example:  '[2,3,2]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed. All houses at this place are
 * arranged in a circle. That means the first house is the neighbor of the last
 * one. Meanwhile, adjacent houses have a security system connected, and it
 * will automatically contact the police if two adjacent houses were broken
 * into on the same night.
 *
 * Given an integer array nums representing the amount of money of each house,
 * return the maximum amount of money you can rob tonight without alerting the
 * police.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money
 * = 2), because they are adjacent houses.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 1000
 *
 *
 */

// @lc code=start
func rob(nums []int) int {
	// nums 是个环，意味着不同能是抢头尾
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robber(nums[1:]), robber(nums[:len(nums)-1]))
}

func robber(nums []int) int {
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
	for -, num := range nums {
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

