/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (37.55%)
 * Likes:    4599
 * Dislikes: 425
 * Total Accepted:    677.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
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
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
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
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 *
 *
*/

/*
 getMin() 要能在 constant time 返回最小值，并且是每次修改 stack 后都可以。需要一个 slice 存储每次操作后的min值。
 另外需要一个 slice 按照 stack 模式存储元素。
*/
// @lc code=start
type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	min := this.GetMin()
	if x < min {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min) // 存每次操作后：栈内的最小值。相当于一个快照。
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	l := len(this.stack)
	this.stack = this.stack[:l-1]
	this.min = this.min[:l-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 1 << 31 // 最小 int
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

