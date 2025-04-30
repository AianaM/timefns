package timefns

import (
	"time"
)

type TimeSpan struct {
	Start, End time.Time
}

func NewTimeSpan(start, end time.Time) TimeSpan {
	return TimeSpan{start, end}
}
