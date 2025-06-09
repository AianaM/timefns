package timefns

import (
	"strings"
	"time"
)

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
