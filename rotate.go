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

type Rotate struct {
	di   Direction
	item *Item
	mode int
}

func (r *Rotate) Init(item *Item) *Rotate {
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

// For the first item only.
func (r *Rotate) Init1(item *Item, box *Box) *Rotate {
	r.di = LWH
	r.item = item
	// 4 modes
	if (r.item.L == r.item.W && r.item.L == r.item.H) ||
		(box.L == box.W && box.L == box.H) {
		r.mode = 1
	} else if r.item.L == r.item.W || box.L == box.W {
		r.mode = 2
	} else if r.item.L == r.item.H || box.L == box.H {
		r.mode = 3
	} else if r.item.W == r.item.H || box.W == box.H {
		r.mode = 4
	} else {
		r.mode = 5
	}

	return r
}

func (r *Rotate) nextFull() {
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
func (r *Rotate) nextLeqW() {
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
func (r *Rotate) nextLeqH() {
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
func (r *Rotate) nextWeqH() {
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
func (r *Rotate) nextLeqWeqH() {
	r.di = 0
}

func (r *Rotate) Next() {
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

func (r *Rotate) NotEnd() bool {
	return r.di > 0
}

func (r *Rotate) GetDiNum() int {
	switch r.mode {
	case 1:
		return 1
	case 2, 3, 4:
		return 3
	case 5:
		return 6
	}
	return -1
}
