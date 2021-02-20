/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (39.46%)
 * Likes:    1422
 * Dislikes: 267
 * Total Accepted:    234.9K
 * Total Submissions: 559.2K
 * Testcase Example:  '[1,3,5]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 *
 * Find the minimum element.
 *
 * The array may contain duplicates.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5]
 * Output: 1
 *
 * Example 2:
 *
 *
 * Input: [2,2,2,0,1]
 * Output: 0
 *
 * Note:
 *
 *
 * This is a follow up problem to Find Minimum in Rotated Sorted Array.
 * Would allow duplicates affect the run-time complexity? How and why?
 *
 *
 */

// @lc code=start
func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		switch mv := nums[mid]; {
		case mv < nums[start]:
			// mid 前面还有更小的数，或者 mid 就是最小值
			// 结果在 [start:mid]
			end = mid
		case mv > nums[end]:
			// mid 之后还有更大的数，或者 mid 本身是最大的数
			// 那么结果在 [mid:end]
			start = mid + 1
		case mv == nums[end]:
			end--
		default:
			end = start
		}
	}
	return nums[start]
}

// @lc code=end

