package timex

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func SecondDuration(interval int64) time.Duration {
	return time.Duration(interval) * time.Second
}

func MinuteDuration(interval int64) time.Duration {
	return time.Duration(interval) * time.Minute
}

func HourDuration(interval int64) time.Duration {
	return time.Duration(interval) * time.Hour
}

//After 当前时间在t之后
func After(t time.Time) bool {
	return time.Now().Before(t)
}

//SecondAfter unix转换后时间加interval在当前时间之后
func SecondAfter(unix, interval int64) bool {
	return After(time.Unix(unix, 10).Add(SecondDuration(interval)))
}

//MinuteAfter unix转换后时间加interval在当前时间之后
func MinuteAfter(unix, interval int64) bool {
	return After(time.Unix(unix, 10).Add(MinuteDuration(interval)))
}

//HourAfter unix转换后时间加interval在当前时间之后
func HourAfter(unix, interval int64) bool {
	return After(time.Unix(unix, 10).Add(HourDuration(interval)))
}

//AfterByUnix unix转换后时间加t在当前时间之后
func AfterByUnix(unix int64, t time.Duration) bool {
	return After(time.Unix(unix, 10).Add(t))
}
