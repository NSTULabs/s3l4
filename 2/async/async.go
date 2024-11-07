package async

import (
	"runtime"
	"sync"
	"test/date"
)

func Run(dates []date.Date) []bool {
	gorsCount := runtime.NumCPU()
	runtime.GOMAXPROCS(gorsCount)

	var wg sync.WaitGroup

	validated := make([]bool, len(dates))
	var start, end int
	step := len(dates) / gorsCount
	wg.Add(gorsCount)
	for i := 0; i < gorsCount; i++ {
		start = i * step
		if i == gorsCount-1 {
			end = len(dates)
		} else {
			end = (i + 1) * step
		}

		go func() {
			for j := start; j < end; j++ {
				if dates[j].IsValid() {
					validated[j] = true
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return validated
}
