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
	di      Direction
	item    *Item
	rotated map[Direction]Item
}

func (r *Rotation) Init(item *Item) *Rotation {
	r.di = LWH
	r.item = item
	return r
}

func (r *Rotation) Next() {
	switch r.di {
	case LWH:
		r.item.W, r.item.H = r.item.H, r.item.W
		r.di = LHW
	case LHW:
		r.item.L, r.item.H = r.item.H, r.item.L
		r.di = WHL
	case WHL:
		r.item.W, r.item.H = r.item.H, r.item.W
		r.di = WLH
	case WLH:
		r.item.L, r.item.H = r.item.H, r.item.L
		r.di = HLW
	case HLW:
		r.item.W, r.item.H = r.item.H, r.item.W
		r.di = HWL
	case HWL:
		r.di = 0 // mark as end
	}
}

func (r *Rotation) NotEnd() bool {
	return r.di > 0
}
