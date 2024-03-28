/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 *
 * https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (45.59%)
 * Likes:    2697
 * Dislikes: 0
 * Total Accepted:    403K
 * Total Submissions: 884K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * Given an array of integers heights representing the histogram's bar height
 * where the width of each bar is 1, return the area of the largest rectangle
 * in the histogram.
 *
 *
 * Example 1:
 *
 *
 * Input: heights = [2,1,5,6,2,3]
 * Output: 10
 * Explanation: The above is a histogram where width of each bar is 1.
 * The largest rectangle is shown in the red area, which has an area = 10
 * units.
 *
 *
 * Example 2:
 *
 *
 * Input: heights = [2,4]
 * Output: 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= heights.length <= 10^5
 * 0 <= heights[i] <= 10^4
 *

 for each fixed-width histogram, the max rectangle is determined by the lowest bar.

 for each height, if its next height is bigger, means the rectangle can extends
 until next height is shorter
 *
*/

package main

// @lc code=start
type wall struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	// track the max result
	result := 0
	// tack the continus increase wall
	stack := make([]wall, 0)

	for i, height := range heights {
		// for each height, if stack is not empty

		// track the position current height starts continus increase
		start := i
		for len(stack) > 0 && stack[len(stack)-1].height > height {
			// get top of stack
			top := stack[len(stack)-1]
			// pop the element
			stack = stack[:len(stack)-1]

			// // if stack's top height <= current height, means the rectangle can continue expand
			// if top.height <= height {
			// 	// break the stack loop, continue process next height
			// 	break
			// }

			// if stack's top height > current height, the rectangle can not continue expand
			// backtrack the stack, find wall[j] <= current height[i]
			// hence, current height, continues from j to i, we want to add current height at position j
			start = top.index
			// and update result during the backtrack
			result = max(result, top.height*(i-top.index))
		}
		stack = append(stack, wall{index: start, height: height})
	}
	// for all the walls left in stack, all heights are in increase order
	// for each wall, the max rectangle starts here are height * (len(heights) - current index)
	// fmt.Println("walls", stack, "max", result)
	for _, wall := range stack {
		result = max(result, (len(heights)-wall.index)*wall.height)
	}
	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
}
