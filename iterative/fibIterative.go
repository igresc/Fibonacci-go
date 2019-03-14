package main

import (
	"fmt"
	"time"
)

func fib(n uint64) uint64 {
	var a uint64 = 0
	var b uint64 = 1
	var i uint64

	for i = 0; i <= n; i++ {
		b = b + a
		a = b - a
	}

	return a
}

func main() {
	var start = time.Now()

	var sequence []uint64
	var n uint64

	for n = 0; n <= 100000; n++ {
		sequence = append(sequence, fib(n))
	}

	fmt.Println(sequence)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
