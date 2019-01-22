package timex

import (
	"github.com/chrisho/sd-helper/mathx"
	"time"
)

var (
	loc, _ = time.LoadLocation("Local") //重要：获取时区

	oneDay, _  = time.ParseDuration("24h")
	oneHour, _ = time.ParseDuration("1h")
)

const (
	YYYYMMDDHHIISS = "2006-01-02 15:04:05"
	YYYYMMDDHHII   = "2006-01-02 15:04"
	YYYYMMDDHH     = "2006-01-02 15"
	YYYYMMDD       = "2006-01-02"
	YYYYMM         = "2006-01"
	RFC3339        = time.RFC3339
)
const TimeIdSortMultiple int64 = 1E9

func TimeNow() time.Time {
	return time.Now()
}

func TodayDate() string {
	return time.Now().Format(YYYYMMDD)
}

func TodayUnix() int64 {
	dateTime, _ := time.ParseInLocation(YYYYMMDD, TodayDate(), time.Local)

	return dateTime.Unix()
}

// get time now unixtime == php time()
func UnixTime() int32 {
	return int32(time.Now().Unix())
}

// get time now date format YYYY-MM-DD HH:II:SS
func DateTime() string {
	return time.Now().Format(YYYYMMDDHHIISS)
}

// 这月的某天
func ThisMonthDayTime(timeTime time.Time, whichDay int) time.Time {
	if whichDay <= 0 {
		whichDay = 1
	}
	return time.Date(timeTime.Year(), timeTime.Month(), whichDay, 0, 0, 0, 0, time.Local)
}

// 下月1号时间
func NextMonth01Time(timeTime time.Time) time.Time {
	return time.Date(timeTime.Year(), timeTime.Month(), 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 1, 0)
}

// RFC3339 time string to format Y-m-d H:i,Y-m-d H ...
func RFC3339Time2Date(timeStr, dateLayout string) string {
	dateTime, _ := time.ParseInLocation(time.RFC3339, timeStr, time.Local)

	return dateTime.Format(dateLayout)
}

// change date to format Y-m-d H:i,Y-m-d H ...
func Date2Date(inputDateLayout, inputDate, resultDateLayout string) string {
	dateTime, _ := time.ParseInLocation(inputDateLayout, inputDate, time.Local)

	return dateTime.Format(resultDateLayout)
}

// date to time.Time
func Date2Time(date, dateFormat string) time.Time {
	timeTime, _ := time.ParseInLocation(dateFormat, date, time.Local)

	return timeTime
}

// Timestamp to time.Time
func Timestamp2Time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

//  date to timestamp
func Date2Timestamp(dateLayout, date string) int64 {
	localTime, _ := time.ParseInLocation(dateLayout, date, time.Local)

	return localTime.Unix()
}

// timestamp to format Y-m-d H:i:s
func Timestamp2Date(timestamp int64, dateLayout string) string {
	return time.Unix(timestamp, 0).Format(dateLayout)
}

// 两个时间相差毫秒数
func TwoTimeMillisecond(t1, t2 time.Time) int64 {
	return t2.UnixNano()/1e6 - t1.UnixNano()/1e6
}

func TimeSince(start time.Time) time.Duration {
	return time.Since(start)
}

func DateSlice() []byte {
	var dateSlice []byte
	dateStr := DateTime()
	dateStrSlice := []byte(dateStr)

	for _, r := range dateStrSlice {
		if r >= '0' && r <= '9' {
			dateSlice = append(dateSlice, r)
		}
	}

	return dateSlice
}

func DateYYYYMMDDHHIISS() string {
	dateSlice := DateSlice()

	return string(dateSlice)
}

func DateYYMMDDHHIISS() string {
	dateSlice := DateSlice()

	return string(dateSlice[2 : len(dateSlice)-1])
}

// 默认长度为19为，time(10位）* 1E9(9位） + id， 修改长度，请修改改1E9...
func TimeIdSort(time, id int32) int64 {
	if time <= 0 {
		return 0
	}
	return int64(time)*TimeIdSortMultiple + int64(id)
}

//获取指定月份第一秒和最后一秒
func GetMonthFirstLast(year, month int) (string, int64, int64) {
	if year < 1970 || month > 12 || month < 1 {
		return "", 0, 0
	}
	var futureMonthFirst time.Time
	if month >= 12 {
		futureMonthFirst = time.Date(year+1, time.Month((month+1)%12), 1, 23, 59, 59, 0, time.Local)
	} else {
		futureMonthFirst = time.Date(year, time.Month(month+1), 1, 23, 59, 59, 0, time.Local)
	}
	// 指定月份最后一天最后一秒时间
	lastDate := futureMonthFirst.Add(-1 * oneDay)
	// 指定月份第一天第一秒
	firstDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	monthStr := firstDate.Format(YYYYMM)
	return monthStr, firstDate.Unix(), lastDate.Unix()
}

//获取某月有多少天
func MonthDays(year, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

//获取指定时间戳下n月后同一天的时间戳
// 如果下n个月没有这一天，则返回该月的最后一天
func NextNMonths(ts int64, months int) int64 {
	dt := time.Unix(ts, 0)
	month := int(dt.Month()) + months
	year := dt.Year() + int(month/12)
	month = month % 12
	day := mathx.MinInt(dt.Day(), MonthDays(year, month))
	dt = time.Date(year, time.Month(month), day, dt.Hour(), dt.Minute(), dt.Second(), 0, loc)
	return dt.Unix()
}
func DateNextNMonths(dt time.Time, months int) time.Time {
	month := int(dt.Month()) + months
	year := dt.Year() + int(month/12)
	month = month % 12
	day := mathx.MinInt(dt.Day(), MonthDays(year, month))
	dt = time.Date(year, time.Month(month), day, dt.Hour(), dt.Minute(), dt.Second(), 0, loc)
	return dt
}

const (
	NoDateEnd = 1 + iota
	NoDateStart
	DateStartEnd
)

func eachMonthLastDay(startDate, endDate time.Time, convertType int) [2]int64 {
	switch convertType {
	case NoDateEnd:
		var nextYear int
		var nextMonth int
		if startDate.Month() == 12 {
			nextYear = startDate.Year() + 1
			nextMonth = 1
		} else {
			nextYear = startDate.Year()
			nextMonth = int(startDate.Month()) + 1
		}
		tEndDate := time.Date(nextYear, time.Month(nextMonth), 1, 23, 59, 59, 0, loc)
		d, _ := time.ParseDuration("-24h")
		tEndDate = tEndDate.Add(d)
		return [2]int64{startDate.Unix(), tEndDate.Unix()}
	case NoDateStart:
		startDate = time.Date(startDate.Year(), startDate.Month(), 26, 0, 0, 0, 0, loc)
		return [2]int64{startDate.Unix(), endDate.Unix()}
	case DateStartEnd:
		return [2]int64{startDate.Unix(), endDate.Unix()}
	default:
		return [2]int64{0, 0}
	}
}

//获取指定时间段内包含26号之后的所有时间段
func GetEachMonthLastDays(startTs, endTs int64) (result [][2]int64) {
	result = make([][2]int64, 0, 5)
	startDate := time.Unix(startTs, 0)
	endDate := time.Unix(endTs, 0)
	if startDate.Month() == endDate.Month() && endDate.Day() < 26 {
		return
	}
	tStartDate := time.Date(startDate.Year(), startDate.Month(), 26, 0, 0, 0, 0, loc)
	if endDate.Day() < 26 {
		result = append(result, eachMonthLastDay(tStartDate, tStartDate, NoDateEnd))
	} else {
		tEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, loc)
		result = append(result, eachMonthLastDay(tEndDate, tEndDate, NoDateStart))
		result = append(result, eachMonthLastDay(tStartDate, tStartDate, NoDateEnd))
	}
	addMonth := (endDate.Year()-startDate.Year())*12 + (int(endDate.Month()) - int(startDate.Month())) - 1
	for i := 0; i < addMonth; i++ {
		nextNMonth := DateNextNMonths(tStartDate, i+1)
		result = append(result, eachMonthLastDay(nextNMonth, nextNMonth, NoDateEnd))
	}
	return
}
