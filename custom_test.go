package timefns

import (
	"testing"
	"time"
)

func TestParseTimeSpan(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		args    args
		want    TimeSpan
		wantErr bool
	}{
		{
			name: "valid timespan with Z format",
			args: args{
				start: "2023-04-15T10:00:00Z",
				end:   "2023-04-15T18:00:00Z",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 4, 15, 18, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "valid timespan with numeric timezone",
			args: args{
				start: "2023-04-15T10:00:00.000-0700",
				end:   "2023-04-15T18:00:00.000-0700",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 10, 0, 0, 0, time.FixedZone("-0700", -7*60*60)),
				End:   time.Date(2023, 4, 15, 18, 0, 0, 0, time.FixedZone("-0700", -7*60*60)),
			},
			wantErr: false,
		},
		{
			name: "valid timespan across days",
			args: args{
				start: "2023-04-15T22:00:00Z",
				end:   "2023-04-16T06:00:00Z",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 22, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 4, 16, 6, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "mixed timezone formats",
			args: args{
				start: "2023-04-15T10:00:00Z",
				end:   "2023-04-15T18:00:00.000+0200",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 10, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 4, 15, 18, 0, 0, 0, time.FixedZone("+0200", 2*60*60)),
			},
			wantErr: false,
		},
		{
			name: "invalid start date format",
			args: args{
				start: "2023-04-15 10:00:00",
				end:   "2023-04-15T18:00:00Z",
			},
			want:    TimeSpan{},
			wantErr: true,
		},
		{
			name: "invalid end date format",
			args: args{
				start: "2023-04-15T10:00:00Z",
				end:   "not-a-date",
			},
			want:    TimeSpan{},
			wantErr: true,
		},
		{
			name: "both dates invalid",
			args: args{
				start: "invalid-start",
				end:   "invalid-end",
			},
			want:    TimeSpan{},
			wantErr: true,
		},
		{
			name: "empty start date",
			args: args{
				start: "",
				end:   "2023-04-15T18:00:00Z",
			},
			want:    TimeSpan{},
			wantErr: true,
		},
		{
			name: "empty end date",
			args: args{
				start: "2023-04-15T10:00:00Z",
				end:   "",
			},
			want:    TimeSpan{},
			wantErr: true,
		},
		{
			name: "same start and end time",
			args: args{
				start: "2023-04-15T12:00:00Z",
				end:   "2023-04-15T12:00:00Z",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "end before start (logically invalid but should parse)",
			args: args{
				start: "2023-04-15T18:00:00Z",
				end:   "2023-04-15T10:00:00Z",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 18, 0, 0, 0, time.UTC),
				End:   time.Date(2023, 4, 15, 10, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "UTC timezone with +0000",
			args: args{
				start: "2023-04-15T10:00:00.000+0000",
				end:   "2023-04-15T18:00:00.000+0000",
			},
			want: TimeSpan{
				Start: time.Date(2023, 4, 15, 10, 0, 0, 0, time.FixedZone("+0000", 0)),
				End:   time.Date(2023, 4, 15, 18, 0, 0, 0, time.FixedZone("+0000", 0)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTimeSpan(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTimeSpan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !got.Start.Equal(tt.want.Start) || !got.End.Equal(tt.want.End) {
					t.Errorf("ParseTimeSpan() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
