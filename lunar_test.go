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
		{
			name:     "Test hour 5 (Mão) for 2023-06-20 - auspicious",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     5,
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

func TestThoTuHour(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Ngày Dậu (2023-06-20) - Thọ Tử ở giờ Tỵ",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			expected: "Tị",
		},
		{
			name:     "Ngày Thìn (2023-03-23) - Thọ Tử ở giờ Tỵ",
			date:     time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
			expected: "Tị",
		},
		{
			name:     "Ngày Sửu (2004-11-30) - Thọ Tử ở giờ Ngọ",
			date:     time.Date(2004, time.November, 30, 0, 0, 0, 0, VietnamLocation()),
			expected: "Ngọ",
		},
		{
			name:     "Ngày Dần (2030-02-12) - Thọ Tử ở giờ Ngọ",
			date:     time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
			expected: "Ngọ",
		},
		{
			name:     "Ngày Dậu (1997-01-19) - Thọ Tử ở giờ Tị",
			date:     time.Date(1997, time.January, 19, 0, 0, 0, 0, VietnamLocation()),
			expected: "Tị",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.ThoTuHour()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestIsThoTuHour(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		hour     int
		expected bool
	}{
		{
			name:     "Ngày Dậu (2023-06-20) - giờ 10 (Tỵ) là Thọ Tử",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     10,
			expected: true,
		},
		{
			name:     "Ngày Dậu (2023-06-20) - giờ 12 (Ngọ) không phải Thọ Tử",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     12,
			expected: false,
		},
		{
			name:     "Ngày Thìn (2023-03-23) - giờ 9 (Tỵ) là Thọ Tử",
			date:     time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
			hour:     9,
			expected: true,
		},
		{
			name:     "Ngày Sửu (2004-11-30) - giờ 11 (Ngọ) là Thọ Tử",
			date:     time.Date(2004, time.November, 30, 0, 0, 0, 0, VietnamLocation()),
			hour:     11,
			expected: true,
		},
		{
			name:     "Ngày Dần (2030-02-12) - giờ 12 (Ngọ) là Thọ Tử",
			date:     time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
			hour:     12,
			expected: true,
		},
		{
			name:     "Ngày Dần (2030-02-12) - giờ 23 (Tý) không phải Thọ Tử",
			date:     time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
			hour:     23,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.IsThoTuHour(tc.hour)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSatChuHour(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Tháng 1 - Sát Chủ ở giờ Dần",
			date:     time.Date(2025, 1, 29, 0, 0, 0, 0, VietnamLocation()),
			expected: "Dần",
		},
		{
			name:     "Tháng 2 - Sát Chủ ở giờ Tỵ",
			date:     time.Date(2025, 3, 1, 0, 0, 0, 0, VietnamLocation()),
			expected: "Tị",
		},
		{
			name:     "Tháng 2 âm lịch (2023-03-23 dương) - Sát Chủ ở giờ Tị",
			date:     time.Date(2023, 3, 23, 0, 0, 0, 0, VietnamLocation()),
			expected: "Tị",
		},
		{
			name:     "Tháng 5 - Sát Chủ ở giờ Dậu",
			date:     time.Date(2023, 6, 20, 0, 0, 0, 0, VietnamLocation()),
			expected: "Dậu",
		},
		{
			name:     "Tháng 10 - Sát Chủ ở giờ Thìn",
			date:     time.Date(2004, 11, 30, 0, 0, 0, 0, VietnamLocation()),
			expected: "Thìn",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.SatChuHour()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestIsSatChuHour(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		hour     int
		expected bool
	}{
		{
			name:     "Tháng 1 - giờ 4 (Dần) là Sát Chủ",
			date:     time.Date(2025, 1, 29, 0, 0, 0, 0, VietnamLocation()),
			hour:     4,
			expected: true,
		},
		{
			name:     "Tháng 1 - giờ 12 (Ngọ) không phải Sát Chủ",
			date:     time.Date(2025, 1, 29, 0, 0, 0, 0, VietnamLocation()),
			hour:     12,
			expected: false,
		},
		{
			name:     "Tháng 2 - giờ 10 (Tỵ) là Sát Chủ",
			date:     time.Date(2025, 3, 1, 0, 0, 0, 0, VietnamLocation()),
			hour:     10,
			expected: true,
		},
		{
			name:     "Tháng 2 âm lịch (2023-03-23 dương) - giờ 10 (Tỵ) là Sát Chủ",
			date:     time.Date(2023, 3, 23, 0, 0, 0, 0, VietnamLocation()),
			hour:     10,
			expected: true,
		},
		{
			name:     "Tháng 5 - giờ 18 (Dậu) là Sát Chủ",
			date:     time.Date(2023, 6, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     18,
			expected: true,
		},
		{
			name:     "Tháng 5 - giờ 23 (Tý) không phải Sát Chủ",
			date:     time.Date(2023, 6, 20, 0, 0, 0, 0, VietnamLocation()),
			hour:     23,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.IsSatChuHour(tc.hour)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Thứ Hai - 2023-06-20",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			expected: "Thứ Hai",
		},
		{
			name:     "Thứ Tư - 2023-03-23",
			date:     time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
			expected: "Thứ Tư",
		},
		{
			name:     "Thứ Bảy - 2004-11-28",
			date:     time.Date(2004, time.November, 28, 0, 0, 0, 0, VietnamLocation()),
			expected: "Thứ Bảy",
		},
		{
			name:     "Thứ Sáu - 1997-01-18",
			date:     time.Date(1997, time.January, 18, 0, 0, 0, 0, VietnamLocation()),
			expected: "Thứ Sáu",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.Weekday()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestHourAlias(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Giờ Tý - 23h",
			date:     time.Date(2023, time.June, 20, 23, 30, 0, 0, VietnamLocation()),
			expected: "Tý",
		},
		{
			name:     "Giờ Ngọ - 12h",
			date:     time.Date(2023, time.June, 20, 12, 0, 0, 0, VietnamLocation()),
			expected: "Ngọ",
		},
		{
			name:     "Giờ Mão - 5h",
			date:     time.Date(2023, time.June, 20, 5, 15, 0, 0, VietnamLocation()),
			expected: "Mão",
		},
		{
			name:     "Giờ Dậu - 18h",
			date:     time.Date(2023, time.June, 20, 18, 45, 0, 0, VietnamLocation()),
			expected: "Dậu",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.HourAlias()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestToSolar(t *testing.T) {
	testCases := []struct {
		name          string
		date          time.Time
		expectedDay   int
		expectedMonth int
		expectedYear  int
	}{
		{
			name:          "Chuyển đổi ngày 2023-06-20",
			date:          time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			expectedDay:   20,
			expectedMonth: 6,
			expectedYear:  2023,
		},
		{
			name:          "Chuyển đổi ngày 2023-03-23",
			date:          time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
			expectedDay:   23,
			expectedMonth: 3,
			expectedYear:  2023,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			solar := l.ToSolar()
			if solar.Day != tc.expectedDay || solar.Month != tc.expectedMonth || solar.Year != tc.expectedYear {
				t.Errorf("Expected %d/%d/%d, got %d/%d/%d",
					tc.expectedDay, tc.expectedMonth, tc.expectedYear,
					solar.Day, solar.Month, solar.Year)
			}
		})
	}
}

func TestLunarString(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Test String format for 2023-06-20",
			date:     time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
			expected: "ngày Kỷ Dậu, tháng Mậu Ngọ, năm Quý Mão",
		},
		{
			name:     "Test String format for 2023-03-23 (tháng nhuận)",
			date:     time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
			expected: "ngày Canh Thìn, tháng Ất Mão Nhuận, năm Quý Mão",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.String()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
