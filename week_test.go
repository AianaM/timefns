package timefns

import (
	"reflect"
	"testing"
	"time"
)

func Test_currentWeek(t *testing.T) {
	// Since currentWeek() uses time.Now() which changes with each call,
	// we test invariants rather than exact values
	got := CurrentWeek()

	// Test start is Monday at midnight
	if got.Start.Weekday() != time.Monday || got.Start.Hour() != 0 ||
		got.Start.Minute() != 0 || got.Start.Second() != 0 || got.Start.Nanosecond() != 0 {
		t.Errorf("currentWeek() start is not midnight of Monday: %v", got.Start)
	}

	// Test end is Saturday at midnight
	if got.End.Weekday() != time.Saturday || got.End.Hour() != 0 ||
		got.End.Minute() != 0 || got.End.Second() != 0 || got.End.Nanosecond() != 0 {
		t.Errorf("currentWeek() end is not midnight of Saturday: %v", got.End)
	}

	// Test duration is 5 days (Monday to Saturday)
	expectedDuration := 5 * 24 * time.Hour
	gotDuration := got.End.Sub(got.Start)
	if gotDuration != expectedDuration {
		t.Errorf("currentWeek() duration = %v, want %v", gotDuration, expectedDuration)
	}

	// Test that the returned week contains current day
	now := time.Now()
	nowStartOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if got.Start.After(nowStartOfDay) || got.End.Before(nowStartOfDay) {
		t.Errorf("currentWeek() does not contain current day. Week: %v to %v, today: %v",
			got.Start, got.End, nowStartOfDay)
	}
}

func Test_week(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "Monday",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC)}, // Monday
			want: TimeSpan{
				Start: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), // Monday
				End:   time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC), // Saturday
			},
		},
		{
			name: "Wednesday",
			args: args{t: time.Date(2023, 5, 17, 10, 0, 0, 0, time.UTC)}, // Wednesday
			want: TimeSpan{
				Start: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), // Monday
				End:   time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC), // Saturday
			},
		},
		{
			name: "Sunday",
			args: args{t: time.Date(2023, 5, 14, 23, 59, 59, 0, time.UTC)}, // Sunday
			want: TimeSpan{
				Start: time.Date(2023, 5, 8, 0, 0, 0, 0, time.UTC),  // Monday
				End:   time.Date(2023, 5, 13, 0, 0, 0, 0, time.UTC), // Saturday
			},
		},
		{
			name: "Week spanning month boundary",
			args: args{t: time.Date(2023, 5, 31, 12, 0, 0, 0, time.UTC)}, // Wednesday, May 31
			want: TimeSpan{
				Start: time.Date(2023, 5, 29, 0, 0, 0, 0, time.UTC), // Monday, May 29
				End:   time.Date(2023, 6, 3, 0, 0, 0, 0, time.UTC),  // Saturday, June 3
			},
		},
		{
			name: "Week spanning year boundary",
			args: args{t: time.Date(2023, 12, 31, 12, 0, 0, 0, time.UTC)}, // Sunday, Dec 31
			want: TimeSpan{
				Start: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC), // Monday, Dec 25
				End:   time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC), // Saturday, Dec 30
			},
		},
		{
			name: "Different timezone",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 0, time.FixedZone("EST", -5*60*60))}, // Monday
			want: TimeSpan{
				Start: time.Date(2023, 5, 15, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)), // Monday
				End:   time.Date(2023, 5, 20, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)), // Saturday
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Week(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("week() = %v, want %v", got, tt.want)
			}
		})
	}
}
