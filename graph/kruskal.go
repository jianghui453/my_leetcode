package graph

type Edge struct {
	src    int
	dst    int
	weight int
}

type Graph struct {
	V, E  int
	Edges []Edge
}

func kruskal(g Graph) Graph {
	sortEdges(g.Edges)

	var retg Graph
	retg.V = g.V

	disjointset := new(DisjointSet)
	disjointset.Constructor(g.V)

	for i := 0; i < g.E; i++ {
		roots := disjointset.Find(g.Edges[i].src)
		rootd := disjointset.Find(g.Edges[i].dst)

		if roots == rootd {
			continue
		}

		disjointset.Union(roots, rootd)
		retg.Edges = append(retg.Edges, g.Edges[i])
		retg.E++
	}

	return retg
}

func sortEdges(edges []Edge) {
	l := len(edges)
	if l <= 1 {
		return
	}
	newedges := make([]Edge, l)

	left, right := 0, l-1
	for i := 1; i < l; i++ {
		if edges[i].weight > edges[0].weight {
			newedges[right] = edges[i]
			right--
		} else if edges[i].weight < edges[0].weight {
			newedges[left] = edges[i]
			left++
		}
	}
	for i := left; i <= right; i++ {
		newedges[i] = edges[0]
	}

	copy(edges, newedges)
	sortEdges(edges[:left])
	if right < l-1 {
		sortEdges(edges[right+1:])
	}
}

// func kruskal(g Graph) Graph {
// 	edges := g.Edges
// 	sortEdges(edges)

// 	eLen := len(edges)
// 	parents := make([]int, g.V)
// 	for i := 0; i < g.V; i++ {
// 		parents[i] = i
// 	}

// 	rank := make([]int, g.V)

// 	mst := Graph{}
// 	for i := 0; i < eLen; i++ {
// 		src := edges[i].src
// 		dst := edges[i].dst
// 		rootSrc := find(parents, src)
// 		rootDst := find(parents, dst)
// 		if rootSrc == rootDst {
// 			continue
// 		}

// 		union(parents, rank, rootSrc, rootDst)

// 		mst.Edges = append(mst.Edges, edges[i])
// 		mst.E++
// 	}

// 	eLen = len(mst.Edges)
// 	sVertex := make([]int, g.V)
// 	for i := 0; i < eLen; i++ {
// 		if sVertex[mst.Edges[i].src] == 0 {
// 			mst.V++
// 			sVertex[mst.Edges[i].src]++
// 		}
// 		if sVertex[mst.Edges[i].dst] == 0 {
// 			mst.V++
// 			sVertex[mst.Edges[i].dst]++
// 		}
// 	}
// 	return mst
// }

// func sortEdges(edges []Edge) {
// 	eLen := len(edges)
// 	if eLen < 2 {
// 		return
// 	}
// 	copyEdges := make([]Edge, eLen)
// 	copy(copyEdges, edges)
// 	cmpWeight := edges[0].weight
// 	equalEdges := []Edge{edges[0]}
// 	left := 0
// 	right := eLen - 1
// 	for i := 1; i < eLen; i++ {
// 		if copyEdges[i].weight > cmpWeight {
// 			copyEdges[i], edges[right] = edges[right], copyEdges[i]
// 			right--
// 		} else if copyEdges[i].weight < cmpWeight {
// 			copyEdges[i], edges[left] = edges[left], copyEdges[i]
// 			left++
// 		} else {
// 			equalEdges = append(equalEdges, copyEdges[i])
// 		}
// 	}
// 	for i := left; i <= right; i++ {
// 		edges[i] = equalEdges[i-left]
// 	}
// 	sortEdges(edges[:left])
// 	if right < eLen-1 {
// 		sortEdges(edges[right+1:])
// 	}
// }

// func find(parents []int, x int) int {
// 	if parents[x] != x {
// 		parents[x] = find(parents, parents[x])
// 	}
// 	return parents[x]
// }

// func union(parents, rank []int, x int, y int) {
// 	rootX := find(parents, x)
// 	rootY := find(parents, y)
// 	if rootX == rootY {
// 		return
// 	}
// 	if rank[rootX] > rank[rootY] {
// 		parents[rootY] = rootX
// 	} else {
// 		parents[rootX] = rootY
// 		if rank[rootX] == rank[rootY] {
// 			rank[rootY]++
// 		}
// 	}
// }
