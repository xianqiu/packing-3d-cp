package packing_3d_cp

type Direction uint8

const (
	LWH Direction = 1
	LHW Direction = 2
	WHL Direction = 3
	WLH Direction = 4
	HLW Direction = 5
	HWL Direction = 6
)

type Rotation struct {
	di   Direction
	item *Item
}

func (r *Rotation) Init(item *Item) *Rotation {
	r.di = LWH
	r.item = item
	return r
}

func (r *Rotation) Next() {
	switch r.di {
	case LWH:
		r.item.H, r.item.W = r.item.W, r.item.H // L, H, W
		r.di = LHW
		break
	case LHW:
		r.item.L, r.item.H = r.item.H, r.item.L // W, H, L
		r.di = WHL
		break
	case WHL:
		r.item.W, r.item.H = r.item.H, r.item.W // W, L, H
		r.di = WLH
		break
	case WLH:
		r.item.L, r.item.H = r.item.H, r.item.L // H, L, W
		r.di = HLW
		break
	case HLW:
		r.item.W, r.item.H = r.item.H, r.item.W // H, W, L
		r.di = HWL
		break
	case HWL:
		r.di = 0 // mark as end
		break
	}
}

func (r *Rotation) NotEnd() bool {
	return r.di > 0
}
