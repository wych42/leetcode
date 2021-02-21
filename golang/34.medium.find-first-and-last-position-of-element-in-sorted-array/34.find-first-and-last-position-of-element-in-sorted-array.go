/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (33.76%)
 * Likes:    4897
 * Dislikes: 190
 * Total Accepted:    641.6K
 * Total Submissions: 1.7M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * If target is not found in the array, return [-1, -1].
 *
 * Follow up: Could you write an algorithm with O(log n) runtime complexity?
 *
 *
 * Example 1:
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * Example 2:
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * Example 3:
 * Input: nums = [], target = 0
 * Output: [-1,-1]
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums is a non-decreasing array.
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
/*
用两次二分查找，一次从左往右，一次从右往左。
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := binarySearchLeft(nums, target)
	right := binarySearchRight(nums, target)
	return []int{left, right}
}

func binarySearchLeft(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	// 查找区间是 [start,stop], 中止条件使用 start<end, 即中止时 start>=end
	for start < end {
		mid := start + (end-start)/2
		switch v := nums[mid]; {
		case v > target: // 结果落在 [start,mid-1]
			end = mid - 1
		case v < target: // 结果落在 [mid+1,end]
			start = mid + 1
		case v==target:
			end = mid // 结果可能就是 mid，也可能在 mid 前面
		}
	}
	if nums[start] != target { // 还差一个元素没有检查
		return -1
	}
	return start
}

func binarySearchRight(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	// 搜索区间是 [start,end], 在 start + 1 = end 的时候停止, 最终会留下两个元素未检查
	for start+1 < end {
		mid := start + (end-start)/2
		switch v := nums[mid]; {
		case v < target: // 结果在 mid 之后，所以 [mid+1,end]
			start = mid + 1
		case v > target: // 结果在 [start,mid-1] 中
			end = mid-1
		case v == target: // 结果可能是 mid，也能在 mid 之后，所以 [mid,end]
			start = mid
		}
	}
	if nums[end] == target {
		return end
	}
	if nums[start] == target {
		return start
	}
	return -1
}

// @lc code=end

