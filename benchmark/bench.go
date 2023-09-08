package benchmark

import (
	"fmt"
	"time"
)

type Benchmark struct {
	Name       string
	startTime  time.Time
	finishTime time.Time
}

func (task *Benchmark) Start() {
	task.startTime = time.Now()
	fmt.Printf("[%s] Benchmark iniciado em %s\n", task.Name, task.startTime.Format("2006-01-02 15:04:05"))
}

func (task *Benchmark) Finish() {
	task.finishTime = time.Now()
	fmt.Printf("[%s] Benchmark concluído em %s\n", task.Name, task.finishTime.Format("2006-01-02 15:04:05"))
}
