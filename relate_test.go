package packing_3d_cp

import "testing"

func TestRelate_Next(t *testing.T) {
	for r := new(Relate).Init(); r.NotEnd(); r.Next() {
		println(r.GetRelation())
	}
}
