package packing_3d_cp

import "testing"

func TestRelate_Next(t *testing.T) {
	itemI := Item{
		1,
		2,
		3,
	}
	itemJ := Item{
		4,
		5,
		6,
	}
	for r := new(Relate).Init(itemI, itemJ); r.NotEnd(); r.Next() {
		println(r.GetRelation())
	}
}
