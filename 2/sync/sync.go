package sync

import "test/date"

func Run(dates []date.Date) []bool {
	validated := make([]bool, len(dates))

	for j := 0; j < len(dates); j++ {
		if dates[j].IsValid() {
			validated[j] = true
		}
	}
	return validated
}
