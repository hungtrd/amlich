package amlich

import (
	"fmt"
	"time"
)

type Lunar struct {
	t     time.Time
	Day   int
	Month int
	Year  int
	Leap  bool
}

func (l *Lunar) String() string {
	return fmt.Sprintf("ngày %s, tháng %s, năm %s", l.DayAlias(), l.MonthAlias(), l.YearAlias())
}

func (l *Lunar) DayAlias() string {
	return fmt.Sprintf("%s %s", l.dayCan(), l.dayChi())
}

func (l *Lunar) dayCan() string {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	return Can[(jd+9)%10]
}

func (l *Lunar) dayChi() string {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	return Chi[(jd+1)%12]
}

func (l *Lunar) MonthAlias() string {
	if l.Leap {
		return fmt.Sprintf("%s %s Nhuận", l.monthCan(), l.monthChi())
	}
	return fmt.Sprintf("%s %s", l.monthCan(), l.monthChi())
}

func (l *Lunar) monthCan() string {
	i := (l.Year*12 + l.Month + 3) % 10
	return Can[i]
}

func (l *Lunar) monthChi() string {
	i := (l.Month + 1) % 12
	return Chi[i]
}

func (l *Lunar) YearAlias() string {
	return fmt.Sprintf("%s %s", l.yearCan(), l.yearChi())
}

func (l *Lunar) yearCan() string {
	i := (l.Year + 6) % 10
	return Can[i]
}

func (l *Lunar) yearChi() string {
	i := (l.Year + 8) % 12
	return Chi[i]
}

func (l *Lunar) HourAlias() string {
	return getHourDiaChi(l.t.Hour()).String()
}

func (l *Lunar) ToSolar() Solar {
	d, m, y := Lunar2Solar(l.Day, l.Month, l.Year, b2i(l.Leap), getTz(l.t))
	return Solar{
		Day:   d,
		Month: m,
		Year:  y,
	}
}

func (l *Lunar) Weekday() string {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	return DaysOfWeek[jd%7]
}

// AuspiciousHours trả về danh sách các giờ hoàng đạo trong ngày
func (l *Lunar) AuspiciousHours() []string {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	dayChi := (jd + 1) % 12

	hourIndices := AuspiciousHours[dayChi]

	hours := make([]string, 0, len(hourIndices))
	for _, idx := range hourIndices {
		hours = append(hours, HourNames[idx])
	}

	return hours
}

// IsAuspiciousHour kiểm tra xem giờ có phải là giờ hoàng đạo không
func (l *Lunar) IsAuspiciousHour(hour int) bool {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	dayChi := (jd + 1) % 12

	hourChi := getHourDiaChi(hour)

	hourIndices := AuspiciousHours[dayChi]
	for _, idx := range hourIndices {
		if DiaChi(idx) == hourChi {
			return true
		}
	}

	return false
}

// ThoTuHour trả về giờ thọ tử của ngày
func (l *Lunar) ThoTuHour() string {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	dayChi := (jd + 1) % 12
	hourChi := ThoTuHour[DiaChi(dayChi)]

	return hourChi.String()
}

// IsThoTuHour kiểm tra xem giờ có phải là giờ thọ tử không
func (l *Lunar) IsThoTuHour(hour int) bool {
	jd := date2JuliusDay(l.t.Day(), int(l.t.Month()), l.t.Year())
	dayChi := (jd + 1) % 12
	hourChi := getHourDiaChi(hour)

	return ThoTuHour[DiaChi(dayChi)] == hourChi
}

// SatChuHour trả về giờ sát chủ của tháng
func (l *Lunar) SatChuHour() string {
	hourChi := SatChuHour[l.Month]
	return hourChi.String()
}

// IsSatChuHour kiểm tra xem giờ có phải là giờ sát chủ không
func (l *Lunar) IsSatChuHour(hour int) bool {
	hourChi := getHourDiaChi(hour)
	return SatChuHour[l.Month] == hourChi
}

// TODO: implement function
func (l *Lunar) Next() *Lunar {
	return nil
}
