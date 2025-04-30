package timefns

import "time"

func CurrentWeek() TimeSpan {
	return Week(time.Now())
}

func Week(t time.Time) TimeSpan {
	t = StartOfDay(t)
	offset := (t.Weekday() + 7 - 1) % 7

	monday := t.AddDate(0, 0, -int(offset))
	saturday := monday.AddDate(0, 0, 5)

	return NewTimeSpan(monday, saturday)
}
