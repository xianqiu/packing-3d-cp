package packing_3d_cp

import "testing"

func TestSolver_Solve(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(1, 4, 4)
	ins.AddItem(3, 2, 4)
	ins.AddItem(3, 2, 3)
	ins.AddItem(3, 2, 1)

	solver := new(Solver).Init(ins)
	println(solver.Solve())
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
	println(solver.Solve())
	println("status", solver.GetStatus())
	//solver.PrintResTree()
	//solver.PrintItems()
}

func TestSolver_Solve2(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 5, 5)
	ins.AddItem(3, 4, 5)
	ins.AddItem(2, 4, 3)
	ins.AddItem(2, 3, 2)
	ins.AddItem(2, 1, 2)

	solver := new(Solver).Init(ins)
	println(solver.Solve())
	println("status", solver.GetStatus())
	//solver.PrintResTree()
	//solver.PrintItems()
}
