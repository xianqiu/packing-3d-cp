package packing_3d_cp

import "testing"

func TestSolver_Solve(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(4, 4, 4)
	ins.AddItem(1, 4, 4)
	ins.AddItem(3, 2, 4)
	ins.AddItem(3, 2, 3)
	ins.AddItem(3, 2, 1)

	solver := new(Solver).New(ins)
	println(solver.Solve())
}

func TestSolver_Solve1(t *testing.T) {
	ins := new(Instance).Init()
	ins.SetBox(5, 5, 5)
	ins.AddItem(2, 5, 5)
	ins.AddItem(3, 4, 5)
	ins.AddItem(3, 5, 1.1)

	solver := new(Solver).New(ins)

	println(solver.Solve())
	solver.PrintResTree()
	solver.PrintItems()
}
