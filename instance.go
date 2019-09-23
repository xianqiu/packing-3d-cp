package packing_3d_cp

type Box struct {
	L, W, H float64
}

type Item Box

type ItemList []*Item

// -----------------------------------//
// Implement interfaces for sort.Sort //
// -----------------------------------//

func (items ItemList) Len() int {
	return len(items)
}

func (items ItemList) Less(i, j int) bool {
	vi := items[i].L * items[i].W * items[i].H
	vj := items[j].L * items[j].W * items[j].H
	return vi > vj
}

func (items ItemList) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

//-------------------------//
// 3D bin packing instance //
//-------------------------//

type Instance struct {
	box   Box
	items ItemList
}

func (s *Instance) Init() *Instance {
	s.box = Box{
		L: 0,
		W: 0,
		H: 0,
	}
	s.items = make(ItemList, 0)
	return s
}

func (s *Instance) Copy() *Instance {
	ins := new(Instance)
	ins.box = s.box
	ins.items = make([]*Item, 0, len(s.items))
	for _, item := range s.items {
		copyItem := Item{
			item.L,
			item.W,
			item.H,
		}
		ins.items = append(ins.items, &copyItem)
	}
	return ins
}

func (s *Instance) SetBox(length, width, height float64) {
	s.box.L = length
	s.box.W = width
	s.box.H = height
}

func (s *Instance) AddItem(length, width, height float64) {
	item := Item{
		L: length,
		W: width,
		H: height,
	}
	s.items = append(s.items, &item)
}

func (s *Instance) GetItem(index int) *Item {
	return s.items[index]
}

func (s *Instance) GetBox() *Box {
	return &s.box
}

func (s *Instance) GetItems() ItemList {
	return s.items
}
