package timefns

import (
	"testing"
	"time"
)

func TestStartOfDay(t *testing.T) {
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
			args: args{t: time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)},
			want: time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "afternoon becomes midnight",
			args: args{t: time.Date(2023, 4, 15, 15, 30, 45, 500, time.UTC)},
			want: time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfDay(tt.args.t); !got.Equal(tt.want) {
				t.Errorf("StartOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "day span",
			args: args{t: time.Date(2023, 4, 15, 15, 30, 45, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 4, 16, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Day(tt.args.t)
			if !got.Start.Equal(tt.want.Start) || !got.End.Equal(tt.want.End) {
				t.Errorf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}
