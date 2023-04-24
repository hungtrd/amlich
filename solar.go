package amlich

import (
	"fmt"
	"time"
)

type Solar struct {
	t     time.Time
	Day   int
	Month int
	Year  int
}

func (s *Solar) String() string {
	return fmt.Sprintf("%s, ngày %02d, tháng %02d, năm %d", s.Weekday(), s.Day, s.Month, s.Year)
}

func (s *Solar) Weekday() string {
	wd := int(s.t.Weekday())
	return DaysOfWeek[wd]
}
