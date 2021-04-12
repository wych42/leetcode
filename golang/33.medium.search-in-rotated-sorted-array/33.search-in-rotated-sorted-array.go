/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (35.84%)
 * Likes:    7391
 * Dislikes: 649
 * Total Accepted:    964.5K
 * Total Submissions: 2.7M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * There is an integer array nums sorted in ascending order (with distinct
 * values).
 *
 * Prior to being passed to your function, nums is rotated at an unknown pivot
 * index k (0 <= k < nums.length) such that the resulting array is [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
 * For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become
 * [4,5,6,7,0,1,2].
 *
 * Given the array nums after the rotation and an integer target, return the
 * index of target if it is in nums, or -1 if it is not in nums.
 *
 *
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2:
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * All values of nums are unique.
 * nums is guaranteed to be rotated at some pivot.
 * -10^4 <= target <= 10^4
 *
 *
 *
 * Follow up: Can you achieve this in O(log n) time complexity?
 */

// @lc code=start
func search(nums []int, target int) int {
	/*
		4,5,6,7,0,1,2;
		还是二分的思路: left, mid, right
		但是列表有两块： 大部, 小部
		mid > target:
		  - 如果 mid 在小部(mid<right), target 不可能在左边, 即 target 一定在右边, target <= right
		  - 如果 mid 在大部(mid>right)，target 在右边; 即 target <= right
		  - 如果 mid 在大部(mid>right)，target 也在左边，即 target > right
		mid < target:
		  - 如果 mid 在大部(mid>right), 那么 target 一定在 mid 和 right 之间
		  - 如果 mid 在小部(mid<right), 那么 target 可能在右边，也可能在左边
		  	- 如果 target 在小部，那么 target <= right
			- 如果 target 在大部，那么 target > right
	*/
	left := 0
	right := len(nums) - 1
	for left+1 < right {
		mid := (right-left)/2 + left
		midVal := nums[mid]
		//fmt.Println(left, right, midVal)
		if midVal > target {
			if midVal < nums[right] { // mid 在小的那部分里, target 也在小部
				right = mid
			} else { // mid 在大部
				if target <= nums[right] { // target 在小部
					left = mid
				} else { // target 在大部
					right = mid
				}
			}
		} else if midVal < target {
			if midVal >= nums[left] { // mid 在大部, target 也在大部
				left = mid
			} else { // 此时 mid 在小部
				if target > nums[right] { // target 在大部
					right = mid
				} else { // target 在小部
					left = mid
				}
			}
		} else {
			return mid
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

// @lc code=end

