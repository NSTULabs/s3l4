package date

var monthDays = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d *Date) IsValid() bool {
	if d.Year%4 == 0 && d.Month == 2 && d.Day == 29 {
		return true
	}
	return d.Month >= 1 && d.Month <= 12 && d.Day >= 0 && d.Day <= monthDays[d.Month-1]
}
