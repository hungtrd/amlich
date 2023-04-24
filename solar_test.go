package amlich

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	tc := []struct {
		day      []int
		expected string
	}{
		{
			day:      []int{21, 5, 2023},
			expected: "Chủ Nhật, ngày 21, tháng 05, năm 2023",
		},
	}

	for _, v := range tc {
		timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		today := time.Date(v.day[2], time.Month(v.day[1]), v.day[0], 0, 0, 0, 0, timeLoc)
		s := Solar{
			t:     today,
			Day:   today.Day(),
			Month: int(today.Month()),
			Year:  today.Year(),
		}

		if s.String() != v.expected {
			t.Errorf("Failed: %v.\nExpected: %s.\nGot: %s", v.day, v.expected, s.String())
		}
	}
}
