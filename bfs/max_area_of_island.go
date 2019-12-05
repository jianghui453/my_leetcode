// 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

// 示例 1:

// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,1,1,0,1,0,0,0,0,0,0,0,0],
//  [0,1,0,0,1,1,0,0,1,0,1,0,0],
//  [0,1,0,0,1,1,0,0,1,1,1,0,0],
//  [0,0,0,0,0,0,0,0,0,0,1,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

// 示例 2:

// [[0,0,0,0,0,0,0,0]]
// 对于上面这个给定的矩阵, 返回 0。

// 注意: 给定的矩阵grid 的长度和宽度都不超过 50。

package dfs

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt := 0
				eles := make([][2]int, 0)
				eles = append(eles, [2]int{i, j})
				for len(eles) > 0 {
					cnt += len(eles)
					newEles := make([][2]int, 0)
					for k := range eles {
						x, y := eles[k][0], eles[k][1]

						if x > 0 && grid[x-1][y] == '1' {
							newEles = append(newEles, [2]int{x-1, y})
							grid[x-1][y] = '2'
						}

						if y > 0 && grid[x][y-1] == '1' {
							newEles = append(newEles, [2]int{x, y-1})
							grid[x][y-1] = '2'
						}

						if x < m-1 && grid[x+1][y] == '1' {
							newEles = append(newEles, [2]int{x+1, y})
							grid[x+1][y] = '2'
						}

						if y < n-1 && grid[x][y+1] == '1' {
							newEles = append(newEles, [2]int{x, y+1})
							grid[xy+1] = '2'
						}
					}

					eles = newEles
				}
				ret = max(ret, cnt)
			}
		}
	}
	
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
