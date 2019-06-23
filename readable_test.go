package readable

import (
	"testing"
	"time"
)

func TestToReadableSize(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"KB", args{2160}, "2.11 KB"},
		{"KB", args{2048}, "2.00 KB"},
		{"MB", args{1024 * 2048}, "2.00 MB"},
		{"MB", args{2212495}, "2.11 MB"},
		{"GB", args{1024 * 1024 * 2048}, "2.00 GB"},
		{"TB", args{1024 * 1024 * 1024 * 2048}, "2.00 TB"},
		{"Bytes", args{211}, "211 Bytes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToReadableSize(tt.args.size); got != tt.want {
				t.Errorf("ToReadableSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToReadableTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ago", args{time.Now().Add(-30 * time.Second)}, "刚刚"},
		{"minutes ago", args{time.Now().Add(-30 * time.Minute)}, "30分钟前"},
		{"hours ago", args{time.Now().Add(-2 * time.Hour)}, "2小时前"},
		{"days ago", args{time.Now().Add(-50 * time.Hour)}, "2天前"},
		{"long long ago", args{time.Date(1982, 2, 11, 5, 20, 0, 0, time.Local)}, "1982-02-11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToReadableTime(tt.args.t); got != tt.want {
				t.Errorf("ToReadableTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
