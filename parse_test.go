package timefns

import (
	"testing"
	"time"
)

func TestParseDateOnly(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid date",
			args:    args{s: "2023-04-15"},
			want:    time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "invalid date format",
			args:    args{s: "15/04/2023"},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateOnly(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ParseDateOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ISO8601 with Z",
			args:    args{s: "2023-04-15T12:30:45Z"},
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "ISO8601 with numeric zone",
			args:    args{s: "2023-04-15T12:30:45.000-0700"},
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.FixedZone("", -7*60*60)),
			wantErr: false,
		},
		{
			name:    "invalid format",
			args:    args{s: "not a date"},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseISO8601(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid ISO8601",
			args:    args{s: "2023-04-15T12:30:45Z"},
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.UTC),
			wantErr: false,
		},
		{
			name: "valid ISO8601 with timezone",
			args: args{s: "2023-04-15T12:30:45+02:00"},
			// Fix: Use time.FixedZone for consistent representation
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.FixedZone("+0200", 2*60*60)),
			wantErr: false,
		},
		{
			name:    "invalid format",
			args:    args{s: "2023-04-15 12:30:45"},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseISO8601(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseISO8601() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Compare Unix timestamps instead of direct comparison
				if !got.Equal(tt.want) {
					t.Errorf("ParseISO8601() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestParseISO8601n(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid ISO8601 with numeric zone (negative)",
			args:    args{s: "2023-04-15T12:30:45.000-0700"},
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.FixedZone("-0700", -7*60*60)),
			wantErr: false,
		},
		{
			name:    "valid ISO8601 with numeric zone (positive)",
			args:    args{s: "2023-04-15T12:30:45.000+0200"},
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.FixedZone("+0200", 2*60*60)),
			wantErr: false,
		},
		{
			name:    "valid ISO8601 with numeric zone (no milliseconds)",
			args:    args{s: "2023-04-15T12:30:45-0700"},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "valid ISO8601 with UTC zone",
			args: args{s: "2023-04-15T12:30:45.000+0000"},
			// For UTC zone with numeric representation, use FixedZone with name "+0000"
			want:    time.Date(2023, 4, 15, 12, 30, 45, 0, time.FixedZone("+0000", 0)),
			wantErr: false,
		},
		{
			name:    "invalid format (missing T)",
			args:    args{s: "2023-04-15 12:30:45.000-0700"},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "invalid format (not a date)",
			args:    args{s: "not a date"},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "invalid format (Z instead of numeric)",
			args:    args{s: "2023-04-15T12:30:45.000Z"},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseISO8601n(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseISO8601n() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Compare time values using Equal method instead of reflect.DeepEqual
				if !got.Equal(tt.want) {
					t.Errorf("ParseISO8601n() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
