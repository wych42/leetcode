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

func permuteUnique(nums []int) [][]int {
	// 用一个map记录每个数的次数
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	res := make([][]int, 0)           // 存结果
	curr := make([]int, 0, len(nums)) // 临时结果
	var dfs func()
	dfs = func() {
		if len(curr) == len(nums) { // 临时结果长度符合要求了，就拷贝一份写到结果里
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		// 限制了只能从 map 里挑元素填空
		for num, count := range m {
			if count > 0 {
				// 挑了一个元素后，把当前的元素 count - 1，递归填剩下的空
				curr = append(curr, num)
				m[num]--
				dfs()
				// 填完一轮，把结果清空
				m[num]++
				curr = curr[:len(curr)-1]
			}
		}
	}
	dfs()
	return res
}

// @lc code=end

