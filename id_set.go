package packing_3d_cp

import "strconv"

type IdSet map[int]bool

func (d *IdSet) Init() *IdSet {
	*d = make(map[int]bool)
	return d
}

func (d *IdSet) Copy() *IdSet {
	e := new(IdSet)
	*e = make(IdSet)
	for id := range *d {
		(*e)[id] = true
	}
	return e
}

func (d *IdSet) IsExist(id int) bool {
	if _, ok := (*d)[id]; ok {
		return true
	}
	return false
}

func (d *IdSet) Add(id int) {
	(*d)[id] = true
}

func (d *IdSet) Union(anotherIdSet *IdSet) {
	for id := range *anotherIdSet {
		(*d)[id] = true
	}
}

func (d *IdSet) Size() int {
	return len(*d)
}

func (d *IdSet) Print() {
	n := 10
	fullStr := ""
	i := 0
	for id := range *d {
		i++
		idStr := strconv.Itoa(id)
		if fullStr == "" {
			fullStr = idStr
			continue
		}
		if fullStr[len(fullStr)-1] == '\n' {
			fullStr += " " + idStr
			continue
		}
		fullStr += ", " + idStr
		if i%n == 0 && i != len(*d) {
			fullStr += "\n"
		}
	}
	println("{" + fullStr + "}")
}
