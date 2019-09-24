package packing_3d_cp

import "testing"

func TestBenchmark_Load(t *testing.T) {
	b := new(Benchmark).Init()
	b.Load()
	b.Run()
}
