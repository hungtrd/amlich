package amlich

import (
	"math"
)

const pi float64 = 3.141592653589793

func i2f(i int) float64 {
	return float64(i)
}

func floor(f float64) int {
	return int(math.Floor(f))
}

func date2JuliusDay(dd, mm, yy int) int {
	a := floor((14 - i2f(mm)) / 12)
	y := i2f(yy) + 4800 - i2f(a)
	m := i2f(mm) + 12*i2f(a) - 3
	jd := i2f(dd) + math.Floor((153*m+2)/5) + 365*y + math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400) - 32045
	if jd < 2299161 {
		jd = i2f(dd) + math.Floor((153*m+2)/5) + 365*y + math.Floor(y/4) - 32083
	}

	return floor(jd)
}

func juliusDay2Date(jd int) (day, month, year int) {
	var a, b, c, d, e, m int
	if jd > 2299160 {
		a = jd + 32044
		b = (4*a + 3) / 146097
		c = a - (b*146097)/4
	} else {
		b = 0
		c = jd + 32082
	}

	d = (4*c + 3) / 1461
	e = c - (1461*d)/4
	m = (5*e + 2) / 153

	day = e - (153*m+2)/5 + 1
	month = m + 3 - 12*(m/10)
	year = b*100 + d - 4800 + (m / 10)

	return
}

func newMoon(k int) float64 {
	var kf, t, t2, t3, dr, jd1, m, mpr, f, c1, deltat, jdNew float64
	kf = i2f(k)
	t = kf / 1236.85 // Time in Julian centuries from 1900 January 0.5
	t2 = t * t
	t3 = t2 * t
	dr = pi / 180
	jd1 = 2415020.75933 + 29.53058868*kf + 0.0001178*t2 - 0.000000155*t3
	jd1 = jd1 + 0.00033*math.Sin((166.56+132.87*t-0.009173*t2)*dr)  // Mean new moon
	m = 359.2242 + 29.10535608*kf - 0.0000333*t2 - 0.00000347*t3    // Sun's mean anomaly
	mpr = 306.0253 + 385.81691806*kf + 0.0107306*t2 + 0.00001236*t3 // Moon's mean anomaly
	f = 21.2964 + 390.67050646*kf - 0.0016528*t2 - 0.00000239*t3    // Moon's argument of latitude
	c1 = (0.1734-0.000393*t)*math.Sin(m*dr) + 0.0021*math.Sin(2*dr*m)
	c1 = c1 - 0.4068*math.Sin(mpr*dr) + 0.0161*math.Sin(dr*2*mpr)
	c1 = c1 - 0.0004*math.Sin(dr*3*mpr)
	c1 = c1 + 0.0104*math.Sin(dr*2*f) - 0.0051*math.Sin(dr*(m+mpr))
	c1 = c1 - 0.0074*math.Sin(dr*(m-mpr)) + 0.0004*math.Sin(dr*(2*f+m))
	c1 = c1 - 0.0004*math.Sin(dr*(2*f-m)) - 0.0006*math.Sin(dr*(2*f+mpr))
	c1 = c1 + 0.0010*math.Sin(dr*(2*f-mpr)) + 0.0005*math.Sin(dr*(2*mpr+m))
	if t < -11 {
		deltat = 0.001 + 0.000839*t + 0.0002261*t2 - 0.00000845*t3 - 0.000000081*t*t3
	} else {
		deltat = -0.000278 + 0.000265*t + 0.000262*t2
	}
	jdNew = jd1 + c1 - deltat

	return jdNew
}

func getNewMoonDay(k, tz int) int {
	return floor(newMoon(k) + 0.5 + i2f(tz)/24)
}

func sunLongitude(jdn float64) float64 {
	var t, t2, dr, m, l0, dl, l float64
	t = (jdn - 2451545.0) / 36525 // Time in Julian centuries from 2000-01-01 12:00:00 GMT
	t2 = t * t
	dr = pi / 180                                                  // degree to radian
	m = 357.52910 + 35999.05030*t - 0.0001559*t2 - 0.00000048*t*t2 // mean anomaly, degree
	l0 = 280.46645 + 36000.76983*t + 0.0003032*t2                  // mean longitude, degree
	dl = (1.914600 - 0.004817*t - 0.000014*t2) * math.Sin(dr*m)
	dl = dl + (0.019993-0.000101*t)*math.Sin(dr*2*m) + 0.00029*math.Sin(dr*3*m)
	l = l0 + dl // true longitude, degree
	l = l * dr
	l = l - pi*2*(math.Floor(l/(pi*2))) // Normalize to (0, 2*PI)
	return l
}

func getSunLongitude(d, tz int) int {
	return floor((sunLongitude(i2f(d)-0.5-i2f(tz)/24) / pi) * 6)
}

/* Find the day that starts the luner month 11 of the given year for the given time zone */
func getLunarMonth11(yy, tz int) int {
	var off, k, nm int

	off = date2JuliusDay(31, 12, yy) - 2415021
	k = floor(i2f(off) / 29.530588853)
	nm = getNewMoonDay(k, tz)
	sunLong := getSunLongitude(nm, tz) // sun longitude at local midnight
	if sunLong >= 9 {
		nm = getNewMoonDay(k-1, tz)
	}
	return nm
}

func getLeapMonthOffset(a11 float64, tz int) int {
	k := floor((a11-2415021.076998695)/29.530588853 + 0.5)
	last := 0
	i := 1 // We start with the month following lunar month 11
	arc := getSunLongitude(getNewMoonDay(k+i, tz), tz)
	for {
		last = arc
		i = i + 1
		newmoon := getNewMoonDay(k+i, tz)
		arc = getSunLongitude(newmoon, tz)

		if arc == last || i >= 14 {
			break
		}
	}
	return i - 1
}

func Solar2Lunar(dd, mm, yy, tz int) (lunarD, lunarM, lunarY int, lunarLeap bool) {
	dayNumber := date2JuliusDay(dd, mm, yy)
	k := floor((i2f(dayNumber) - 2415021.076998695) / 29.530588853)
	monthStart := getNewMoonDay(k+1, tz)
	if monthStart > dayNumber {
		monthStart = getNewMoonDay(k, tz)
	}
	a11 := getLunarMonth11(yy, tz)
	b11 := a11
	if a11 >= monthStart {
		lunarY = yy
		a11 = getLunarMonth11(yy-1, tz)
	} else {
		lunarY = yy + 1
		b11 = getLunarMonth11(yy+1, tz)
	}
	lunarD = dayNumber - monthStart + 1
	diff := floor((i2f(monthStart) - i2f(a11)) / 29)
	lunarLeap = false
	lunarM = diff + 11
	if b11-a11 > 365 {
		leapMonthDiff := getLeapMonthOffset(i2f(a11), tz)
		if diff >= leapMonthDiff {
			lunarM = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = false
			}
		}
	}
	if lunarM > 12 {
		lunarM = lunarM - 12
	}
	if lunarM >= 11 && diff < 4 {
		lunarY -= 1
	}

	return
}

func Lunar2Solar(lunarDay, lunarMonth, lunarYear, lunarLeap, tz int) (d, m, y int) {
	var k, a11, b11, off, leapOff, leapMonth, monthStart int
	if lunarMonth < 11 {
		a11 = getLunarMonth11(lunarYear-1, tz)
		b11 = getLunarMonth11(lunarYear, tz)
	} else {
		a11 = getLunarMonth11(lunarYear, tz)
		b11 = getLunarMonth11(lunarYear+1, tz)
	}
	k = floor(0.5 + (i2f(a11)-2415021.076998695)/29.530588853)
	off = lunarMonth - 11
	if off < 0 {
		off += 12
	}
	if b11-a11 > 365 {
		leapOff = getLeapMonthOffset(i2f(a11), tz)
		leapMonth = leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
		if lunarLeap != 0 && lunarMonth != leapMonth {
			return
		} else if lunarLeap != 0 || off >= leapOff {
			off += 1
		}
	}
	monthStart = getNewMoonDay(k+off, tz)

	return juliusDay2Date(monthStart + lunarDay - 1)
}
