// Package span formats how long something took only as precisely as a
// reader needs: about four significant figures, with each larger unit
// that comes in pushing out a smaller one.
package span

import (
	"fmt"
	"math"
	"time"
)

const (
	day  = 24 * time.Hour
	year = 365 * day
)

// Format renders d to its tier's precision:
//
//	below 1µs        whole nanoseconds      162ns
//	1µs to 1m        4 significant figures  16.23µs 162.3ms 1.623s 16.23s
//	1m to 1h         seconds to 0.1         2m3.5s
//	1h to 1d         whole seconds          1h46m51s
//	1d to 10d        days, hours, minutes   1d4h33m
//	10d to 1y        days, hours            12d4h
//	1y (365d) and up years, days            1y23d
//
// Below a day the result is a Go duration string, so time.ParseDuration
// reads it back. A negative span is the positive one with a sign.
func Format(d time.Duration) string {
	if d == math.MinInt64 {
		// Its magnitude does not fit; one nanosecond less rounds the same.
		d++
	}
	if d < 0 {
		return "-" + Format(-d)
	}
	r := d.Round(resolution(d))
	// Rounding up can carry r into a coarser tier; round again to that.
	r = r.Round(resolution(r))
	switch {
	case r < day:
		return r.String()
	case r < 10*day:
		return fmt.Sprintf("%dd%dh%dm", r/day, r%day/time.Hour, r%time.Hour/time.Minute)
	case r < year:
		return fmt.Sprintf("%dd%dh", r/day, r%day/time.Hour)
	default:
		// Near the maximum, Round saturates rather than rounding up past
		// it, so round to the day here, where the count cannot overflow.
		n := uint64(r / day)
		if r%day >= day/2 {
			n++
		}
		return days(n)
	}
}

// resolution is the finest unit shown for a span of d.
func resolution(d time.Duration) time.Duration {
	switch {
	case d < 10*time.Microsecond:
		return time.Nanosecond
	case d < time.Minute:
		// Four significant figures: a thousandth of d's leading power of ten.
		p := 10 * time.Microsecond
		for p*10 <= d {
			p *= 10
		}
		return p / 1000
	case d < time.Hour:
		return 100 * time.Millisecond
	case d < day:
		return time.Second
	case d < 10*day:
		return time.Minute
	case d < year:
		return time.Hour
	default:
		return day
	}
}

// Days renders a count of days, such as simulated days run, in days and
// from 365 on in years of 365 days and days. Its unit is the day, so it
// never shows hours.
func Days(n int) string {
	if n < 0 {
		// Negating the minimum int wraps, but read unsigned it is still
		// the magnitude.
		return "-" + days(uint64(-int64(n)))
	}
	return days(uint64(n))
}

func days(n uint64) string {
	if n < 365 {
		return fmt.Sprintf("%dd", n)
	}
	return fmt.Sprintf("%dy%dd", n/365, n%365)
}
