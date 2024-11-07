package main

import (
	"math/rand"
	"test/async"
	"test/date"
	"test/sync"
	"testing"
)

func BenchmarkDate(b *testing.B) {
	cases := []struct {
		Name  string
		Dates []date.Date
	}{
		{
			Name:  "Count 100",
			Dates: generateRandomDates(100),
		},
		{
			Name:  "Count 10000",
			Dates: generateRandomDates(10000),
		},
		{
			Name:  "Count 1000000",
			Dates: generateRandomDates(1000000),
		},
	}

	// Async benchmark tests
	for _, c := range cases {
		b.Run(c.Name+" async", func(bb *testing.B) {
			for i := 0; i < bb.N; i++ {
				async.Run(c.Dates)
			}
		})
	}
	// Sync benchmark tests
	for _, c := range cases {
		b.Run(c.Name+" sync", func(bb *testing.B) {
			for i := 0; i < bb.N; i++ {
				sync.Run(c.Dates)
			}
		})
	}
}

func generateRandomDates(n int) []date.Date {
	result := make([]date.Date, n)
	for i := 0; i < n; i++ {
		result[i] = generateRandomDate()
	}
	return result
}

func generateRandomDate() date.Date {
	return date.Date{
		Day:   rand.Intn(50),
		Month: rand.Intn(50),
		Year:  rand.Intn(3000),
	}
}
