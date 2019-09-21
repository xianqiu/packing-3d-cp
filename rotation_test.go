package packing_3d_cp

import "testing"

func TestRotation_Next(t *testing.T) {
	item := &Item{1, 2, 3}

	r := new(Rotation).Init(item)
	for ; r.NotEnd(); r.Next() {
		println(item.L, item.W, item.H)
	}
}
