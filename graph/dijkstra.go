// 迪杰斯特拉算法
// 获取大小为V的图的最小生成树
// 把起始点的距离值设为0
// （1）找到未被标记的顶点中离已被标记的顶点u最近的一个顶点v
// （2）更新这个顶点的距离值（u的距离值加u到v的距离）
// （3）如果还有未被标记的顶点回到（1）

package dijkstra

import (
	"fmt"
	// "math"
)

func dijkstra(V int, graph [][]int) []int {
	spt := make([]int, V)

	for flag := true; flag; {
		var (
			idx int
			min int
		)

		flag = false
		for i := 0; i < V; i++ {
			if i != 0 && spt[i] == 0 {
				flag = true
				continue
			}

			for j := 1; j < V; j++ {
				if spt[j] == 0 && graph[i][j] > 0 && (min == 0 || spt[i]+graph[i][j] < min) {
					min = spt[i] + graph[i][j]
					idx = j
				}
			}
		}

		spt[idx] = min
	}

	return spt
}

// func dijkstra(V int, graph [][]int) []int {
// 	sptMark := make([]bool, V) // 记录点是否已经标记了
// 	spt := make([]int, V)      // 记录点的最小距离值
// 	for i := 0; i < V; i++ {
// 		spt[i] = -1
// 	}
// 	// 原点的距离值标记为0
// 	spt[0] = 0
// 	for i := 1; i < V; i++ {
// 		// 查找能达到的，距离值最小的点
// 		minIdx := -1
// 		minVal := math.MaxInt64
// 		for i := 0; i < V; i++ {
// 			if sptMark[i] == false && spt[i] != -1 && spt[i] < minVal {
// 				minIdx = i
// 				minVal = spt[i]
// 			}
// 		}
// 		// 把这个点能达到的点标上距离值
// 		sptMark[minIdx] = true
// 		for i := 0; i < V; i++ {
// 			if sptMark[i] == false && graph[minIdx][i] > 0 &&
// 				(spt[i] == -1 || spt[minIdx]+graph[minIdx][i] < spt[i]) {
// 				spt[i] = spt[minIdx] + graph[minIdx][i]
// 			}
// 		}
// 	}
// 	return spt
// }
