/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 *
 * https://leetcode.com/problems/largest-divisible-subset/description/
 *
 * algorithms
 * Medium (38.45%)
 * Likes:    1825
 * Dislikes: 90
 * Total Accepted:    111.3K
 * Total Submissions: 289.4K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct positive integers nums, return the largest subset
 * answer such that every pair (answer[i], answer[j]) of elements in this
 * subset satisfies:
 *
 *
 * answer[i] % answer[j] == 0, or
 * answer[j] % answer[i] == 0
 *
 *
 * If there are multiple solutions, return any of them.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3]
 * Output: [1,2]
 * Explanation: [1,3] is also accepted.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,4,8]
 * Output: [1,2,4,8]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 2 * 10^9
 * All the integers in nums are unique.
 *
 *
 */

// @lc code=start
import "sort"

func largestDivisibleSubset(nums []int) []int {
	/*
		参考 300.longest-increasing-subsequence
		最长递增的判断条件是 nums[j] > nums[i]; j > i
		最长divisible 的判断条件是 nums[j] % nums[i] == 0; j > i(列表有序)

		这样可以得到一个列表：在 i 处，最长 divisible 的数量是多少
	*/
	if len(nums) <= 1 {
		return nums
	}
	// 从小到大排列
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dp := make([]int, len(nums))    // 记录到 i 时，最长的长度是多少
	dpIdx := make([]int, len(nums)) // 记录到 i 时，最长长度的前一个下标是什么
	maxLength := 0
	maxIdx := 0
	for i, num := range nums {
		length := 1 // 如果都匹配不上，也就是当前元素对贡献最长序列没有帮助， 默认就是1
		prevIdx := i
		for j := 0; j < i; j++ {
			if num%nums[j] == 0 {
				if dp[j]+1 > length {
					length = dp[j] + 1
					prevIdx = j
				}
			}
		}
		if length > maxLength {
			maxLength = length
			maxIdx = i
		}
		dp[i] = length
		dpIdx[i] = prevIdx
	}
	res := make([]int, 0, dp[len(nums)-1])
	for len(res) != maxLength {
		res = append(res, nums[maxIdx])
		maxIdx = dpIdx[maxIdx]
	}
	return res
}

// @lc code=end

