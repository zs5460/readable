package readable

import (
	"fmt"
	"time"
)

const (
	_ = 1 << (10 * iota)
	kb
	mb
	gb
	tb
)

//ToReadableSize returns readble size string.
func ToReadableSize(size int) string {
	switch {
	case size >= tb:
		return fmt.Sprintf("%4.2f TB", float64(size)/tb)
	case size >= gb:
		return fmt.Sprintf("%4.2f GB", float64(size)/gb)
	case size >= mb:
		return fmt.Sprintf("%4.2f MB", float64(size)/mb)
	case size >= kb:
		return fmt.Sprintf("%4.2f KB", float64(size)/kb)
	default:
		return fmt.Sprintf("%d Bytes", size)
	}
}

// ToReadableTime returns easy to read format.
func ToReadableTime(t time.Time) string {
	seconds := int(time.Now().Sub(t).Seconds())
	switch {
	case seconds > 0 && seconds < 60:
		return "刚刚"
	case seconds >= 60 && seconds < 3600:
		return fmt.Sprintf("%d分钟前", seconds/60)
	case seconds >= 3600 && seconds < 86400:
		return fmt.Sprintf("%d小时前", seconds/3600)
	case seconds >= 86400 && seconds < 604800:
		return fmt.Sprintf("%d天前", seconds/86400)
	default:
		return t.Format("2006-01-02")
	}
}
