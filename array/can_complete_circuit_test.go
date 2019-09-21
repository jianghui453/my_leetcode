package array

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
    var gas, cost []int
    var h, r int

    gas  = []int{1,2,3,4,5}
    cost = []int{3,4,5,1,2}
    h = 3
    r = canCompleteCircuit(gas, cost)
    t.Logf("h=%d r=%d", h, r)

    gas  = []int{2,3,4}
    cost = []int{3,4,3}
    h = -1
    r = canCompleteCircuit(gas, cost)
    t.Logf("h=%d r=%d", h, r)
}
