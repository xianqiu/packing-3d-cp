package packing_3d_cp

func compare(tree *SearchTree, i, j int, result *bool) {
	ins := tree.ins
	item := ins.GetItem(j)
	for r := new(Rotation).Init(item); r.NotEnd(); r.Next() {
		for a := Relation(1); a <= 6; a++ {
			//println(">> old tree ...")
			//tree.PrintTree()
			newTree := tree.Copy()
			newTree.AddArc(i, j, a)
			if newTree.IsFeasible() {
				// DEBUG INFO
				//println("feasible: (i, j, r, a) =", i, j, r.di, a)
				//newTree.PrintTree()
				//newTree.PrintItems()
				// ---
				if j == i+1 && j == (len(ins.items)-1) {
					*result = true
					//newTree.PrintItems()
					return
				}
				var p, q int
				if j == i+1 {
					p = 0
					q = j + 1
				} else {
					p = i + 1
					q = j
				}
				compare(newTree, p, q, result)
				if *result {
					return
				}
			}
			//println("infeasible: (i, j, r, a) =", i, j, r.di, a)
		}
	}
}

type Solver struct {
	ins *Instance
}

func (c *Solver) New(ins *Instance) *Solver {
	c.ins = ins
	return c
}

func (c *Solver) Solve() bool {
	// TODO: Uncomment.
	//sort.Sort(c.ins.items)
	res := false
	t := new(SearchTree).Init(c.ins)
	item0 := c.ins.GetItem(0)
	for r := new(Rotation).Init(item0); r.NotEnd(); r.Next() {
		compare(t, 0, 1, &res)
		if res {
			return true
		}
	}
	return false
}
