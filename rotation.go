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
	mode    int
	rotated map[Direction]Item
}

func (r *Rotation) Init(item *Item) *Rotation {
	r.di = LWH
	r.item = item
	// 4 modes
	if r.item.L == r.item.W && r.item.L == r.item.H {
		r.mode = 1
	} else if r.item.L == r.item.W {
		r.mode = 2
	} else if r.item.L == r.item.H {
		r.mode = 3
	} else if r.item.W == r.item.H {
		r.mode = 4
	} else {
		r.mode = 5
	}

	return r
}

func (r *Rotation) nextFull() {
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

// L = W
func (r *Rotation) nextLeqW() {
	switch r.di {
	case LWH:
		r.item.W, r.item.H = r.item.H, r.item.W
		r.di = LHW
	case LHW:
		r.item.L, r.item.W = r.item.W, r.item.L
		r.di = HLW
	case HLW:
		r.di = 0 // mark as end
	}
}

// L = H
func (r *Rotation) nextLeqH() {
	switch r.di {
	case LWH:
		r.item.W, r.item.H = r.item.H, r.item.W
		r.di = LHW
	case LHW:
		r.item.L, r.item.H = r.item.H, r.item.L
		r.di = WHL
	case WHL:
		r.di = 0 // mark as end
	}
}

// W = H
func (r *Rotation) nextWeqH() {
	switch r.di {
	case LWH:
		r.item.L, r.item.W = r.item.W, r.item.L
		r.di = WHL
	case WHL:
		r.item.W, r.item.H = r.item.H, r.item.W
		r.di = WLH
	case WLH:
		r.di = 0 // mark as end
	}
}

// L = W = H
func (r *Rotation) nextLeqWeqH() {
	r.di = 0
}

func (r *Rotation) Next() {
	if r.mode == 1 {
		r.nextLeqWeqH()
		return
	}
	if r.mode == 2 {
		r.nextLeqW()
		return
	}
	if r.mode == 3 {
		r.nextLeqH()
		return
	}
	if r.mode == 4 {
		r.nextWeqH()
		return
	}
	r.nextFull()
}

func (r *Rotation) NotEnd() bool {
	return r.di > 0
}
