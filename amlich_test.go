package amlich

import (
	"testing"
)

func TestSolar2Lunar(t *testing.T) {
	testCases := []struct {
		Solar []int
		Lunar []int
	}{
		{
			Solar: []int{23, 2, 2023, 7},
			Lunar: []int{4, 2, 2023, 0},
		},
		{
			Solar: []int{19, 1, 2023, 7},
			Lunar: []int{28, 12, 2022, 0},
		},
		{
			Solar: []int{31, 12, 2022, 7},
			Lunar: []int{9, 12, 2022, 0},
		},
	}

	for _, v := range testCases {
		d, m, y, leap := Solar2Lunar(v.Solar[0], v.Solar[1], v.Solar[2], v.Solar[3])

		l := v.Lunar[3] == 1
		if d != v.Lunar[0] || m != v.Lunar[1] || y != v.Lunar[2] || l != leap {
			t.Errorf("Failed test case: %v\nExpect: %v\nGot: %v %v %v %v", v.Solar, v.Lunar, d, m, y, leap)
		}
	}
}
