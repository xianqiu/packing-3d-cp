package packing_3d_cp

import "testing"

func TestIdSet_Copy(t *testing.T) {
	idSet := new(IdSet).Init()
	for i := 0; i < 20; i++ {
		idSet.Add(i)
	}
	copySet := idSet.Copy()
	copySet.Print()
}

func TestIdSet_Union(t *testing.T) {
	set1 := new(IdSet).Init()
	set1.Add(1)
	set1.Add(2)

	set2 := new(IdSet).Init()
	set2.Add(2)
	set2.Add(3)

	set1.Print()
	set1.Union(set2)
	set1.Print()
}
