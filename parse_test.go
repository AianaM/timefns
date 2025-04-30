package timefns

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseDateOnly(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "standard date",
			args: args{s: "2023-04-25"},
			want: time.Date(2023, 4, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "leap year date",
			args: args{s: "2024-02-29"},
			want: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "first day of year",
			args: args{s: "2025-01-01"},
			want: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "last day of year",
			args: args{s: "2022-12-31"},
			want: time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDateOnly(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDateOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "UTC time",
			args: args{s: "2023-04-25T12:30:45Z"},
			want: time.Date(2023, 4, 25, 12, 30, 45, 0, time.UTC),
		},
		{
			name: "time with positive offset",
			args: args{s: "2023-04-25T15:30:00+03:00"},
			want: time.Date(2023, 4, 25, 15, 30, 0, 0, time.FixedZone("", 3*60*60)),
		},
		{
			name: "time with negative offset",
			args: args{s: "2023-04-25T08:45:30-05:00"},
			want: time.Date(2023, 4, 25, 8, 45, 30, 0, time.FixedZone("", -5*60*60)),
		},
		{
			name: "local time with zero seconds",
			args: args{s: "2025-12-31T23:59:00Z"},
			want: time.Date(2025, 12, 31, 23, 59, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
