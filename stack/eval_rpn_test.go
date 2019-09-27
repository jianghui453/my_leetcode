package stack

import "testing"

func TestEvalRPN(t *testing.T) {
    var tokens []string
    var h, r int

    tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
    h = 22
    r = evalRPN(tokens)
    t.Logf("tokens=%v \n h=%d \n r=%d", tokens, h, r)

    tokens = []string{"4", "13", "5", "/", "+"}
    h = 6
    r = evalRPN(tokens)
    t.Logf("tokens=%v \n h=%d \n r=%d", tokens, h, r)

    tokens = []string{"2", "1", "+", "3", "*"}
    h = 9
    r = evalRPN(tokens)
    t.Logf("tokens=%v \n h=%d \n r=%d", tokens, h, r)

    tokens = []string{"2", "1", "+"}
    h = 3
    r = evalRPN(tokens)
    t.Logf("tokens=%v \n h=%d \n r=%d", tokens, h, r)

    tokens = []string{"2", "1", "-"}
    h = 1
    r = evalRPN(tokens)
    t.Logf("tokens=%v \n h=%d \n r=%d", tokens, h, r)
}
