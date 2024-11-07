package semaphoreimpl

import (
	"math/rand"
	"runtime"
	"sync"
)

type Chars struct {
	semaphore chan struct{}
	data      []rune
	len       int
}

func fill(chars *Chars, numChars int) {
	for {
		randRune := rune(rand.Intn(128))
		chars.semaphore <- struct{}{}

		if chars.len == numChars {
			<-chars.semaphore
			break
		}
		chars.data[chars.len] = randRune
		chars.len++

		<-chars.semaphore
	}
}

func Run(numChars int) []rune {
	gorsCount := runtime.NumCPU()
	runtime.GOMAXPROCS(gorsCount)
	var wg sync.WaitGroup
	chars := Chars{
		data:      make([]rune, numChars),
		len:       0,
		semaphore: make(chan struct{}, 1),
	}

	wg.Add(gorsCount)
	for i := 0; i < gorsCount; i++ {
		go func() {
			fill(&chars, numChars)
			wg.Done()
		}()
	}
	wg.Wait()

	return chars.data
}
