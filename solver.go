package packing_3d_cp

func compare(tree *SearchTree, ins *Instance, i, j int, result *bool) {
	item := ins.GetItem(j)
	for r := new(Rotation).Init(item); r.NotEnd(); r.Next() {
		for a := Relation(1); a <= 6; a++ {
			println(">> old tree ...")
			tree.Print()
			newTree := tree.Copy()
			newTree.AddArc(ins, i, j, a)
			if newTree.IsFeasible(ins) {
				// DEBUG INFO
				println("i, j, r, di:", i, j, a, r.di)
				newTree.Print()
				for i := 0; i < len(newTree.nodes); i++ {
					xa, ya, za := newTree.GetXyzOfA(i)
					xb, yb, zb := newTree.GetXyzOfB(ins, i)
					println("item", i, "A =", xa, ya, za, "B =", xb, yb, zb, "Dim =",
						ins.GetItem(i).L, ins.GetItem(i).W, ins.GetItem(i).H)
				}
				// ---
				if j == i+1 && j == (len(ins.items)-1) {
					*result = true
					return
				}
				if j == i+1 {
					i = 0
					j += 1
				} else {
					i += 1
				}
				compare(newTree, ins, i, j, result)
				if *result {
					return
				}
			}
		}
	}
}

type Solver struct {
	ins *Instance
	s   *SearchTree
}

func (c *Solver) New(ins *Instance) *Solver {
	c.ins = ins
	return c
}

func (c *Solver) Solve() bool {
	// TODO: Uncomment.
	//sort.Sort(c.ins.items)
	res := false
	t := new(SearchTree).Init()
	item0 := c.ins.GetItem(0)
	for r := new(Rotation).Init(item0); r.NotEnd(); r.Next() {
		compare(t, c.ins, 0, 1, &res)
		if res {
			return true
		}
	}
	return false
}

func (c *Solver) GetLocation(id int) (float64, float64, float64) {
	x := c.s.rl.GetNode(id).location
	y := c.s.rw.GetNode(id).location
	z := c.s.rh.GetNode(id).location
	return x, y, z
}
