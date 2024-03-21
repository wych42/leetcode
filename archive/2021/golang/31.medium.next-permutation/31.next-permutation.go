/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (33.77%)
 * Likes:    5325
 * Dislikes: 1828
 * Total Accepted:    501K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such an arrangement is not possible, it must rearrange it as the lowest
 * possible order (i.e., sorted in ascending order).
 *
 * The replacement must be in place and use only constant extra memory.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [1,3,2]
 * Example 2:
 * Input: nums = [3,2,1]
 * Output: [1,2,3]
 * Example 3:
 * Input: nums = [1,1,5]
 * Output: [1,5,1]
 * Example 4:
 * Input: nums = [1]
 * Output: [1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
func nextPermutation(nums []int) {
	/*
		将输入的数进行全排列，可以发现一些规律:
		1234
		1243
		1324
		1342
		1423
		3421 -> 4321
		4123
		...
		下一个更大的排列，是这样得到的：
		- 从后往前，找连续的升序子序列，直到第一个打破升序的数（idx)
		- 再从后往前，找到第一个大于 idx 值的数, 与 idx 交换
		- 把升序子序列，进行正向的从小到大的排列
	*/
	idx := len(nums) - 2                                // 处理最前面的边界
	for ; idx >= 0 && nums[idx] >= nums[idx+1]; idx-- { // 找到打破升序的元素
	}
	if idx >= 0 {
		for j := len(nums) - 1; j > idx; j-- {
			if nums[j] > nums[idx] {
				// 找到要交换的元素
				nums[idx], nums[j] = nums[j], nums[idx]
				break
			}
		}
	}

	// 直接翻转升序序列
	l := idx + 1
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return
}

// @lc code=end

