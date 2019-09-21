package packing_3d_cp

import (
	"testing"
)

func TestRelationTree_AddChild(t *testing.T) {
	r := new(RelationTree).Init()
	r.AddNode(1, 1)
	r.AddNode(2, 2)
	r.AddArc(1, 2)
	r.AddNode(3, 3)
	r.AddArc(2, 3)
	r.AddNode(4, 4)
	r.AddArc(2, 4)
	r.AddNode(5, 5)
	r.AddArc(5, 1)
	r.AddNode(6, 6)
	r.AddArc(6, 2)
	r.AddArc(5, 6)

	result := make(map[int]float64)
	//result[r.Root] = 0
	result[1] = 5
	result[2] = 11
	result[3] = 13
	result[4] = 13
	result[5] = 0
	result[6] = 5

	for id, node := range r.nodes {
		if result[id] != node.location {
			t.Error("ERROR >> id", id, "location:", node, "expected:", result[id])
		} else {
			t.Log("PASS >> id", id, "location:", node.location)
		}
	}
	if r.boundary != 17 {
		t.Error("boundary=", r.boundary, "expected =", 17)
	}

}

func TestRelationTreeNode_Copy(t *testing.T) {
	r1 := &RelationNode{
		id:     1,
		weight: 1,
	}
	r2 := r1.Copy()
	r2.id = 2
	r2.weight = 2

	if r2.id != r1.id+1 || r2.weight != r1.weight+1 {
		t.Error("r1.id =", r1.id, "r1.weight =", r1.weight, "r2.id=", r2.id, "r2.weight=", r2.weight)
	}
}

func TestRelationTree_Copy(t *testing.T) {
	r := new(RelationTree).Init()
	r.AddNode(1, 1)
	r.AddNode(2, 2)
	r.AddArc(1, 2)
	r.AddNode(3, 3)
	r.AddArc(2, 3)
	r.AddNode(4, 4)
	r.AddArc(2, 4)

	newTree := new(RelationTree)
	newTree = r.Copy()
	// Add a child to the new tree
	newTree.AddNode(5, 5)
	newTree.AddArc(5, 1)
	n1 := len(newTree.nodes)
	n2 := len(r.nodes)
	if n1 != n2+1 {
		t.Error("Node numbers: new tree =", n1, "old tree =", n2)
	}

	// Modify the weight of node 3.
	newTree.GetNode(3).weight = 300
	node := r.GetNode(3)
	if node.weight != 3 {
		t.Error("node: id=", node.id, "weight =", node.weight)
	}

	// Modify the location of node 3.
	newTree.GetNode(3).location = 100
	if r.GetNode(3).location != 3 {
		t.Error("node 3(old): location=", r.GetNode(3).location)
	}
}

func TestRelationTree_BoundaryNodes1(t *testing.T) {
	r := new(RelationTree).Init()
	r.AddNode(1, 1)
	r.AddNode(2, 2)
	r.AddArc(1, 2)
	r.AddNode(3, 3)
	r.AddArc(2, 3)
	r.AddNode(4, 3)
	r.AddArc(2, 4)
	r.AddNode(5, 5)
	r.AddArc(1, 5)

	for i := 3; i < 6; i++ {
		if _, ok := r.boundaryIds[i]; !ok {
			t.Error(i, " is expected to be a boundary node")
		}
	}
	if r.boundary != 6 {
		t.Error("boundary =", r.boundary, "expected = 6")
	}
}

func TestRelationTree_BoundaryNodes2(t *testing.T) {
	r := new(RelationTree).Init()
	r.AddNode(1, 1)
	r.AddNode(2, 2)
	r.AddArc(1, 2)
	r.AddNode(3, 3)
	r.AddArc(2, 3)
	r.AddNode(4, 3)
	r.AddArc(2, 4)
	r.Print()
	if len(r.boundaryIds) != 2 {
		t.Error("|boundary ids| =", len(r.boundaryIds))
	}
	if _, ok := r.boundaryIds[3]; !ok {
		t.Error("3 is expected to be a boundary node")
	}
	if _, ok := r.boundaryIds[4]; !ok {
		t.Error("4 is expected to be a boundary node")
	}

	r.AddNode(5, 10)
	r.AddArc(1, 5)

	if len(r.boundaryIds) != 1 {
		t.Error("|boundary ids| =", len(r.boundaryIds))
	}
	if _, ok := r.boundaryIds[5]; !ok {
		t.Error("5 is expected")
	}

	r = r.Init()
	r.AddNode(1, 1)
	if len(r.boundaryIds) != 1 {
		t.Error("|boundary ids| =", len(r.boundaryIds))
	}
}
