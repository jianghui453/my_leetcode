package stack

import (
	"testing"
)

func Test_MinStack(t *testing.T) {
	var minStack *MinStack
	var hope, ret int

	c := Constructor()
	minStack = &c
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	hope = -3
	ret = minStack.GetMin()
	t.Logf("GetMin: %t hope=%d ret=%d", ret == hope, hope, ret)

	minStack.Pop()
	hope = 0
	ret = minStack.Top()
	t.Logf("Top: %t hope=%d ret=%d", ret == hope, hope, ret)

	hope = -2
	ret = minStack.GetMin()
	t.Logf("GetMin: %t hope=%d ret=%d", ret == hope, hope, ret)
}