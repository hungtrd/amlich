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
