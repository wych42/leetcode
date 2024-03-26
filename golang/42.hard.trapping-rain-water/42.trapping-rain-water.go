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

 consider each height is a wall.
 for two walls: Left and Right
 only when Right >= Left, water can be trapped between:
 - if no walls between them, amount is Min(Right.Height,Left.Height)*(Right.Index-Left.Index)
 - if there walls between them, substract wall heights from amount above

 for each wall in the list, we need find previous most suitable wall.
 *
*/

package main

// @lc code=start

type Wall struct {
	index  int
	height int
}

func trap(height []int) int {

	result := 0
	// use a stack to track
	stack := make([]Wall, 0)
	for i, h := range height {
		wall := Wall{index: i, height: h}
		// push current wall to stack if stack is empty(means first wall)
		if len(stack) == 0 {
			stack = append(stack, wall)
			continue
		}
		// or current wall is shorter than stack top wall()
		if h < stack[len(stack)-1].height {
			stack = append(stack, wall)
			continue
		}
		// if current wall is *not shorter* than stack top wall
		// there's room between current wall[i], and one of previous walls
		// The Right wall is i with height h

		for len(stack) > 0 && stack[len(stack)-1].height <= h {
			// fmt.Println("stack", stack, "currentWall", wall)
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// no walls before, can not trap
			if len(stack) == 0 {
				break
			}
			prev := stack[len(stack)-1]
			width := i - prev.index - 1

			result += (min(prev.height, h) - top.height) * width
			// fmt.Println("result", result)
		}
		stack = append(stack, wall)

	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	trap([]int{2, 0, 1, 0, 2})
}

// @lc code=end
