/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		mid := (start + end) / 2
		switch v := nums[mid]; {
		case v == target:
			return mid
		case v < target:
			start = mid
		case v > target:
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
// @lc code=end

