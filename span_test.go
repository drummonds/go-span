package span_test

import (
	"math"
	"strconv"
	"testing"
	"time"

	span "git.bytestone.uk/hum3/go-span"
)

const (
	day  = 24 * time.Hour
	year = 365 * day
)

func TestFormat(t *testing.T) {
	cases := []struct {
		d    time.Duration
		want string
	}{
		// Below a microsecond: whole nanoseconds, the finest unit.
		{0, "0s"},
		{162 * time.Nanosecond, "162ns"},
		{999 * time.Nanosecond, "999ns"},
		// Microseconds to seconds: four significant figures.
		{16234 * time.Nanosecond, "16.23µs"},
		{162345 * time.Nanosecond, "162.3µs"},
		{162345 * time.Microsecond, "162.3ms"},
		{1623456 * time.Microsecond, "1.623s"},
		{16234567 * time.Microsecond, "16.23s"},
		{1600 * time.Millisecond, "1.6s"},
		// Minutes: one decimal on the seconds.
		{2*time.Minute + 3456*time.Millisecond, "2m3.5s"},
		// Hours: whole seconds.
		{time.Hour + 46*time.Minute + 51360895289*time.Nanosecond, "1h46m51s"},
		// Days: hours and minutes, then hours only from ten days.
		{day + 4*time.Hour + 32*time.Minute + 40*time.Second, "1d4h33m"},
		{12*day + 4*time.Hour + 20*time.Minute, "12d4h"},
		// Years of 365 days: days only.
		{year + 23*day + 5*time.Hour, "1y23d"},
		// Rounding up carries into the next, coarser tier.
		{999999 * time.Nanosecond, "1ms"},
		{59996 * time.Millisecond, "1m0s"},
		{23*time.Hour + 59*time.Minute + 59600*time.Millisecond, "1d0h0m"},
		// A negative span is the positive one with a sign.
		{-1623456 * time.Microsecond, "-1.623s"},
	}
	for _, c := range cases {
		if got := span.Format(c.d); got != c.want {
			t.Errorf("Format(%d) = %q, want %q", int64(c.d), got, c.want)
		}
	}
}

// Below a day the output is a Go duration string, so it round-trips
// through time.ParseDuration to the rounded span.
func TestFormatParsesBelowADay(t *testing.T) {
	for _, d := range []time.Duration{162, 16234, 162345678, 16234567890, 123456789012, 5432109876543} {
		if _, err := time.ParseDuration(span.Format(d)); err != nil {
			t.Errorf("Format(%d) = %q does not parse: %v", int64(d), span.Format(d), err)
		}
	}
}

func TestDays(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{0, "0d"},
		{340, "340d"},
		{365, "1y0d"},
		{388, "1y23d"},
		{-3, "-3d"},
	}
	for _, c := range cases {
		if got := span.Days(c.n); got != c.want {
			t.Errorf("Days(%d) = %q, want %q", c.n, got, c.want)
		}
	}
}

// The largest spans round to the nearest day even where that day is past
// the maximum Duration. The minimum signed value has no positive
// counterpart, so negating it overflows; it still shows as its magnitude
// with a sign.
func TestExtremes(t *testing.T) {
	if got, want := span.Format(math.MaxInt64), "292y172d"; got != want {
		t.Errorf("Format(MaxInt64) = %q, want %q", got, want)
	}
	if got, want := span.Format(math.MinInt64), "-292y172d"; got != want {
		t.Errorf("Format(MinInt64) = %q, want %q", got, want)
	}
	if strconv.IntSize != 64 {
		t.Skip("Days expectation is for 64-bit int")
	}
	if got, want := span.Days(math.MinInt), "-25269512429739111y293d"; got != want {
		t.Errorf("Days(MinInt) = %q, want %q", got, want)
	}
}
