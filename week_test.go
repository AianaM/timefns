package timefns

import (
	"testing"
	"time"
)

func TestWeek(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "starting on Monday",
			args: args{t: time.Date(2023, 4, 10, 12, 0, 0, 0, time.UTC)}, // Monday
			want: TimeSpan{
				Start: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC), // Monday
				End:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC), // Saturday
			},
		},
		{
			name: "starting on Wednesday",
			args: args{t: time.Date(2023, 4, 12, 12, 0, 0, 0, time.UTC)}, // Wednesday
			want: TimeSpan{
				Start: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC), // Monday
				End:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC), // Saturday
			},
		},
		{
			name: "starting on Sunday",
			args: args{t: time.Date(2023, 4, 16, 12, 0, 0, 0, time.UTC)}, // Sunday
			want: TimeSpan{
				Start: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC), // Monday
				End:   time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC), // Saturday
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Week(tt.args.t)
			if !got.Start.Equal(tt.want.Start) || !got.End.Equal(tt.want.End) {
				t.Errorf("Week() = %v, want %v", got, tt.want)
			}
		})
	}
}
