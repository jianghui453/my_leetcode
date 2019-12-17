package stack

import (
	"testing"
)

func Test_myQueue(t *testing.T) {
	var q *MyQueue
	
	_q := Constructor()
	q = &_q
	q.Push(1)
	q.Push(2)
	t.Log("peek =", q.Peek(), "hope =", 1)
	t.Log("pop =", q.Pop(), "hope =", 1)
	t.Log("pop =", q.Pop(), "hope =", 2)
	
	t.Log("empty =", q.Empty(), "hope =", true)
}