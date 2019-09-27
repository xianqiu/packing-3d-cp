package packing_3d_cp

import "fmt"

type SearchTree struct {
	ins        *Instance
	nodes      map[int]bool
	rl, rw, rh *RelationTree
}

func (s *SearchTree) Init(ins *Instance) *SearchTree {
	s.ins = ins
	s.rl = new(RelationTree).Init()
	s.rw = new(RelationTree).Init()
	s.rh = new(RelationTree).Init()
	s.nodes = make(map[int]bool)
	return s
}

func (s *SearchTree) Copy() *SearchTree {
	t := new(SearchTree)
	t.ins = s.ins
	t.rl = s.rl.Copy()
	t.rw = s.rw.Copy()
	t.rh = s.rh.Copy()
	t.nodes = make(map[int]bool)
	for k := range s.nodes {
		t.nodes[k] = true
	}
	return t
}

func (s *SearchTree) addNode(index int) {
	s.nodes[index] = true
	s.rl.AddNode(index, s.ins.GetItem(index).L)
	s.rw.AddNode(index, s.ins.GetItem(index).W)
	s.rh.AddNode(index, s.ins.GetItem(index).H)
}

func (s *SearchTree) AddArc(i, j int, a Relation) {
	s.addNode(i)
	s.addNode(j)
	if s.IsArcFeasible(j, i, a) {
		return
	}
	switch a {
	case LEFT: // i is to the left of j
		s.rl.AddArc(i, j) // Note: "break" is automatic
	case RIGHT: // i is to the right of j
		s.rl.AddArc(j, i)
	case BACK: // i is in the back of j
		s.rw.AddArc(i, j)
	case FRONT: // i is in the front of j
		s.rw.AddArc(j, i)
	case BELOW: // i is below j
		s.rh.AddArc(i, j)
	case ABOVE: // i is above j
		s.rh.AddArc(j, i)
	}
}

func (s *SearchTree) IsArcFeasible(i, j int, a Relation) bool {
	res := false
	switch a {
	case LEFT: // i is to the left of j
		res = s.rl.IsArcFeasible(i, j) // Note: "break" is automatic
	case RIGHT: // i is to the right of j
		res = s.rl.IsArcFeasible(j, i)
	case BACK: // i is in the back of j
		res = s.rw.IsArcFeasible(i, j)
	case FRONT: // i is in the front of j
		res = s.rw.IsArcFeasible(j, i)
	case BELOW: // i is below j
		res = s.rh.IsArcFeasible(i, j)
	case ABOVE: // i is above j
		res = s.rh.IsArcFeasible(j, i)
	}
	return res
}

func (s *SearchTree) getBoundaryIds() map[int]bool {
	ids := make(map[int]bool)
	for k := range s.rl.boundaryIds {
		ids[k] = true
	}
	for k := range s.rw.boundaryIds {
		ids[k] = true
	}
	for k := range s.rh.boundaryIds {
		ids[k] = true
	}
	return ids
}

func (s *SearchTree) GetXyzOfA(id int) (float64, float64, float64) {
	x := s.rl.GetNode(id).location
	y := s.rw.GetNode(id).location
	z := s.rh.GetNode(id).location
	return x, y, z
}

func (s *SearchTree) GetXyzOfB(id int) (float64, float64, float64) {
	x := s.rl.GetNode(id).location + s.ins.GetItem(id).L
	y := s.rw.GetNode(id).location + s.ins.GetItem(id).W
	z := s.rh.GetNode(id).location + s.ins.GetItem(id).H
	return x, y, z
}

func (s *SearchTree) IsFeasible() bool {
	boundaryIds := s.getBoundaryIds()
	eps := 1e-6
	for id := range boundaryIds {
		l, w, h := s.GetXyzOfB(id)
		if l > s.ins.box.L+eps || w > s.ins.box.W+eps || h > s.ins.box.H+eps {
			return false
		}
	}
	return true
}

func (s *SearchTree) PrintTree() {
	s.rl.PrintTree("L")
	s.rw.PrintTree("W")
	s.rh.PrintTree("H")
}

func (s *SearchTree) PrintItem(id int) {
	if s == nil {
		return
	}
	xa, ya, za := s.GetXyzOfA(id)
	xb, yb, zb := s.GetXyzOfB(id)
	l, w, h := s.ins.GetItem(id).L, s.ins.GetItem(id).W, s.ins.GetItem(id).H
	fmt.Printf("+ item %d\n", id)
	fmt.Printf("  - location: A=(%.2f, %.2f, %.2f) B=(%.2f, %.2f, %.2f)\n", xa, ya, za, xb, yb, zb)
	fmt.Printf("  - size: (%.2f, %.2f, %.2f)\n", l, w, h)
}

func (s *SearchTree) PrintItems() {
	for id := range s.nodes {
		s.PrintItem(id)
	}
}
