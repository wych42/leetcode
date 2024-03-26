/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (63.33%)
 * Likes:    5079
 * Dislikes: 0
 * Total Accepted:    908.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it can trap after raining.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * Explanation: The above elevation map (black section) is represented by array
 * [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue
 * section) are being trapped.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [4,2,0,3,2,5]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 1. find out how many units of water at specific position
 for each position, the unit of water can be trapped is determined by the max of left and max of right:
 e.g.:  current is 2, maxLeft is 3, maxRight is 5, 1 unit water can be trapped(min(3,5)-2)
 2. sum all the units

 *
*/

package main

// @lc code=start

func trap(height []int) int {

	result := 0

	l := 0
	r := len(height) - 1
	maxLeft := height[l]
	maxRight := height[r]

	// skip first and last, as they can not hold water
	for l < r {
		if maxLeft < maxRight {
			l++
			result += max(0, min(maxLeft, maxRight)-height[l])
			maxLeft = max(maxLeft, height[l])
		} else {
			r--
			result += max(0, min(maxLeft, maxRight)-height[r])
			maxRight = max(maxRight, height[r])
		}
	}
	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
// }

// @lc code=end
