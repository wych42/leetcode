/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	min []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: make([]int, 0),
		stack: make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	min := this.GetMin()
	if x < min {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min)
	}
}

func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	return this.min[len(this.min) - 1]
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

