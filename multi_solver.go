package packing_3d_cp

type MultiSolver struct {
	ins *Instance
}

func (m *MultiSolver) Init(ins *Instance) {
	m.ins = ins
}

func (m *MultiSolver) Solve() STATUS {
	item0 := m.ins.GetItem(0)
	result := make(chan *Solver, 6)
	for r := new(Rotate).Init1(item0, &m.ins.box); r.NotEnd(); r.Next() {
		solver := new(Solver).Init(m.ins.Copy())
		go solver.solve1(result)
	}
	defer close(result)
	// TODO: collect result.
	return UNEXPECTED
}
