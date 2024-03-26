/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.cn/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (60.07%)
 * Likes:    4909
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * You are given an integer array height of length n. There are n vertical
 * lines drawn such that the two endpoints of the i^th line are (i, 0) and (i,
 * height[i]).
 *
 * Find two lines that together with the x-axis form a container, such that the
 * container contains the most water.
 *
 * Return the maximum amount of water a container can store.
 *
 * Notice that you may not slant the container.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array
 * [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the
 * container can contain is 49.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [1,1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4

 *
 *
 */

package main

// @lc code=start

func maxArea(height []int) int {
	// at start, the max area is start-end
	// iterate from start, end
	// if height is large than previous start or end
	// replace the result
	result := 0
	l, r := 0, len(height)-1
	for l < r {
		hl := height[l]
		hr := height[r]
		result = max(result, (r-l)*min(hl, hr))
		if hl < hr { // left is shorter, move left
			l++
		} else { // move right
			r--
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
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// }

// @lc code=end
