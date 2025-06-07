package timefns

import "fmt"

func ParseTimeSpan(start, end string) (TimeSpan, error) {
	startDate, err := Parse(start)
	if err != nil {
		return TimeSpan{}, fmt.Errorf("error parsing start date: %w", err)
	}
	endDate, err := Parse(end)
	if err != nil {
		return TimeSpan{}, fmt.Errorf("error parsing end date: %w", err)
	}
	return TimeSpan{Start: startDate, End: endDate}, nil
}
