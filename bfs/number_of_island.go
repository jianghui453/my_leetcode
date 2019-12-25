// 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

// 示例 1:

// 输入:
// 11110
// 11010
// 11000
// 00000

// 输出: 1
// 示例 2:

// 输入:
// 11000
// 11000
// 00100
// 00011

// 输出: 3

package bfs

func numIslands(grid [][]byte) int {
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
				ret++
				eles := make([][2]int, 0)
				eles = append(eles, [2]int{i, j})
				grid[i][j] = '2'
				for len(eles) > 0 {
					newEles := make([][2]int, 0)
					for k := range eles {
						x, y := eles[k][0], eles[k][1]
						if x > 0 && grid[x-1][y] == '1' {
							newEles = append(newEles, [2]int{x - 1, y})
							grid[x-1][y] = '2'
						}
						if x < m-1 && grid[x+1][y] == '1' {
							newEles = append(newEles, [2]int{x + 1, y})
							grid[x+1][y] = '2'
						}
						if y > 0 && grid[x][y-1] == '1' {
							newEles = append(newEles, [2]int{x, y - 1})
							grid[x][y-1] = '2'
						}
						if y < n-1 && grid[x][y+1] == '1' {
							newEles = append(newEles, [2]int{x, y + 1})
							grid[x][y+1] = '2'
						}
					}
					eles = newEles
				}
			}
		}
	}

	return ret
}
