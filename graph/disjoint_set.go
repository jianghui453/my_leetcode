package graph

type DisjointSet struct {
	ranks   []int
	subsets []int
}

func (d *DisjointSet) Constructor(cnt int) {
	d.ranks = make([]int, cnt)
	d.subsets = make([]int, cnt)
	for i := 0; i < cnt; i++ {
		d.ranks[i] = 1
		d.subsets[i] = i
	}
}

func (d *DisjointSet) Find(i int) int {
	if d.subsets[i] != i {
		d.subsets[i] = d.Find(d.subsets[i])
	}
	return d.subsets[i]
}

func (d *DisjointSet) Union(i, j int) {
	rooti := d.Find(i)
	rootj := d.Find(j)

	if rooti == rootj {
		return
	}

	if d.ranks[rooti] <= d.ranks[rootj] {
		d.subsets[rooti] = rootj
		if d.ranks[rooti] == d.ranks[rootj] {
			d.ranks[rootj]++
		}
	} else if d.ranks[rooti] > d.ranks[rootj] {
		d.subsets[rootj] = rooti
	}
}
