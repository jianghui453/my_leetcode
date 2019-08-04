// 迪杰斯特拉算法
// 获取大小为V的图的最小生成树
// 把起始点的距离值设为0
// （1）找到未被标记的顶点中离已被标记的顶点u最近的一个顶点v
// （2）更新这个顶点的距离值（u的距离值加u到v的距离）
// （3）如果还有未被标记的顶点回到（1）

package dijkstra

func dijkstra(V int, graph [V][V]int) []int {
    sptTree := [V]bool{}
    sptTree[0] = true
    spt := []int{0}

    for i := 1; i < V; i ++ {
        for i := 0; i < V; i ++ {
            if sptTree[i] == true {
                continue
            }
        }
    }
}