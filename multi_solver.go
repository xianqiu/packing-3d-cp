package packing_3d_cp

import "time"

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
	//TODO: exclude trivially infeasible case
	item0 := m.ins.GetItem(0)
	result := make(chan STATUS)
	resTree := make(chan *SearchTree)
	infeasibleCount := make(chan int)
	r := new(Rotate).Init1(item0, &m.ins.box)
	go func() {
		infeasibleCount <- r.GetDiNum()
	}()
	for ; r.NotEnd(); r.Next() {
		solver := new(Solver).Init(m.ins.Copy())
		go solver.solve1(resTree, result, infeasibleCount)
	}
	select {
	case res := <-result:
		m.status = res
		if res == FEASIBLE {
			<-resTree
		}
	case <-time.After(1000 * time.Millisecond):
		m.status = TIMEOUT
	}
	close(result)
	close(resTree)
	close(infeasibleCount)
}

func (m *MultiSolver) GetStatus() STATUS {
	return m.status
}
