package main

import (
	"fmt"
	"time"
	"runtime"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("%d cpus will be used\n", runtime.GOMAXPROCS(0))

	memStats := &runtime.MemStats{}
	start := time.Now()
	loops := 999999
	fmt.Println("go routines concurency will be used")
	for i := 0; i < loops; i++ {
		go Sqrt(float64(i))
	}
	runtime.ReadMemStats(memStats)
	fmt.Printf("%d loops in %s\n", loops, time.Since(start))
	fmt.Println("======")
	fmt.Printf("%dkb/%dkb\n", memStats.Alloc/1024, memStats.TotalAlloc/1024)
	fmt.Printf("%dkb heap alloc\n", memStats.HeapAlloc/1024)
	fmt.Println("======")
}
