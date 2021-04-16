// 给定一个按顺序连接的多边形的顶点，判断该多边形是否为凸多边形。（凸多边形的定义）

// 注：

// 顶点个数至少为 3 个且不超过 10,000。
// 坐标范围为 -10,000 到 10,000。
// 你可以假定给定的点形成的多边形均为简单多边形（简单多边形的定义）。换句话说，保证每个顶点处恰好是两条边的汇合点，并且这些边 互不相交 。
//  

// 示例 1：

// [[0,0],[0,1],[1,1],[1,0]]

// 输出： True

// 解释：
// 示例 2：

// [[0,0],[0,10],[10,10],[10,0],[5,5]]

// 输出： False

// 解释：

package math_algorithm

import (
	// "fmt"
)

func isConvex(points [][]int) bool {
	n := len(points)
	pre := 0
	points = append(points, points[0])
	points = append(points, points[1])
	
	for i := 0; i < n; i++ {
		dx1 := points[i+1][0] - points[i][0]
		dy1 := points[i+1][1] - points[i][1]
		dx2 := points[i+2][0] - points[i][0]
		dy2 := points[i+2][1] - points[i][1]

		cur := dx1 * dy2 - dx2 * dy1

		if cur != 0 {
			if cur * pre < 0 {
				return false
			}
			pre = cur
		}
	}
	
	return true
}