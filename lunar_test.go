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
