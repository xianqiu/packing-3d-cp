package packing_3d_cp

import (
	"sort"
	"time"
)

type Solver struct {
	ins          *Instance
	resTree      *SearchTree
	timeout      int // millisecond
	status       STATUS
	threadStatus chan STATUS
}

type STATUS uint8

const (
	UNEXPECTED = 0
	FEASIBLE   = 1
	INFEASIBLE = 2
	TIMEOUT    = 3
)

func (s *Solver) Init(ins *Instance) *Solver {
	s.ins = ins
	s.resTree = new(SearchTree)
	s.timeout = 1000
	s.status = UNEXPECTED
	return s
}

func (s *Solver) SetTimeout(miniSeconds int) {
	s.timeout = miniSeconds
}

func (s *Solver) isTimeout(t0 int64) bool {
	t1 := time.Now().UnixNano()
	if t1-t0 > int64(s.timeout)*1e6 {
		return true
	}
	return false
}

func (s *Solver) returnFeasible(resTree *SearchTree) bool {
	s.resTree = resTree
	s.status = FEASIBLE
	return true
}

func (s *Solver) returnInfeasible() bool {
	s.status = INFEASIBLE
	return false
}

func (s *Solver) returnTimeout() bool {
	s.status = TIMEOUT
	return false
}

// Given Items i and j, compute the next pair of Items to be compared
func (s *Solver) nextPair(i, j int) (int, int) {
	var p, q int
	if j == i+1 {
		p = 0
		q = j + 1
	} else {
		p = i + 1
		q = j
	}
	return p, q
}

func (s *Solver) compare(tree *SearchTree, i, j int, t0 int64) bool {
	ins := tree.ins
	item := ins.GetItem(j)
	for r := new(Rotate).Init(item); r.NotEnd(); r.Next() {
		if s.isTimeout(t0) {
			return s.returnTimeout()
		}
		a := new(Relate).Init(*ins.GetItem(i), *ins.GetItem(j))
		// consider LEFT, BACK, BELOW for Items 1 and 2
		if i == 1 && j == 2 {
			a.SetSpecial()
		}
		for ; a.NotEnd(); a.Next() {
			newTree := tree.Copy()
			newTree.AddArc(i, j, a.GetRelation())
			if newTree.IsFeasible() {
				if j == i+1 && j == (len(ins.items)-1) {
					return s.returnFeasible(newTree)
				}
				p, q := s.nextPair(i, j)
				if res := s.compare(newTree, p, q, t0); res {
					return true
				}
			}
		}
	}
	return s.returnInfeasible()
}

func (s *Solver) checkTriviallyInfeasible() bool {
	// TODO:
	return false
}

func (s *Solver) Solve() {
	now := time.Now().UnixNano()
	// Step 1: Exclude trivial cases
	if isTrivial := s.checkTriviallyInfeasible(); isTrivial {
		s.status = INFEASIBLE
		return
	}
	// Step 2: Sort instance w.r.t. volume
	sort.Sort(s.ins.items)
	// Step 3: Initialize search tree
	tree := new(SearchTree).Init(s.ins)
	// Step 4: Rotate item 1 and compare (1, 2) and the pairs after it.
	item0 := s.ins.GetItem(0)
	markAsTimeout := false
	for r := new(Rotate).Init1(item0, &s.ins.box); r.NotEnd(); r.Next() {
		s.compare(tree, 0, 1, now)
		if s.status == FEASIBLE {
			return
		} else if s.status == TIMEOUT {
			markAsTimeout = true
		}
	}
	if markAsTimeout {
		s.status = TIMEOUT
	}
}

// For multi-thread
// Note: Do not use it alone
func (s *Solver) solve1(resTree chan *SearchTree, result chan STATUS, infeasibleCount chan int) {
	now := time.Now().UnixNano()
	// Step 1: Sort instance w.r.t. volume
	sort.Sort(s.ins.items)
	// Step 2: Initialize search tree
	tree := new(SearchTree).Init(s.ins)
	// Step 3: compare (1, 2) and the pairs after it.
	s.compare(tree, 0, 1, now)
	if s.status == FEASIBLE {
		result <- FEASIBLE
		resTree <- s.resTree
	} else if s.status == INFEASIBLE {
		a := <-infeasibleCount - 1
		if a == 0 {
			result <- INFEASIBLE
		} else {
			infeasibleCount <- a
		}
	}
}

func (s *Solver) PrintResTree() {
	if s.resTree == nil {
		return
	}
	s.resTree.PrintTree()
}

func (s *Solver) PrintItems() {
	if s.resTree == nil {
		return
	}
	s.resTree.PrintItems()
}

func (s *Solver) GetStatus() STATUS {
	return s.status
}
