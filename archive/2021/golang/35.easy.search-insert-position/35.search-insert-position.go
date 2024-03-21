/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.97%)
 * Likes:    3143
 * Dislikes: 285
 * Total Accepted:    762.4K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array of distinct integers and a target value, return the
 * index if the target is found. If not, return the index where it would be if
 * it were inserted in order.
 *
 *
 * Example 1:
 * Input: nums = [1,3,5,6], target = 5
 * Output: 2
 * Example 2:
 * Input: nums = [1,3,5,6], target = 2
 * Output: 1
 * Example 3:
 * Input: nums = [1,3,5,6], target = 7
 * Output: 4
 * Example 4:
 * Input: nums = [1,3,5,6], target = 0
 * Output: 0
 * Example 5:
 * Input: nums = [1], target = 0
 * Output: 0
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums contains distinct values sorted in ascending order.
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
/*
用二分查找，找到 start, end，使 nums[start]和nums[end] 都最接近 target.
再比较 nums[start], nums[end] 与 target 的大小，决定插入位置
*/
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
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
		return end + 1
	}
}

// @lc code=end

