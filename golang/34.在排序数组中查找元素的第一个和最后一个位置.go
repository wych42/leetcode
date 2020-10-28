/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package golang

// @lc code=start
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
	for start+1 < end {
		mid := (start + end) / 2
		switch v := nums[mid]; {
		case v >= target:
			end = mid
		case v < target:
			start = mid
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
func binarySearchRight(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		switch v := nums[mid]; {
		case v <= target:
			start = mid
		case v > target:
			end = mid
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
