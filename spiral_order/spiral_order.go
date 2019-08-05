package spiral_order

import "fmt"

func spiralOrder(matrix [][]int) []int {
    var ret []int
    m := len(matrix)
    if m == 0 {
        return ret
    }
    n := len(matrix[0])
    if n == 0 {
        return ret
    }
    for m > 0 && n > 0 {
        fmt.Printf("m=%d n=%d matrix=%v\n", m, n, matrix)
        if m == 1 {
            return append(ret, matrix[0]...)
        }
        if n == 1 {
            for i := 0; i < m; i ++ {
                ret = append(ret, matrix[i][0])
            }
            return ret
        }
        ret = append(ret, matrix[0]...)
        for i := 1; i < m; i ++ {
            ret = append(ret, matrix[i][n-1])
        }
        for i := n-2; i >=0; i -- {
            ret = append(ret, matrix[m-1][i])
        }
        for i := m-2; i > 0; i -- {
            ret = append(ret, matrix[i][0])
        }
        if m == 2 || n == 2 {
            return ret
        }
        matrix = matrix[1:m-1]
        for i := 0; i < m-2; i ++ {
            fmt.Printf("matrix=%v i=%d\n", matrix, i)
            matrix[i] = matrix[i][1:n-1]
        }
        m -= 2
        n -= 2
    }
    return ret
}