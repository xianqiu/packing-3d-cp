package packing_3d_cp

import "testing"

func TestSearchTree_IsFeasible1(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(1, 4, 4)
	ins.AddItem(3, 4, 4)
	s := new(SearchTree).Init(ins)
	s.AddArc(0, 1, LEFT)
	if !s.IsFeasible() {
		t.Error("LEFT fail")
	}
	s.Init(ins)
	s.AddArc(0, 1, RIGHT)
	if !s.IsFeasible() {
		t.Error("RIGHT fail")
	}

	ins.Init()
	ins.AddItem(1.1, 4, 4)
	ins.AddItem(3, 4, 4)
	s.Init(ins)
	s.AddArc(0, 1, LEFT)
	if s.IsFeasible() {
		t.Error("LEFT fail")
	}
	s.Init(ins)
	s.AddArc(0, 1, RIGHT)
	if s.IsFeasible() {
		t.Error("RIGHT fail")
	}
}

func TestSearchTree_IsFeasible2(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(4, 1, 4)
	ins.AddItem(4, 3, 4)
	s := new(SearchTree).Init(ins)
	s.AddArc(0, 1, BACK)
	if !s.IsFeasible() {
		t.Error("BACK fail")
	}
	s.Init(ins)
	s.AddArc(0, 1, FRONT)
	if !s.IsFeasible() {
		t.Error("FRONT fail")
	}

	ins.Init()
	ins.AddItem(4, 1, 4)
	ins.AddItem(4, 3.1, 4)
	s.Init(ins)
	s.AddArc(0, 1, BACK)
	if s.IsFeasible() {
		t.Error("BACK fail")
	}
	s.Init(ins)
	s.AddArc(0, 1, FRONT)
	if s.IsFeasible() {
		t.Error("FRONT fail")
	}
}

func TestSearchTree_IsFeasible3(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(4, 4, 1)
	ins.AddItem(4, 4, 3)
	s := new(SearchTree).Init(ins)
	s.AddArc(0, 1, BELOW)
	if !s.IsFeasible() {
		t.Error("BELOW fail")
	}
	s.Init(ins)
	s.AddArc(0, 1, ABOVE)
	if !s.IsFeasible() {
		t.Error("ABOVE fail")
	}

	ins.Init()
	ins.AddItem(4, 4, 1.1)
	ins.AddItem(4, 4, 3)
	s.Init(ins)
	s.AddArc(0, 1, BELOW)
	if s.IsFeasible() {
		t.Error("BELOW fail")
	}
	s.Init(ins)
	s.AddArc(0, 1, ABOVE)
	if s.IsFeasible() {
		t.Error("ABOVE fail")
	}
}

func TestSearchTree_IsFeasible4(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(1, 4, 4)
	ins.AddItem(3, 2, 4)
	ins.AddItem(3, 2, 3)
	ins.AddItem(3, 2, 1.1)
	s := new(SearchTree).Init(ins)
	s.AddArc(0, 1, LEFT)
	s.AddArc(0, 2, LEFT)
	s.AddArc(1, 2, BACK)
	s.AddArc(0, 3, LEFT)
	s.AddArc(1, 3, BACK)
	s.AddArc(2, 3, FRONT)
	println(s.IsFeasible())
	for i := 0; i < len(s.nodes); i++ {
		xa, ya, za := s.GetXyzOfA(i)
		xb, yb, zb := s.GetXyzOfB(i)
		println("item", i, "A =", xa, ya, za, "B =", xb, yb, zb, "Dim =",
			ins.GetItem(i).L, ins.GetItem(i).W, ins.GetItem(i).H)
	}
}
