package packing_3d_cp

func compare(tree *SearchTree, i, j int, resTree **SearchTree) bool {
	ins := tree.ins
	item := ins.GetItem(j)
	for r := new(Rotation).Init(item); r.NotEnd(); r.Next() {
		for a := Relation(1); a <= 6; a++ {
			newTree := tree.Copy()
			newTree.AddArc(i, j, a)
			if newTree.IsFeasible() {
				if j == i+1 && j == (len(ins.items)-1) {
					*resTree = newTree
					return true
				}
				var p, q int
				if j == i+1 {
					p = 0
					q = j + 1
				} else {
					p = i + 1
					q = j
				}
				if res := compare(newTree, p, q, resTree); res {
					return true
				}
			}
		}
	}
	return false
}

type Solver struct {
	ins     *Instance
	resTree *SearchTree
}

func (c *Solver) New(ins *Instance) *Solver {
	c.ins = ins
	c.resTree = new(SearchTree)
	return c
}

func (c *Solver) Solve() bool {
	// TODO: Uncomment.
	//sort.Sort(c.ins.items)
	t := new(SearchTree).Init(c.ins)
	item0 := c.ins.GetItem(0)
	for r := new(Rotation).Init(item0); r.NotEnd(); r.Next() {
		if res := compare(t, 0, 1, &c.resTree); res {
			return true
		}
	}
	return false
}

func (c *Solver) PrintResTree() {
	if c.resTree == nil {
		return
	}
	c.resTree.PrintTree()
}

func (c *Solver) PrintItems() {
	if c.resTree == nil {
		return
	}
	c.resTree.PrintItems()
}
