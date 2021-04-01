/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (66.49%)
 * Likes:    5709
 * Dislikes: 131
 * Total Accepted:    786K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an array nums of distinct integers, return all the possible
 * permutations. You can return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * Example 2:
 * Input: nums = [0,1]
 * Output: [[0,1],[1,0]]
 * Example 3:
 * Input: nums = [1]
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * All the integers of nums are unique.
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrack func(rest []int, curr []int)
	backtrack = func(rest []int, curr []int) {
		if len(curr) == len(nums) { // 排列完了
			result = append(result, curr)
			return
		}
		for i := 0; i <= len(rest)-1; i++ {
			val := rest[i] // 单独拿出来，放在第一个

			// 递归算剩下的
			r := make([]int, 0, len(rest)-1)
			r = append(r, rest[:i]...)
			r = append(r, rest[i+1:]...)
			backtrack(r, append([]int{val}, curr...))
		}
	}
	backtrack(nums, []int{})
	return result
}

// @lc code=end

