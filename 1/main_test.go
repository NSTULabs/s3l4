package main

import (
	"fmt"
	"test/condimpl"
	"test/muteximpl"
	"test/semaphoreimpl"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	numChars := []int{100, 10000, 1000000}
	for _, num := range numChars {
		b.Run(
			fmt.Sprintf("Count %d", num),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					muteximpl.Run(num)
				}
			},
		)
	}
}

func BenchmarkSemaphore(b *testing.B) {
	numChars := []int{100, 10000, 1000000}
	for _, num := range numChars {
		b.Run(
			fmt.Sprintf("Count %d", num),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					semaphoreimpl.Run(num)
				}
			},
		)
	}
}

func BenchmarkCond(b *testing.B) {
	numChars := []int{100, 10000, 1000000}
	for _, num := range numChars {
		b.Run(
			fmt.Sprintf("Count %d", num),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					condimpl.Run(num)
				}
			},
		)
	}
}
