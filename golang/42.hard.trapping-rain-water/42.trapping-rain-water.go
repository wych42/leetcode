/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (51.10%)
 * Likes:    10654
 * Dislikes: 159
 * Total Accepted:    713.1K
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
 * 0 <= n <= 3 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */

// @lc code=start
type wall struct {
	idx    int
	height int
}

func trap(heights []int) int {
	/*
	 用栈来计算.
	 1. 当前高度 i < top, 就把 i 入站
	 2. i >= top 时，说明左侧和 i 之间有空出来位置，可以接雨水
	    一个 for 循环（只要栈不为空，或者 i >= top)
	    1. 出栈，此时地表在出栈元素的高度上
	 	2. 计算i与此时栈顶的宽度，相乘就是这一块的积水
	*/
	stack := make([]wall, 0)
	amount := 0
	for i, h := range heights {
		w := wall{idx: i, height: h}
		if len(stack) == 0 {
			stack = append(stack, w)
			continue
		}
		if h < stack[len(stack)-1].height {
			stack = append(stack, w)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1].height <= h { //栈顶 <= 当前高度，中间可以积水
			top := stack[len(stack)-1] //栈顶(出栈)
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break // 说明前面没有墙了，不能积水
			}

			// 栈顶前面一个
			prev := stack[len(stack)-1]
			width := i - prev.idx - 1
			amount += (min(prev.height, h) - top.height) * width
		}
		stack = append(stack, w)
	}
	return amount
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

