package timefns

import (
	"reflect"
	"testing"
	"time"
)

func Test_newTimeSpan(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name string
		args args
		want TimeSpan
	}{
		{
			name: "regular case with different times",
			args: args{
				start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			want: TimeSpan{
				Start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "same start and end time",
			args: args{
				start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: TimeSpan{
				Start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "end time before start time",
			args: args{
				start: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: TimeSpan{
				Start: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "different timezones",
			args: args{
				start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2023, 1, 1, 12, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
			},
			want: TimeSpan{
				Start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 1, 1, 12, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
			},
		},
		{
			name: "different dates",
			args: args{
				start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2023, 1, 5, 10, 0, 0, 0, time.UTC),
			},
			want: TimeSpan{
				Start: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 1, 5, 10, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTimeSpan(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTimeSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}
