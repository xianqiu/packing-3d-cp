package packing_3d_cp

import (
	"container/list"
	"strconv"
	"strings"
)

//--------------//
// RelationNode //
//--------------//

type RelationNode struct {
	id       int
	weight   float64
	location float64
}

func (n *RelationNode) New(id int, weight float64) *RelationNode {
	m := &RelationNode{
		id:       id,
		weight:   weight,
		location: 0,
	}
	return m
}

func (n *RelationNode) Copy() *RelationNode {
	m := new(RelationNode)
	m.id = n.id
	m.weight = n.weight
	m.location = n.location
	return m
}

//--------------//
// RelationTree //
//--------------//

type RelationTree struct {
	nodes       map[int]*RelationNode
	arcs        map[int]*list.List
	boundary    float64
	boundaryIds map[int]bool
}

func (r *RelationTree) Init() *RelationTree {
	r.nodes = make(map[int]*RelationNode)
	r.arcs = make(map[int]*list.List)
	r.boundary = 0
	r.boundaryIds = make(map[int]bool)
	return r
}

func (r *RelationTree) Copy() *RelationTree {
	s := new(RelationTree)
	s.nodes = make(map[int]*RelationNode)
	s.arcs = make(map[int]*list.List)
	s.boundary = r.boundary
	s.boundaryIds = make(map[int]bool)

	for k := range r.nodes {
		s.nodes[k] = r.nodes[k].Copy()
	}
	println(">> copy a tree: start")
	for k := range r.arcs {
		s.arcs[k] = list.New()
		children := r.arcs[k]
		println("k has", children.Len(), "children")
		for e := children.Front(); e != nil; e = e.Next() {
			s.arcs[k].PushBack(e.Value.(int))
			println("copy:", k, "->", e.Value.(int))
		}
	}
	println(">> copy done.")
	for k := range r.boundaryIds {
		s.boundaryIds[k] = true
	}
	return s
}

// Do the following things:
// 1. Add a new node
// 2. Init arc set
// 3. Update boundary ids
func (r *RelationTree) AddNode(id int, weight float64) {
	if _, ok := r.nodes[id]; ok {
		return
	}
	r.nodes[id] = new(RelationNode).New(id, weight)
	r.arcs[id] = list.New()
	r.updateBoundaryIds(id)
}

func (r *RelationTree) GetNode(id int) *RelationNode {
	return r.nodes[id]
}

// Do The following things:
// 1. Add arc i -> j
// 2. AddArc the locations of the subtree rooted at the child
// 3. AddArc the boundary ids of the relation tree
func (r *RelationTree) AddArc(i, j int) {
	r.arcs[i].PushBack(j)
	r.updateLocations(i, j)
	r.updateBoundaryIds(j)
}

func (r *RelationTree) isLeaf(id int) bool {
	return r.arcs[id].Len() == 0
}

// AddArc the boundary ids of the subtree rooted at the node
func (r *RelationTree) updateBoundaryIds(id int) {
	node := r.GetNode(id)
	if r.isLeaf(id) {
		b := node.location + node.weight
		if b > r.boundary {
			r.boundary = b
			r.boundaryIds = map[int]bool{node.id: true}
		} else if b == r.boundary {
			r.boundaryIds[node.id] = true
		}
		return
	}
	// else recursively AddArc childIds of the child
	childIds := r.arcs[node.id]
	for e := childIds.Front(); e != nil; e = e.Next() {
		r.updateBoundaryIds(e.Value.(int))
	}
}

// AddArc locations of the subtree rooted at the child
func (r *RelationTree) updateLocations(i, j int) {
	parent := r.GetNode(i)
	child := r.GetNode(j)
	newLoc := parent.weight + parent.location
	if newLoc > child.location {
		child.location = newLoc
	}
	// if child has no children, then return
	if r.isLeaf(child.id) {
		return
	}
	// else recursively AddArc children of the child
	children := r.arcs[child.id]
	for e := children.Front(); e != nil; e = e.Next() {
		r.updateLocations(child.id, e.Value.(int))
	}
}

func (r *RelationTree) Print() {
	for k := range r.arcs {
		children := r.arcs[k]
		c := make([]string, 0, children.Len())
		for e := children.Front(); e != nil; e = e.Next() {
			c = append(c, r.formatNodeStr(e.Value.(int)))
		}
		println(r.formatNodeStr(k), "->", strings.Join(c, ", "))
	}
}

func (r *RelationTree) formatNodeStr(id int) string {
	node := r.GetNode(id)
	strId := strconv.Itoa(node.id)
	strWeight := strconv.FormatFloat(node.weight, 'f', 2, 64)
	return strId + "(w:" + strWeight + ")"
}
