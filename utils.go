package amlich

import "time"

// b2i converts a boolean value to an integer value.
func b2i(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue #6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

// getTz computes offset in hours east of UTC
// e.g. UTC+7 returns 7, UTC-7 returns -7
func getTz(t time.Time) int {
	_, offset := t.Zone()
	tz := offset / 3600

	return tz
}

// loc2tz computes offset in hours east of UTC from time.Location
// e.g. UTC+7 returns 7, UTC-7 returns -7
func loc2tz(l time.Location) int {
	t := time.Now().In(&l)
	return getTz(t)
}

// VietnamLocation returns the *time.Location for Vietnam
func VietnamLocation() *time.Location {
	loc, _ := time.LoadLocation(VietnamTimezone)
	return loc
}

func getHourChi(hour int) int {
	if hour >= 23 || hour < 1 {
		return 0
	}
	if hour >= 1 && hour < 3 {
		return 1
	}
	if hour >= 3 && hour < 5 {
		return 2
	}
	if hour >= 5 && hour < 7 {
		return 3
	}
	if hour >= 7 && hour < 9 {
		return 4
	}
	if hour >= 9 && hour < 11 {
		return 5
	}
	if hour >= 11 && hour < 13 {
		return 6
	}
	if hour >= 13 && hour < 15 {
		return 7
	}
	if hour >= 15 && hour < 17 {
		return 8
	}
	if hour >= 17 && hour < 19 {
		return 9
	}
	if hour >= 19 && hour < 21 {
		return 10
	}
	return 11
}
