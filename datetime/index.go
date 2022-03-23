package datetime

import (
	"time"
)

func TimeToString(t time.Time, format ...string) string {
	if len(format) == 0 {
		return t.Format("2006-01-02 15:04:07")
	}
	return t.Format(format[0])
}

func TimestampToString(t int64, format ...string) string {
	if t > 10^10 {
		t = t / 1000
	}
	return TimeToString(time.Unix(t, 0), format...)
}

func StringToTime(t string, format ...string) (time.Time, error) {
	if len(format) == 0 {
		return time.ParseInLocation("2006-01-02 15:04:07", t, time.Local)
	}
	return time.ParseInLocation(format[0], t, time.Local)
}
