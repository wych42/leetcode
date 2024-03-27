/*
* @lc app=leetcode.cn id=739 lang=golang
*
* [739] Daily Temperatures
*
* https://leetcode.cn/problems/daily-temperatures/description/
*
* algorithms
* Medium (68.76%)
* Likes:    1732
* Dislikes: 0
* Total Accepted:    528.7K
* Total Submissions: 768.9K
* Testcase Example:  '[73,74,75,71,69,72,76,73]'
*
* Given an array of integers temperatures represents the daily temperatures,
* return an array answer such that answer[i] is the number of days you have to
* wait after the i^th day to get a warmer temperature. If there is no future
* day for which this is possible, keep answer[i] == 0 instead.
*
*
* Example 1:
* Input: temperatures = [73,74,75,71,69,72,76,73]
* Output: [1,1,4,2,1,1,0,0]
* Example 2:
* Input: temperatures = [30,40,50,60]
* Output: [1,1,1,0]
* Example 3:
* Input: temperatures = [30,60,90]
* Output: [1,1,0]
*
*
* Constraints:
*
*
* 1 <= temperatures.length <= 10^5
* 30 <= temperatures[i] <= 100
*

the result of each day i, is index - i, index is the first value > tmp[i]
- use a stack keep track of unsolved days
- iterate through each day: for each day in stack(top), check :
  - if current tmp > top.tmp, result[top.index] = current.index - top.index
  - else, append current to stack

*
*/
package main

// @lc code=start
type temperature struct {
	value int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]temperature, 0)
	for i, t := range temperatures {
		for len(stack) > 0 && t > stack[len(stack)-1].value {
			res[stack[len(stack)-1].index] = i - stack[len(stack)-1].index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, temperature{index: i, value: t})
	}

	return res
}

// @lc code=end
