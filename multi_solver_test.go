package packing_3d_cp

import (
	"testing"
	"time"
)

func TestMultiSolver_Solve(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 5, 6)
	ins.AddItem(3, 4, 3)
	ins.AddItem(2, 4, 3)
	ins.AddItem(2, 3, 2)
	ins.AddItem(2, 1, 2)
	ins.AddItem(3, 1, 2)
	ins.AddItem(2, 3, 2)
	ins.AddItem(1, 3, 2.1)
	ins.AddItem(4, 5, 1)

	solver := new(MultiSolver).Init(ins)
	t0 := time.Now().UnixNano()
	solver.Solve()
	println("time cost:", (time.Now().UnixNano()-t0)/1e6)
	println("status", solver.GetStatus())
	//solver.PrintResTree()
	//solver.PrintItems()
}
