package amlich

import (
	"testing"
	"time"
)

var lunarTestCases = []struct {
	name               string
	date               time.Time
	expectedDayAlias   string
	expectedMonthAlias string
	expectedYearAlias  string
}{
	{
		name:               "Test Case 1",
		date:               time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Kỷ Dậu",
		expectedMonthAlias: "Mậu Ngọ",
		expectedYearAlias:  "Quý Mão",
	},
	{
		name:               "Test Case 2",
		date:               time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Canh Thìn",
		expectedMonthAlias: "Ất Mão Nhuận",
		expectedYearAlias:  "Quý Mão",
	},
	{
		name:               "Test Case 3",
		date:               time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Mậu Dần",
		expectedMonthAlias: "Mậu Dần",
		expectedYearAlias:  "Canh Tuất",
	},
	{
		name:               "Test Case 4",
		date:               time.Date(2004, time.November, 30, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Quý Sửu",
		expectedMonthAlias: "Ất Hợi",
		expectedYearAlias:  "Giáp Thân",
	},
	{
		name:               "Test Case 5",
		date:               time.Date(1997, time.January, 19, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Tân Dậu",
		expectedMonthAlias: "Tân Sửu",
		expectedYearAlias:  "Bính Tý",
	},
}

func TestDayAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.DayAlias()

			if result != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedDayAlias, result)
			}
		})
	}
}

func TestMonthAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.MonthAlias()

			if result != tc.expectedMonthAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedMonthAlias, result)
			}
		})
	}
}

func TestYearAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.YearAlias()

			if result != tc.expectedYearAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedYearAlias, result)
			}
		})
	}
}

func TestAuspiciousHours(t *testing.T) {
	testCases := []struct {
		name          string
		date          time.Time
		expectedCount int
	}{
		{
			name:          "Test auspicious hours count for 2023-06-20",
			date:          time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			expectedCount: 6,
		},
		{
			name:          "Test auspicious hours count for 2023-03-23",
			date:          time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
			expectedCount: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			hours := l.AuspiciousHours()

			if len(hours) != tc.expectedCount {
				t.Errorf("Expected %d auspicious hours, got %d", tc.expectedCount, len(hours))
			}

			for _, hour := range hours {
				if hour == "" {
					t.Errorf("Found empty hour name")
				}
			}
		})
	}
}

func TestIsAuspiciousHour(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		hour     int
		expected bool
	}{
		{
			name:     "Test hour 23 (Tý) for 2023-06-20 - auspicious",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     23,
			expected: true,
		},
		{
			name:     "Test hour 2 (Sửu) for 2023-06-20 - not auspicious",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     2,
			expected: false,
		},
		{
			name:     "Test hour 9 (Tỵ) for 2023-06-20 - not auspicious",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     9,
			expected: false,
		},
		{
			name:     "Test hour 15 (Thân) for 2023-06-20 - not auspicious",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     15,
			expected: false,
		},
		{
			name:     "Test hour 17 (Dậu) for 2023-06-20 - auspicious",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     17,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.IsAuspiciousHour(tc.hour)

			if result != tc.expected {
				t.Errorf("Expected IsAuspiciousHour(%d) to be %v, got %v", tc.hour, tc.expected, result)
			}
		})
	}
}
