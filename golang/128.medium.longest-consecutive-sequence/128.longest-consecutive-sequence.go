/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.cn/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (52.07%)
 * Likes:    2027
 * Dislikes: 0
 * Total Accepted:    588.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers nums, return the length of the longest
 * consecutive elements sequence.
 *
 * You must write an algorithm that runs in O(n) time.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *

 * sort and check: o(nlogn) + o(n)

 * use a hashmap store each element.
 * iterate each element, check its greater/smaller siblings exists in the hashmap, mark checked.
 *
 */

package main

// @lc code=start
func longestConsecutive(nums []int) int {
	elemMap := make(map[int][]int, len(nums))
	for _, num := range nums {
		if _, ok := elemMap[num]; ok {
			continue
		}
		elemMap[num] = []int{num}
	}
	visited := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := visited[num]; ok {
			continue
		}
		visited[num] = struct{}{}
		min := num
		for {
			min--
			if _, ok := elemMap[min]; ok {
				visited[min] = struct{}{}
				elemMap[num] = append(elemMap[num], min)
			} else {
				break
			}
		}
		max := num
		for {
			max++
			if _, ok := elemMap[max]; ok {
				visited[max] = struct{}{}
				elemMap[num] = append(elemMap[num], max)
			} else {
				break
			}
		}
	}
	result := 0
	for _, v := range elemMap {
		if len(v) > result {
			result = len(v)
		}
	}
	return result
}

// func main() {
// 	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
// }

// @lc code=end
