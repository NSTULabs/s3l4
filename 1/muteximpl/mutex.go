package muteximpl

import (
	"math/rand"
	"runtime"
	"sync"
)

type chars struct {
	mu   sync.Mutex
	data []rune
	len  int
}

func fill(ch *chars, numChars int) {
	for {
		randRune := rune(rand.Intn(128))
		ch.mu.Lock()

		if ch.len == numChars {
			ch.mu.Unlock()
			break
		}
		ch.data[ch.len] = randRune
		ch.len++

		ch.mu.Unlock()
	}
}

func Run(numChars int) []rune {
	gorsCount := runtime.NumCPU()
	runtime.GOMAXPROCS(gorsCount)
	var wg sync.WaitGroup
	ch := chars{
		data: make([]rune, numChars),
		len:  0,
	}

	wg.Add(gorsCount)
	for i := 0; i < gorsCount; i++ {
		go func() {
			fill(&ch, numChars)
			wg.Done()
		}()
	}
	wg.Wait()

	return ch.data
}
