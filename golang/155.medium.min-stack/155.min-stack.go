/*
* @lc app=leetcode.cn id=155 lang=golang
*
* [155] Min Stack
*
* https://leetcode.cn/problems/min-stack/description/
*
  - algorithms
  - Medium (59.62%)
  - Likes:    1741
  - Dislikes: 0
  - Total Accepted:    563.7K
  - Total Submissions: 945.4K
  - Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
    '[[],[-2],[0],[-3],[],[],[],[]]'

*
* Design a stack that supports push, pop, top, and retrieving the minimum
* element in constant time.
*
* Implement the MinStack class:
*
*
* MinStack() initializes the stack object.
* void push(int val) pushes the element val onto the stack.
* void pop() removes the element on the top of the stack.
* int top() gets the top element of the stack.
* int getMin() retrieves the minimum element in the stack.
*
*
* You must implement a solution with O(1) time complexity for each
* function.
*
*
* Example 1:
*
*
* Input
* ["MinStack","push","push","push","getMin","pop","top","getMin"]
* [[],[-2],[0],[-3],[],[],[],[]]
*
* Output
* [null,null,null,null,-3,null,0,-2]
*
* Explanation
* MinStack minStack = new MinStack();
* minStack.push(-2); [-2,]
* minStack.push(0); [-2, 0]
* minStack.push(-3); [-2, 0, -3]
* minStack.getMin(); // return -3
* minStack.pop();
* minStack.top();    // return 0
* minStack.getMin(); // return -2
*
*
*
* Constraints:
*
*
* -2^31 <= val <= 2^31 - 1
* Methods pop, top and getMin operations will always be called on non-empty
* stacks.
* At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
*
*
*/
package main

// @lc code=start
type MinStack struct {
	// use slice track normal stack
	stack []int
	// use slice track current minimal
	min []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min) > 0 {
		this.min = append(this.min, min(this.min[len(this.min)-1], val))
	} else {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	// operations will be called on non-empty stack
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	// operations will be called on non-empty stack
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	// operations will be called on non-empty stack
	return this.min[len(this.min)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
