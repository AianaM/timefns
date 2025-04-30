package timefns

import (
	"log"
	"time"
)

// DateOnly
func ParseDateOnly(s string) time.Time {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		log.Fatal("can't parse, date string must be parseable using the format '2006-01-02'")
	}
	return t
}

func Parse(s string) time.Time {
	t, err := time.Parse(ISO8601, s)
	if err != nil {
		log.Fatal("can't parse, date string must be parseable using the format RFC3339 '2025-04-25T21:12:51.720Z'")
	}
	return t
}
