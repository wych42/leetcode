package golang

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		switch v := nums[mid]; {
		// result in [start:mid]
		case v < nums[start]:
			end = mid
		// result in [mid:end]
		case v > nums[end]:
			start = mid + 1 // skip mid
		case v == nums[end]:
			end = end - 1
		default:
			end = start
		}
	}
	return nums[start]

}

// @lc code=end
