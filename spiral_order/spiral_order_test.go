package spiral_order

import "testing"

func TestSpiralOrder(t *testing.T) {
    var matrix [][]int
    var hope, ret []int

    matrix = [][]int{
        { 1, 2, 3 },
        { 4, 5, 6 },
        { 7, 8, 9 },
    }
    hope = []int{1,2,3,6,9,8,7,4,5}
    ret = spiralOrder(matrix)
    t.Logf("hope=%v \nret=%v", hope, ret)

    matrix = [][]int{
        { 1, 2, 3, 4 },
        { 5, 6, 7, 8 },
        { 9,10,11,12 },
    }
    hope = []int{ 1,2,3,4,8,12,11,10,9,5,6,7 }
    ret = spiralOrder(matrix)
    t.Logf("hope=%v \nret=%v", hope, ret)
}