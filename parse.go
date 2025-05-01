package timefns

import (
	"strings"
	"time"
)

// DateOnly "2006-01-02"
func ParseDateOnly(s string) (time.Time, error) {
	return time.Parse(time.DateOnly, s)
}

func Parse(s string) (time.Time, error) {
	if strings.Contains(s, "Z") {
		return ParseISO8601(s)
	}
	return ParseISO8601n(s)
}

func ParseISO8601(s string) (time.Time, error) {
	return time.Parse(ISO8601, s)
}

// with numeric zone
func ParseISO8601n(s string) (time.Time, error) {
	return time.Parse(ISO8601n, s)
}
