package packing_3d_cp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Benchmark struct {
	instances     map[string]Instance
	answers       map[string]STATUS
	solvedNum     int
	timeCost      float64
	errId         string
	maxItemNumber int
}

func (b *Benchmark) Init() *Benchmark {
	b.instances = make(map[string]Instance)
	b.answers = make(map[string]STATUS)
	b.maxItemNumber = 0
	return b
}

type JsonItems struct {
	Length float64
	Width  float64
	Height float64
	Count  int
}

type JsonInstance struct {
	Bin   JsonItems
	Items []JsonItems
}

func (b *Benchmark) parseRow(row []string) (string, JsonInstance, string) {
	id := row[0]
	answer := row[2]
	res := JsonInstance{}
	err := json.Unmarshal([]byte(row[1]), &res)
	if err != nil {
		println("json.Unmarshal:", err.Error())
	}
	return id, res, answer
}

func (b *Benchmark) setAnswer(id, answer string) {
	a, err := strconv.ParseInt(answer, 10, 64)
	if err != nil {
		a = 0
	}
	switch a {
	case 1:
		b.answers[id] = FEASIBLE
	case -1:
		b.answers[id] = INFEASIBLE
	default:
		b.answers[id] = UNEXPECTED
	}
}

func (b *Benchmark) setInstance(id string, jsonIns JsonInstance) {
	ins := new(Instance).Init()
	ins.SetBox(jsonIns.Bin.Length, jsonIns.Bin.Width, jsonIns.Bin.Height)
	for _, jsonItems := range jsonIns.Items {
		for i := 0; i < jsonItems.Count; i++ {
			ins.AddItem(jsonItems.Length, jsonItems.Width, jsonItems.Height)
		}
	}
	b.instances[id] = *ins
}

func (b *Benchmark) Load() {
	fileName := "benchmark.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	br := bufio.NewReader(file)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		row := strings.Split(string(line), "\t")
		id, jsonIns, answer := b.parseRow(row)
		b.setAnswer(id, answer)
		b.setInstance(id, jsonIns)
	}
	if err := file.Close(); err != nil {
		println(err.Error())
	}
}

func (b *Benchmark) Run() {
	t0 := time.Now().UnixNano()
	solver := new(Solver)
	for id, ins := range b.instances {
		solver.Init(&ins)
		if b.maxItemNumber > 0 && len(ins.items) > b.maxItemNumber {
			continue
		}
		solver.Solve()
		if solver.GetStatus() == FEASIBLE && b.answers[id] == INFEASIBLE ||
			solver.GetStatus() == INFEASIBLE && b.answers[id] == FEASIBLE {
			b.errId = id
			break
		}
	}
	b.timeCost = float64(time.Now().UnixNano()-t0) / 1e9
}

func (b *Benchmark) SetMaxItemNumber(n int) {
	b.maxItemNumber = n
}

func (b *Benchmark) PrintReport() {
	if b.errId != "" {
		println("Error found in instance:", b.errId)
		return
	}
	unsolvedNum := len(b.answers) - b.solvedNum
	println("instances solved:", b.solvedNum, " instances unsolved:", unsolvedNum)
	fmt.Printf("total time cost: %.2f\n", b.timeCost)
}

func (b *Benchmark) PrintInstance(instanceId string) {
	println("instance id:", instanceId)
	ins := new(Instance)
	*ins = b.instances[instanceId]
	ins.Print()
}
