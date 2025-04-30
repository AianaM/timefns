package timefns

import (
	"reflect"
	"testing"
	"time"
)

func Test_currentMonth(t *testing.T) {
	// Since currentMonth() uses time.Now() which changes with each call,
	// we test invariants rather than exact values
	got := CurrentMonth()

	// Test start is first day of month at midnight
	if got.Start.Day() != 1 || got.Start.Hour() != 0 || got.Start.Minute() != 0 ||
		got.Start.Second() != 0 || got.Start.Nanosecond() != 0 {
		t.Errorf("currentMonth() start is not midnight of first day: %v", got.Start)
	}

	// Test end is first day of next month at midnight
	expectedEndMonth := got.Start.Month() + 1
	expectedEndYear := got.Start.Year()
	// Handle December -> January transition
	if expectedEndMonth > 12 {
		expectedEndMonth = 1
		expectedEndYear++
	}

	if got.End.Day() != 1 || int(got.End.Month()) != int(expectedEndMonth) ||
		got.End.Year() != expectedEndYear || got.End.Hour() != 0 ||
		got.End.Minute() != 0 || got.End.Second() != 0 || got.End.Nanosecond() != 0 {
		t.Errorf("currentMonth() end is not midnight of first day of next month: %v", got.End)
	}

	// Test that the returned month matches current month
	now := time.Now()
	if got.Start.Year() != now.Year() || got.Start.Month() != now.Month() {
		t.Errorf("currentMonth() returns incorrect month: got %v-%v, expected %v-%v",
			got.Start.Year(), got.Start.Month(), now.Year(), now.Month())
	}
}

func Test_month(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "Middle of month",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "First day of month",
			args: args{t: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Last day of month",
			args: args{t: time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "December (year boundary)",
			args: args{t: time.Date(2023, 12, 15, 12, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "February in non-leap year",
			args: args{t: time.Date(2023, 2, 15, 0, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "February in leap year",
			args: args{t: time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Different timezone",
			args: args{t: time.Date(2023, 5, 15, 14, 30, 45, 0, time.FixedZone("EST", -5*60*60))},
			want: TimeSpan{
				Start: time.Date(2023, 5, 1, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
				End:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Month(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("month() = %v, want %v", got, tt.want)
			}
		})
	}
}
