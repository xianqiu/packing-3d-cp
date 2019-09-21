package packing_3d_cp

import (
	"sort"
	"testing"
)

func TestInstance_Init(t *testing.T) {
	s := new(Instance).Init()
	s.SetBox(4, 5, 6)
	s.AddItem(1, 1, 1)
	s.AddItem(2, 2, 2)
	println("item number =", len(s.items))

	for i := 0; i < len(s.items); i++ {
		println("item", i, s.GetItem(i).L, s.GetItem(i).W, s.GetItem(i).H)
	}
	println("Sort W.r.t. volume")
	// sort item list
	sort.Sort(s.GetItems())
	for i := 0; i < len(s.items); i++ {
		println("item", i, s.GetItem(i).L, s.GetItem(i).W, s.GetItem(i).H)
	}

	s.GetItem(0).L = 100
	println(s.GetItem(0).L)
}
