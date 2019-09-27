package packing_3d_cp

import (
	"testing"
	"time"
)

func TestSolver_Solve(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(1, 4, 4)
	ins.AddItem(3, 2, 4)
	ins.AddItem(3, 2, 3)
	ins.AddItem(3, 2, 1)

	solver := new(Solver).Init(ins)
	solver.Solve()
	println("status", solver.GetStatus())
}

func TestSolver_Solve1(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(5, 5, 5)
	ins.AddItem(2, 5, 5)
	ins.AddItem(3, 4, 5)
	ins.AddItem(3, 1, 1)
	ins.AddItem(3, 2, 1)
	ins.AddItem(3, 2, 1)
	solver := new(Solver).Init(ins)
	solver.Solve()
	println("status", solver.GetStatus())
	//solver.PrintResTree()
	//solver.PrintItems()
}

func TestSolver_Solve2(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 5, 6)
	ins.AddItem(3, 4, 3)
	ins.AddItem(2, 4, 3)
	ins.AddItem(2, 3, 2)
	ins.AddItem(2, 1, 2)
	ins.AddItem(3, 1, 2)
	ins.AddItem(2, 3, 2)
	ins.AddItem(1, 3, 2)
	ins.AddItem(4, 5, 1)

	solver := new(Solver).Init(ins)
	t0 := time.Now().UnixNano()
	solver.Solve()
	println("time cost:", (time.Now().UnixNano()-t0)/1e6)
	println("status", solver.GetStatus())
	//solver.PrintResTree()
	//solver.PrintItems()
}

func TestSolver_Solve3(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 5, 6)
	ins.AddItem(3, 4, 3)
	ins.AddItem(2, 4, 3)
	ins.AddItem(2, 3, 2)
	ins.AddItem(2, 1, 2)
	ins.AddItem(3, 1, 2.1)
	ins.AddItem(3, 3, 2)
	ins.AddItem(4, 5, 1)

	solver := new(Solver).Init(ins)
	t0 := time.Now().UnixNano()
	solver.Solve()
	println("time cost:", (time.Now().UnixNano()-t0)/1e6)
	println("status", solver.GetStatus())
	//solver.PrintResTree()
	//solver.PrintItems()
}

func TestSolver_Solve4(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(76.4, 88.8, 64.1)
	ins.AddItem(76.4, 88.8, 7.9)
	ins.AddItem(76.4, 37.3, 54.9)
	ins.AddItem(76.4, 51.5, 54.9)
	ins.AddItem(55.8, 88.8, 1.3)
	ins.AddItem(20.6, 88.8, 1.3)

	solver := new(Solver).Init(ins)
	t0 := time.Now().UnixNano()
	solver.Solve()
	println("time cost:", (time.Now().UnixNano()-t0)/1e6)
	println("status", solver.GetStatus())
}

func TestSolver_Solve5(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(44.60, 41.80, 35.70)
	ins.AddItem(28.40, 17.00, 44.60)
	ins.AddItem(44.60, 24.80, 19.40)
	ins.AddItem(8.80, 24.80, 44.60)
	ins.AddItem(24.80, 44.60, 7.50)
	ins.AddItem(7.30, 17.00, 44.60)

	solver := new(Solver).Init(ins)
	t0 := time.Now().UnixNano()
	solver.Solve()
	println("time cost:", (time.Now().UnixNano()-t0)/1e6)
	println("status", solver.GetStatus())
}
