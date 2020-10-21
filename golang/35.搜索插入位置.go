/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
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
	switch t := target; {
	case nums[start] >= t:
		return start
	case nums[end] >= t:
		return end
	default:
		return end+1
	}
}
// @lc code=end

