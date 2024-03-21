/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (49.33%)
 * Likes:    2848
 * Dislikes: 78
 * Total Accepted:    443.9K
 * Total Submissions: 892.9K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers, nums, that might contain duplicates, return
 * all possible unique permutations in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,2]
 * Output:
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start

// 不用 map 怎么做呢？
func permuteUnique(nums []int) [][]int {
	// 有n个格子填空
	// 选中一个放第一格, 记录用过
	// 用剩下的填其余的格子
	// 再选择数填当前格子的时候，要判断有没有用过
	res := make([][]int, 0)
	used := make([]bool, len(nums)) // 记录一次循环内，用过了哪些值
	backtrack(&nums, []int{}, &used, &res)
	return res
}

func backtrack(nums *[]int, curr []int, used *[]bool, res *[][]int) {
	if len(curr) == len(*nums) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
		return
	}

	// 记录当前格子，已经用过哪些了
	m := make(map[int]bool)
	for i, num := range *nums {
		if m[num] { // 用过了就不处理
			continue
		}
		if used[i] {
			continue
		}
		m[num] = true
		used[i] = true
		backtrack(nums, append(curr, num), used, res)
		used[i] = false
	}
	return
}

// @lc code=end

