package packing_3d_cp

import (
	"fmt"
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
	arcs        map[int][]int
	boundary    float64
	boundaryIds map[int]bool
}

func (r *RelationTree) Init() *RelationTree {
	r.nodes = make(map[int]*RelationNode)
	r.arcs = make(map[int][]int)
	r.boundary = 0
	r.boundaryIds = make(map[int]bool)
	return r
}

func (r *RelationTree) Copy() *RelationTree {
	s := new(RelationTree)
	s.nodes = make(map[int]*RelationNode)
	s.arcs = make(map[int][]int)
	s.boundary = r.boundary
	s.boundaryIds = make(map[int]bool)

	for k := range r.nodes {
		s.nodes[k] = r.nodes[k].Copy()
	}
	for k := range r.arcs {
		s.arcs[k] = make([]int, 0, len(r.arcs))
		childIds := r.arcs[k]
		for _, id := range childIds {
			s.arcs[k] = append(s.arcs[k], id)
		}
	}
	for k := range r.boundaryIds {
		s.boundaryIds[k] = true
	}
	return s
}

// Do the following things:
// 1. If node exists, update node weight and the locations of the subtree rooted at the node
// 2. If node does not exist, add a new node and init arc set
// 3. Update boundary ids
func (r *RelationTree) AddNode(id int, weight float64) {
	if _, ok := r.nodes[id]; ok {
		r.updateNode(id, weight)
	} else {
		r.nodes[id] = new(RelationNode).New(id, weight)
		r.arcs[id] = make([]int, 0)
	}
	r.updateBoundaryIds(id)
}

// Do the following things:
// 1. Update node weight
// 2. Update locations of the subtree rooted at the node
func (r *RelationTree) updateNode(id int, weight float64) {
	if r.GetNode(id).weight == weight {
		return
	}
	r.GetNode(id).weight = weight
	childIds := r.arcs[id]
	for _, childId := range childIds {
		r.updateLocations(id, childId)
	}
}

func (r *RelationTree) GetNode(id int) *RelationNode {
	return r.nodes[id]
}

// Do The following things:
// 1. Add arc i -> j
// 2. AddArc the locations of the subtree rooted at the child
// 3. AddArc the boundary ids of the relation tree
func (r *RelationTree) AddArc(i, j int) {
	r.arcs[i] = append(r.arcs[i], j)
	r.updateLocations(i, j)
	r.updateBoundaryIds(j)
}

func (r *RelationTree) isLeaf(id int) bool {
	return len(r.arcs[id]) == 0
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
	for _, id := range childIds {
		r.updateBoundaryIds(id)
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
	// if child has no childIds, then return
	if r.isLeaf(child.id) {
		return
	}
	// else recursively add arc to childIds of the child
	childIds := r.arcs[child.id]
	for _, id := range childIds {
		r.updateLocations(child.id, id)
	}
}

func (r *RelationTree) PrintTree(name string) {
	fmt.Printf("+ relation tree: %s\n", name)
	if r == nil {
		return
	}
	for k := range r.arcs {
		childIds := r.arcs[k]
		c := make([]string, 0, len(childIds))
		for _, id := range childIds {
			c = append(c, r.formatNodeStr(id))
		}
		fmt.Printf("  - %s -> %s\n", r.formatNodeStr(k), strings.Join(c, ", "))
	}
}

func (r *RelationTree) formatNodeStr(id int) string {
	node := r.GetNode(id)
	strId := strconv.Itoa(node.id)
	strWeight := strconv.FormatFloat(node.weight, 'f', 2, 64)
	return strId + "(w:" + strWeight + ")"
}
