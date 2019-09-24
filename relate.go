package packing_3d_cp

type Relation uint8

const (
	LEFT  Relation = 1
	RIGHT Relation = 2
	BACK  Relation = 3
	FRONT Relation = 4
	BELOW Relation = 5
	ABOVE Relation = 6
)

type Relate struct {
	relations []Relation
	index     int
}

func (r *Relate) Init(itemI Item, itemJ Item) *Relate {
	r.index = 0
	r.relations = []Relation{LEFT, RIGHT, BACK, FRONT, BELOW, ABOVE}
	return r
}

func (r *Relate) NotEnd() bool {
	return r.index != len(r.relations)
}

func (r *Relate) Next() {
	r.index += 1
}

func (r *Relate) GetRelation() Relation {
	return r.relations[r.index]
}

// For item 1 and item 2 only.
// Only three relations need to be considered.
func (r *Relate) SetSpecial() {
	r.relations = []Relation{LEFT, BACK, BELOW}
}
