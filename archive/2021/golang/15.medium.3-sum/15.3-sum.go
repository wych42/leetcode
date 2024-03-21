/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (27.99%)
 * Likes:    10046
 * Dislikes: 1035
 * Total Accepted:    1.3M
 * Total Submissions: 4.4M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an integer array nums, return all the triplets [nums[i], nums[j],
 * nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] +
 * nums[k] == 0.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 *
 * Example 1:
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Example 2:
 * Input: nums = []
 * Output: []
 * Example 3:
 * Input: nums = [0]
 * Output: []
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start
import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// 先排序 [-1,0,1,2]
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 所有数都小于0或者大于0，就不用处理了
	if nums[len(nums)-1] < 0 || nums[0] > 0 {
		return nil
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		// 计算 [i:] 的结果
		// 因为是有序的，nums[i] > 0, 意味着后面都大于0, 不用继续了
		if nums[i] > 0 {
			return res
		}
		// 判断下重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 用左右指针, 往中间靠 [i+1:]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 { // 右边太大
				r--
			} else { // 左边太小
				l++
			}
		}
	}
	return res
}

// @lc code=end

