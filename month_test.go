package timefns

import (
	"testing"
	"time"
)

func TestMonth(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "start of month",
			args: args{t: time.Date(2023, 4, 1, 12, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "middle of month",
			args: args{t: time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "end of month",
			args: args{t: time.Date(2023, 4, 30, 12, 0, 0, 0, time.UTC)},
			want: TimeSpan{
				Start: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Month(tt.args.t)
			if !got.Start.Equal(tt.want.Start) || !got.End.Equal(tt.want.End) {
				t.Errorf("Month() = %v, want %v", got, tt.want)
			}
		})
	}
}
