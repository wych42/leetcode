/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (37.80%)
 * Likes:    6778
 * Dislikes: 0
 * Total Accepted:    1.7M
 * Total Submissions: 4.6M
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
 *
 *
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Explanation:
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 * The distinct triplets are [-1,0,1] and [-1,-1,2].
 * Notice that the order of the output and the order of the triplets does not
 * matter.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,1]
 * Output: []
 * Explanation: The only possible triplet does not sum up to 0.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0,0,0]
 * Output: [[0,0,0]]
 * Explanation: The only possible triplet sums up to 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *

 we need two numbers sum to zero, there're cases:
 - 0+0+0
 - -x + -y + z
 - -x + 0 + y
 - -x + y + z

 solution:
 1. sort nums in non-desc order
 2. if all num < 0 or > 0, no result
 2. iterate nums, nums[i]
    - resolv two-sum problem with multiple result in nums[i+1:]
	- resolve two-sum of sorted array input by two-pointer
 *
*/

package main

// @lc code=start
import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// don't check nums length

	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}

	for i := 0; i < len(nums)-1; i++ {
		// no more result if nums[i] > 0
		if nums[i] > 0 {
			return result
		}

		// skip duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// two pointer
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// skip duplicate [-2, 0, 0, 2, 2]
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// move to next
				left++
				right--
				continue
			}
			if sum > 0 {
				right--
				continue
			}
			if sum < 0 {
				left++
				continue
			}
		}
	}
	return result

}

// @lc code=end
