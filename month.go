package timefns

import "time"

func CurrentMonth() TimeSpan {
	return Month(time.Now())
}

func Month(t time.Time) TimeSpan {
	t = StartOfDay(t)
	start := t.AddDate(0, 0, -int(t.Day())+1)
	end := start.AddDate(0, 1, 0)
	return NewTimeSpan(start, end)
}
