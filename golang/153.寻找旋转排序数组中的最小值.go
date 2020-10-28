/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package golang

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		switch v := nums[mid]; {
		case v <= nums[end]:
			end = mid
		default:
			start = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}

// @lc code=end
