package timefns

import (
	"reflect"
	"testing"
	"time"
)

func Test_startOfDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "midnight stays the same",
			args: args{t: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)},
			want: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "afternoon truncates to midnight",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 123456789, time.UTC)},
			want: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "preserves timezone",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 123456789, time.FixedZone("EST", -5*60*60))},
			want: time.Date(2023, 5, 15, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfDay(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("startOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_today(t *testing.T) {
	// Since today() uses time.Now() which changes with each call,
	// we can only test certain invariants rather than exact values
	got := Today()

	// Test that the timespan returned is exactly 24 hours
	wantDuration := 24 * time.Hour
	gotDuration := got.End.Sub(got.Start)
	if gotDuration != wantDuration {
		t.Errorf("today() duration = %v, want %v", gotDuration, wantDuration)
	}

	// Test that start time is at midnight (hour, min, sec, nsec all zero)
	if got.Start.Hour() != 0 || got.Start.Minute() != 0 || got.Start.Second() != 0 || got.Start.Nanosecond() != 0 {
		t.Errorf("today() start time is not midnight: %v", got.Start)
	}

	// Test that the returned day matches today's date
	now := time.Now()
	if got.Start.Year() != now.Year() || got.Start.Month() != now.Month() || got.Start.Day() != now.Day() {
		t.Errorf("today() returns incorrect date: got %v, expected date matching %v", got.Start, now)
	}
}

func Test_day(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "UTC midnight",
			args: args{t: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 5, 16, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "UTC afternoon",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 123456789, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 5, 16, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Different timezone",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 123456789, time.FixedZone("EST", -5*60*60))},
			want: TimeSpan{
				Start: time.Date(2023, 5, 15, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
				End:   time.Date(2023, 5, 16, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
			},
		},
		{
			name: "Month boundary",
			args: args{t: time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 5, 31, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Year boundary",
			args: args{t: time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Leap year",
			args: args{t: time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day() = %v, want %v", got, tt.want)
			}
		})
	}
}
