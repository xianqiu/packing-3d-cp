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
	id          int
	weight      float64
	location    float64
	childrenIds *IdSet
	parentId    int
}

func (n *RelationNode) New(id int, weight float64) *RelationNode {
	m := &RelationNode{
		id:          id,
		weight:      weight,
		location:    0,
		childrenIds: new(IdSet).Init(),
		parentId:    -1,
	}
	return m
}

func (n *RelationNode) Copy() *RelationNode {
	m := new(RelationNode)
	m.id = n.id
	m.weight = n.weight
	m.location = n.location
	m.childrenIds = n.childrenIds.Copy()
	m.parentId = n.parentId
	return m
}

func (n *RelationNode) HasParent() bool {
	return n.parentId != -1
}

//--------------//
// RelationTree //
//--------------//

type RelationTree struct {
	nodes       map[int]*RelationNode
	boundary    float64
	boundaryIds *IdSet
}

func (r *RelationTree) Init() *RelationTree {
	r.nodes = make(map[int]*RelationNode)
	r.boundary = 0
	r.boundaryIds = new(IdSet).Init()
	return r
}

func (r *RelationTree) Copy() *RelationTree {
	s := new(RelationTree)
	s.nodes = make(map[int]*RelationNode)
	s.boundary = r.boundary
	for k := range r.nodes {
		s.nodes[k] = r.nodes[k].Copy()
	}
	s.boundaryIds = r.boundaryIds.Copy()
	return s
}

// Do the following things:
// 1. If node exists, update node weight and the locations of its children
// 2. If node does not exist, add a new node
// 3. Update boundary ids
func (r *RelationTree) AddNode(id int, weight float64) {
	if _, ok := r.nodes[id]; ok {
		r.updateNodeWeight(id, weight)
	} else {
		r.nodes[id] = new(RelationNode).New(id, weight)
	}
	r.updateBoundaryIds(id)
}

func (r *RelationTree) GetNode(id int) *RelationNode {
	return r.nodes[id]
}

// Do The following things:
// 1. Add arc i -> j and mark i as j's parent
// 2. Add arc i -> j's children
// 3. Add i's children to i's parent
// 4. Update the locations of j and j's children
// 5. Update the boundary ids of the relation tree
func (r *RelationTree) AddArc(i, j int) {
	parent := r.GetNode(i)
	child := r.GetNode(j)
	child.parentId = parent.id
	parent.childrenIds.Add(child.id)
	parent.childrenIds.Union(child.childrenIds)
	if parent.HasParent() {
		grandpa := r.GetNode(parent.parentId)
		grandpa.childrenIds.Union(parent.childrenIds)
	}
	r.updateNodeLocationFromArc(i, j)
	r.updateBoundaryIds(j)
	for childId := range *r.GetNode(j).childrenIds {
		r.updateBoundaryIds(childId)
	}
}

func (r *RelationTree) HasChildren(id int) bool {
	return r.GetNode(id).childrenIds.Size() > 0
}

// If the node is a boundary node
// Then add its id to the boundary id set
func (r *RelationTree) updateBoundaryIds(id int) {
	if r.HasChildren(id) {
		return
	}
	node := r.GetNode(id)
	b := node.location + node.weight
	if b > r.boundary {
		r.boundary = b
		r.boundaryIds = new(IdSet).Init()
		r.boundaryIds.Add(node.id)
	} else if b == r.boundary {
		r.boundaryIds.Add(node.id)
	}
}

// Update node j with new location: i.location + i.weight
func (r *RelationTree) updateNodeLocationFromArc(i, j int) {
	parent := r.GetNode(i)
	child := r.GetNode(j)
	r.updateNodeLocation(child.id, parent.weight+parent.location)
}

// Do the following things:
// 1. Update location of node
// 2. Update the locations of its children
func (r *RelationTree) updateNodeLocation(id int, newLoc float64) {
	node := r.GetNode(id)
	delta := newLoc - node.location
	if delta == 0 {
		return
	}
	node.location = newLoc
	for childId := range *node.childrenIds {
		r.GetNode(childId).location += delta
	}
}

// Do the following things:
// 1. Update node weight
// 2. Update the locations of its children
// 3. Update boundary ids
func (r *RelationTree) updateNodeWeight(id int, weight float64) {
	node := r.GetNode(id)
	delta := weight - node.weight
	if delta == 0 {
		return
	}
	for childId := range *node.childrenIds {
		r.GetNode(childId).location += delta
	}
}

func (r *RelationTree) IsArcExist(i, j int) bool {
	node := r.GetNode(i)
	if node == nil {
		return false
	}
	return node.childrenIds.IsExist(j)
}

func (r *RelationTree) PrintTree(name string) {
	fmt.Printf("+ relation tree: %s\n", name)
	if r == nil {
		return
	}
	for k, node := range r.nodes {
		childIds := node.childrenIds
		c := make([]string, 0, childIds.Size())
		for id := range *childIds {
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
