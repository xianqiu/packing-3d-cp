package packing_3d_cp

type MultiSolver struct {
	ins    *Instance
	status STATUS
}

func (m *MultiSolver) Init(ins *Instance) *MultiSolver {
	m.ins = ins
	m.status = UNEXPECTED
	return m
}

func (m *MultiSolver) Solve() {
	item0 := m.ins.GetItem(0)
	result := make(chan *Solver, 6)
	goroutineNum := 0
	for r := new(Rotate).Init1(item0, &m.ins.box); r.NotEnd(); r.Next() {
		goroutineNum += 1
		solver := new(Solver).Init(m.ins.Copy())
		go solver.solve1(result)
	}
	defer close(result)
	// Collect result
	timeoutCount := 0
	infeasibleCount := 0
	index := 0
	for r := range result {
		index += 1
		if r.GetStatus() == FEASIBLE {
			m.status = FEASIBLE
			return
		}
		if r.GetStatus() == TIMEOUT {
			timeoutCount += 1
		} else if r.GetStatus() == INFEASIBLE {
			infeasibleCount += 1
		}
		if index == goroutineNum {
			break
		}
	}
	if timeoutCount > 0 {
		m.status = TIMEOUT
	} else {
		m.status = INFEASIBLE
	}
}

func (m *MultiSolver) GetStatus() STATUS {
	return m.status
}
