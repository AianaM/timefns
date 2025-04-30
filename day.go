package timefns

import "time"

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Today() TimeSpan {
	return Day(time.Now())
}

func Day(t time.Time) TimeSpan {
	t = StartOfDay(t)
	return NewTimeSpan(t, t.AddDate(0, 0, 1))
}
