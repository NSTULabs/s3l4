package condimpl

import (
	"math/rand"
	"runtime"
	"sync"
)

type chars struct {
	mu   sync.Mutex
	cond *sync.Cond
	data []rune
	len  int
}

func fill(ch *chars, numChars int) {
	for {
		randRune := rune(rand.Intn(128))
		ch.mu.Lock()
		if ch.len == numChars {
			ch.mu.Unlock()
			return
		}
		ch.data[ch.len] = randRune
		ch.len++
		ch.cond.Signal() // Сигналим, что состояние изменилось
		ch.mu.Unlock()
	}
}

func Run(numChars int) []rune {
	gorsCount := runtime.NumCPU()
	runtime.GOMAXPROCS(gorsCount)

	ch := chars{
		data: make([]rune, numChars),
		len:  0,
	}
	ch.cond = sync.NewCond(&ch.mu) // создаем условную переменную на основе ch.mu

	// Запускаем горутины для заполнения данных
	for i := 0; i < gorsCount; i++ {
		go fill(&ch, numChars)
	}

	ch.mu.Lock()
	for ch.len != numChars {
		ch.cond.Wait()
	}
	ch.mu.Unlock()

	return ch.data
}
