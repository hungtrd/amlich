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

func getHourDiaChi(hour int) DiaChi {
	switch hour {
	case 23, 0:
		return Ty
	case 1, 2:
		return Suu
	case 3, 4:
		return Dan
	case 5, 6:
		return Mao
	case 7, 8:
		return Thin
	case 9, 10:
		return Ti
	case 11, 12:
		return Ngo
	case 13, 14:
		return Mui
	case 15, 16:
		return Than
	case 17, 18:
		return Dau
	case 19, 20:
		return Tuat
	case 21, 22:
		return Hoi
	default:
		panic("hour must be in range 0..23")
	}
}
