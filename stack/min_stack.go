// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) -- 将元素 x 推入栈中。
// pop() -- 删除栈顶的元素。
// top() -- 获取栈顶元素。
// getMin() -- 检索栈中的最小元素。
// 示例:

// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

package stack

type MinStack struct {
	mainStack []int
	trackStack []int
	length int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	var m MinStack
	m.mainStack = make([]int, 0)
	m.trackStack = make([]int, 0)
	return m
}


func (this *MinStack) Push(x int) {
	this.mainStack = append(this.mainStack, x)
	if this.length > 0 && x > this.trackStack[this.length-1] {
		x = this.trackStack[this.length-1]
	}
	this.trackStack = append(this.trackStack, x)
	this.length++
}


func (this *MinStack) Pop()  {
    if this.length > 0 {
		this.mainStack = this.mainStack[: this.length-1]
		this.trackStack = this.trackStack[: this.length-1]
		this.length--
	}
}


func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	} else {
		return this.mainStack[this.length-1]
	}
}


func (this *MinStack) GetMin() int {
    if this.length == 0 {
		return 0
	}
	return this.trackStack[this.length-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */