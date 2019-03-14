package main

import (
	"fmt"
)

func fib(n uint64) uint64 {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	var sequence []uint64
	var n uint64

	for n = 0; n <= 40; n++ {
		sequence = append(sequence, fib(n))
	}
	fmt.Println(sequence)
}
