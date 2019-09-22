package packing_3d_cp

import "testing"

func TestRotation_Next1(t *testing.T) {
	item := Item{1, 2, 3}

	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.Next() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_Next2(t *testing.T) {
	item := Item{1, 1, 1}

	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.Next() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_Next3(t *testing.T) {
	item := Item{1, 1, 3}

	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.Next() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_Next4(t *testing.T) {
	item := Item{1, 2, 1}

	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.Next() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_Next5(t *testing.T) {
	item := Item{1, 2, 2}

	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.Next() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_NextLeqW(t *testing.T) {
	item := Item{1, 1, 3}
	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.nextLeqW() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_NextLeqH(t *testing.T) {
	item := Item{1, 2, 1}
	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.nextLeqH() {
		println(item.L, item.W, item.H)
	}
}

func TestRotation_NextWeqH(t *testing.T) {
	item := Item{1, 2, 2}
	items := ItemList{&item}
	r := new(Rotate).Init(items[0])
	for ; r.NotEnd(); r.nextWeqH() {
		println(item.L, item.W, item.H)
	}
}
