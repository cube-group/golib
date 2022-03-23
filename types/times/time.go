package times

import (
	"fmt"
	"math"
	"time"
)

var months = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

//json时间格式化
type JsonTime struct {
	time.Time
}

func NowTime() JsonTime {
	return JsonTime{time.Now()}
}

//实现json序列化方法,格式化时间
func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", FormatDatetime(t.Time))
	return []byte(stamp), nil
}

//TODO datetime默认值
func Default() time.Time {
	value, _ := time.Parse("0000-00-00 00:00:00", "1970-01-01 08:00:01")
	return value
}

//格式化为文件使用的年月日时分秒
func FormatFileDatetime(date time.Time) string {
	return date.Format("20060102150405")
}

//格式化为年月日时分秒
func FormatDatetime(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

//格式化为月日时分秒
func FormatDaytime(date time.Time) string {
	return date.Format("01-02 15:04:05")
}

//格式化为月日时分秒
func FormatTime(date time.Time) string {
	return date.Format("15:04:05")
}

//格式华为年月日
func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

//格式华为年月日时
func FormatDay(date time.Time) string {
	return date.Format("02")
}

//格式华为年月日时
func FormatDayHour(date time.Time) string {
	return date.Format("02 15:00")
}

//格式华为年月日时
func FormatHour(date time.Time) string {
	return date.Format("15:00")
}

func GetFormatMonth(options ...time.Time) int {
	var current time.Time
	if len(options) == 0 {
		current = time.Now()
	} else {
		current = options[0]
	}
	_, month, _ := current.Date()
	return months[month.String()]
}

/**
 * 1. 如果距当前时间60s内,则显示x秒内
 * 2. 如果距当前时间60m内,则显示x分钟内
 * 3. 如果距当前时间24h内,则显示x小时内
 * 4. 如果超过24小时,则显示x天前
 * @param $time
 * @return string
 */
func TimeFormat(v time.Time) string {
	if v == Default() {
		return ""
	}
	now := time.Now()
	diff := now.Sub(v).Seconds()

	if diff < 0 {
		return "非法时间"
	}
	if diff < 60 {
		return "刚刚"
	}
	if diff < (60 * 60) {
		return fmt.Sprintf("%v分钟前", math.Floor(diff/60))
	}
	if diff < (60 * 60 * 24) {
		return fmt.Sprintf("%v小时前", math.Floor(diff/(60*60)))
	}
	return fmt.Sprintf("%v天前", math.Floor(diff/(60*60*24)))

	if diff < (60 * 60 * 24 * 30) {
		return fmt.Sprintf("%v天前", math.Floor(diff/(60*60*24)))
	}
	if diff < (60 * 60 * 24 * 30 * 12) {
		return fmt.Sprintf("%v月前", math.Floor(diff/(60*60*24*30)))
	}
	return fmt.Sprintf("%v年前", math.Floor(diff/(60*60*24*30*12)))

}

//获取最近指定天数的datetime
func LastDatetime(currentTime time.Time, days int) string {
	return FormatDatetime(time.Unix(currentTime.Unix()-int64(days*24*3600), 0))
}

//格式化为当天开始时间
func DayStarTime(current time.Time) time.Time {
	return time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
}

//格式化为当天结束时间
func DayEndTime(current time.Time) time.Time {
	return time.Date(current.Year(), current.Month(), current.Day(), 23, 59, 59, 0, current.Location())
}

//根据标准字符串日志格式时间转为time.Time
func StrToTime(v string) (res time.Time, err error) {
	res, err = time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
	if err == nil {
		return
	}
	res, err = time.ParseInLocation("2006-01-02", v, time.Local)
	if err == nil {
		return
	}
	res, err = time.ParseInLocation("2006-01-02 15:04", v, time.Local)
	if err == nil {
		return
	}
	res, err = time.ParseInLocation("2006-01", v, time.Local)
	if err == nil {
		return
	}
	res, err = time.ParseInLocation("2006-01-02 15", v, time.Local)
	if err == nil {
		return
	}
	return
}

//根据标准字符串日志格式时间转为时间戳
func StrToTimeStamp(v string) int64 {
	res, err := StrToTime(v)
	if err != nil {
		return 0
	}
	return res.Unix()
}

func StrTimeFormat(v string) string {
	res, _ := StrToTime(v)
	return TimeFormat(res)
}

//获取当前utc时间无小数点
//例如：2018-11-01T07:15:56Z
func GetNowUTCTimeString() (string, int64, error) {
	return GetUTCTimeString(time.Now())
}

//获取utc时间无小数点
//例如：2018-11-01T07:15:56Z
func GetUTCTimeString(v time.Time) (string, int64, error) {
	utcTimeStamp := v.UTC().Unix()
	return FormatISO8601Date(utcTimeStamp), utcTimeStamp, nil
}
