package packing_3d_cp

type SearchTree struct {
	nodes      map[int]bool
	rl, rw, rh *RelationTree
}

func (s *SearchTree) Init() *SearchTree {
	s.rl = new(RelationTree).Init()
	s.rw = new(RelationTree).Init()
	s.rh = new(RelationTree).Init()
	s.nodes = make(map[int]bool)
	return s
}

func (s *SearchTree) Copy() *SearchTree {
	t := new(SearchTree)
	t.rl = s.rl.Copy()
	t.rw = s.rw.Copy()
	t.rh = s.rh.Copy()
	t.nodes = make(map[int]bool)
	for k := range s.nodes {
		t.nodes[k] = true
	}
	return t
}

type Relation uint8

const (
	LEFT  Relation = 1
	RIGHT Relation = 2
	BACK  Relation = 3
	FRONT Relation = 4
	BELOW Relation = 5
	ABOVE Relation = 6
)

func (s *SearchTree) addNode(ins *Instance, index int) {
	s.nodes[index] = true
	s.rl.AddNode(index, ins.GetItem(index).L)
	s.rw.AddNode(index, ins.GetItem(index).W)
	s.rh.AddNode(index, ins.GetItem(index).H)
}

func (s *SearchTree) AddArc(ins *Instance, i, j int, r Relation) {
	s.addNode(ins, i)
	s.addNode(ins, j)
	switch r {
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

func (s *SearchTree) getBoundaryIds(ins *Instance) map[int]bool {
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

func (s *SearchTree) GetXyzOfB(ins *Instance, id int) (float64, float64, float64) {
	x := s.rl.GetNode(id).location + ins.GetItem(id).L
	y := s.rw.GetNode(id).location + ins.GetItem(id).W
	z := s.rh.GetNode(id).location + ins.GetItem(id).H
	return x, y, z
}

func (s *SearchTree) IsFeasible(ins *Instance) bool {
	boundaryIds := s.getBoundaryIds(ins)
	for id := range boundaryIds {
		l, w, h := s.GetXyzOfB(ins, id)
		if l > ins.box.L || w > ins.box.W || h > ins.box.H {
			return false
		}
	}
	return true
}

func (s *SearchTree) Print() {
	println("-- relation tree: L --")
	s.rl.Print()
	println("-- relation tree: W --")
	s.rw.Print()
	println("-- relation tree: H --")
	s.rh.Print()
}
