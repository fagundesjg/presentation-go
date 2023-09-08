package main

import (
	"presentation-go/benchmark"
	fib "presentation-go/fibbonati"
)

func main() {
	bench := benchmark.Benchmark{Name: "Tempo total"}
	bench.Start()
	fib.FibbonatiExample()
	bench.Finish()
}
