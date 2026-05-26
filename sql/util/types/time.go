// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"bytes"
	"regexp"
	gotime "time"

	"github.com/chrislusf/gleam/sql/mysql"
	"github.com/juju/errors"
)

// Portable analogs of some common call errors.
var (
	ErrInvalidTimeFormat = errors.New("invalid time format")
	ErrInvalidYearFormat = errors.New("invalid year format")
	ErrInvalidYear       = errors.New("invalid year")
)

// Time format without fractional seconds precision.
const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
	// TimeFSPFormat is time format with fractional seconds precision.
	TimeFSPFormat = "2006-01-02 15:04:05.000000"
)

const (
	// MinYear is the minimum for mysql year type.
	MinYear int16 = 1901
	// MaxYear is the maximum for mysql year type.
	MaxYear int16 = 2155

	// MinTime is the minimum for mysql time type.
	MinTime = -gotime.Duration(838*3600+59*60+59) * gotime.Second
	// MaxTime is the maximum for mysql time type.
	MaxTime = gotime.Duration(838*3600+59*60+59) * gotime.Second

	zeroDatetimeStr = "0000-00-00 00:00:00"
	zeroDateStr     = "0000-00-00"
)

// Zero values for different types.
var (
	// ZeroDuration is the zero value for Duration type.
	ZeroDuration = Duration{Duration: gotime.Duration(0), Fsp: DefaultFsp}

	// ZeroTime is the zero value for TimeInternal type.
	ZeroTime = mysqlTime{}

	// ZeroDatetime is the zero value for datetime Time.
	ZeroDatetime = Time{
		Time: ZeroTime,
		Type: mysql.TypeDatetime,
		Fsp:  DefaultFsp,
	}

	// ZeroTimestamp is the zero value for timestamp Time.
	ZeroTimestamp = Time{
		Time: ZeroTime,
		Type: mysql.TypeTimestamp,
		Fsp:  DefaultFsp,
	}

	// ZeroDate is the zero value for date Time.
	ZeroDate = Time{
		Time: ZeroTime,
		Type: mysql.TypeDate,
		Fsp:  DefaultFsp,
	}

	local = gotime.Local
)

var (
	// minDatetime is the minimum for mysql datetime type.
	minDatetime = FromDate(1000, 1, 1, 0, 0, 0, 0)
	// maxDatetime is the maximum for mysql datetime type.
	maxDatetime = FromDate(9999, 12, 31, 23, 59, 59, 999999)

	// minTimestamp is the minimum for mysql timestamp type.
	minTimestamp = gotime.Date(1970, 1, 1, 0, 0, 1, 0, gotime.UTC)
	// maxTimestamp is the maximum for mysql timestamp type.
	maxTimestamp = gotime.Date(2038, 1, 19, 3, 14, 7, 999999, gotime.UTC)

	// WeekdayNames lists names of weekdays, which are used in builtin time function `dayname`.
	WeekdayNames = []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	// MonthNames lists names of months, which are used in builtin time function `monthname`.
	MonthNames = []string{
		"January", "February",
		"March", "April",
		"May", "June",
		"July", "August",
		"September", "October",
		"November", "December",
	}
)

// TimeInternal is the internal representation for mysql time in TiDB.
type TimeInternal interface {
	Year() int
	Month() int
	Day() int
	Hour() int
	Minute() int
	Second() int
	Weekday() gotime.Weekday
	YearDay() int
	YearWeek(mode int) (int, int)
	Week(mode int) int
	Microsecond() int
	GoTime(*gotime.Location) (gotime.Time, error)
}

// FromGoTime translates time.Time to mysql time internal representation.
func FromGoTime(t gotime.Time) TimeInternal { _ = "STUB: not implemented"; return *new(TimeInternal) }

// FromDate makes a internal time representation from the given date.
func FromDate(year int, month int, day int, hour int, minute int, second int, microsecond int) TimeInternal {
	_ = "STUB: not implemented"
	return *new(TimeInternal)
}

// Clock returns the hour, minute, and second within the day specified by t.
func (t Time) Clock() (hour int, minute int, second int) { _ = "STUB: not implemented"; return 0, 0, 0 }

// Time is the struct for handling datetime, timestamp and date.
// TODO: check if need a NewTime function to set Fsp default value?
type Time struct {
	Time TimeInternal
	Type uint8
	// Fsp is short for Fractional Seconds Precision.
	// See http://dev.mysql.com/doc/refman/5.7/en/fractional-seconds.html
	Fsp int
}

// CurrentTime returns current time with type tp.
func CurrentTime(tp uint8) Time { _ = "STUB: not implemented"; return *new(Time) }

func (t Time) String() string { _ = "STUB: not implemented"; return "" }

// We control the format, so no error would occur.

// IsZero returns a boolean indicating whether the time is equal to ZeroTime.
func (t Time) IsZero() bool { _ = "STUB: not implemented"; return false }

// InvalidZero returns a boolean indicating whether the month or day is zero.
func (t Time) InvalidZero() bool { _ = "STUB: not implemented"; return false }

const numberFormat = "%Y%m%d%H%i%s"
const dateFormat = "%Y%m%d"

// ToNumber returns a formatted number.
// e.g,
// 2012-12-12 -> 20121212
// 2012-12-12T10:10:10 -> 20121212101010
// 2012-12-12T10:10:10.123456 -> 20121212101010.123456
func (t Time) ToNumber() *MyDecimal { _ = "STUB: not implemented"; return nil }

// Fix issue #1046
// Prevents from converting 2012-12-12 to 20121212000000

// We skip checking error here because time formatted string can be parsed certainly.

// Convert converts t with type tp.
func (t Time) Convert(tp uint8) (Time, error) { _ = "STUB: not implemented"; return *new(Time), nil }

// ConvertToDuration converts mysql datetime, timestamp and date to mysql time type.
// e.g,
// 2012-12-12T10:10:10 -> 10:10:10
// 2012-12-12 -> 0
func (t Time) ConvertToDuration() (Duration, error) {
	_ = "STUB: not implemented"
	return *new(Duration), nil
}

// TODO: check convert validation

// Compare returns an integer comparing the time instant t to o.
// If t is after o, return 1, equal o, return 0, before o, return -1.
func (t Time) Compare(o Time) int { _ = "STUB: not implemented"; return 0 }

func compareTime(a, b TimeInternal) int { _ = "STUB: not implemented"; return 0 }

// CompareString is like Compare,
// but parses string to Time then compares.
func (t Time) CompareString(str string) (int, error) {
	_ = "STUB: not implemented"
	// use MaxFsp to parse the string
	return 0, nil
}

// roundTime rounds the time value according to digits count specified by fsp.
func roundTime(t gotime.Time, fsp int) gotime.Time {
	_ = "STUB: not implemented"
	return *new(gotime.Time)
}

func (t Time) roundFrac(fsp int) (Time, error) { _ = "STUB: not implemented"; return *new(Time), nil }

// date type has no fsp

// have same fsp

// TODO: Consider time_zone variable.

// Take the hh:mm:ss part out to avoid handle month or day = 0.

// TODO: when hh:mm:ss overflow one day after rounding, it should be add to yy:mm:dd part,
// but mm:dd may contain 0, it makes the code complex, so we ignore it here.

// RoundFrac rounds fractional seconds precision with new fsp and returns a new one.
// We will use the “round half up” rule, e.g, >= 0.5 -> 1, < 0.5 -> 0,
// so 2011:11:11 10:10:10.888888 round 0 -> 2011:11:11 10:10:11
// and 2011:11:11 10:10:10.111111 round 0 -> 2011:11:11 10:10:10
func RoundFrac(t gotime.Time, fsp int) (gotime.Time, error) {
	_ = "STUB: not implemented"
	return *new(gotime.Time), nil
}

// ToPackedUint encodes Time to a packed uint64 value.
//
//	 1 bit  0
//	17 bits year*13+month   (year 0-9999, month 0-12)
//	 5 bits day             (0-31)
//	 5 bits hour            (0-23)
//	 6 bits minute          (0-59)
//	 6 bits second          (0-59)
//	24 bits microseconds    (0-999999)
//
//	Total: 64 bits = 8 bytes
//
//	0YYYYYYY.YYYYYYYY.YYdddddh.hhhhmmmm.mmssssss.ffffffff.ffffffff.ffffffff
func (t Time) ToPackedUint() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO: Consider time_zone variable.

// mysql timestamp month and day can't be zero.

// FromPackedUint decodes Time from a packed uint64 value.
func (t *Time) FromPackedUint(packed uint64) error { _ = "STUB: not implemented"; return nil }

func (t *Time) check() error { _ = "STUB: not implemented"; return nil }

// Sub subtracts t1 from t, returns a duration value.
// Note that sub should not be done on different time types.
func (t *Time) Sub(t1 *Time) Duration { _ = "STUB: not implemented"; return *new(Duration) }

// TODO: Consider time_zone variable.

// TimestampDiff returns t2 - t1 where t1 and t2 are date or datetime expressions.
// The unit for the result (an integer) is given by the unit argument.
// The legal values for unit are "YEAR" "QUARTER" "MONTH" "DAY" "HOUR" "SECOND" and so on.
func TimestampDiff(unit string, t1 Time, t2 Time) int64 { _ = "STUB: not implemented"; return 0 }

func parseDateFormat(format string) []string { _ = "STUB: not implemented"; return nil }

// Date format must start and end with number.

// Separator is a single none-number char.

func parseDatetime(str string, fsp int) (Time, error) {
	_ = "STUB: not implemented"
	// Try to split str with delimiter.
	// TODO: only punctuation can be the delimiter for date parts or time parts.
	// But only space and T can be the delimiter between the date and time part.
	return *new(Time), nil
}

// No delimiter.

// YYYYMMDDHHMMSS

// YYMMDDHHMMSS

// YYYYMMDD

// YYMMDD

// YYYYMMDDHHMMSS.fraction

// YYMMDDHHMMSS.fraction

// YYYY-MM-DD

// We don't have fractional seconds part.
// YYYY-MM-DD HH-MM-SS

// We have fractional seconds part.
// YYY-MM-DD HH-MM-SS.fraction

// If str is sepereated by delimiters, the first one is year, and if the year is 2 digit,
// we should adjust it.
// TODO: ajust year is very complex, now we only consider the simplest way.

// Convert to Go time and add 1 second, to handle input like 2017-01-05 08:40:59.575601

func scanTimeArgs(seps []string, args ...*int) error { _ = "STUB: not implemented"; return nil }

// ParseYear parses a formatted string and returns a year number.
func ParseYear(str string) (int16, error) { _ = "STUB: not implemented"; return 0, nil }

// Nothing to do.

// See https://dev.mysql.com/doc/refman/5.7/en/two-digit-years.html
func adjustYear(y int) int { _ = "STUB: not implemented"; return 0 }

// AdjustYear is used for adjusting year and checking its validation.
func AdjustYear(y int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Duration is the type for MySQL time type.
type Duration struct {
	gotime.Duration
	// Fsp is short for Fractional Seconds Precision.
	// See http://dev.mysql.com/doc/refman/5.7/en/fractional-seconds.html
	Fsp int
}

// String returns the time formatted using default TimeFormat and fsp.
func (d Duration) String() string { _ = "STUB: not implemented"; return "" }

func (d Duration) formatFrac(frac int) string { _ = "STUB: not implemented"; return "" }

// ToNumber changes duration to number format.
// e.g,
// 10:10:10 -> 101010
func (d Duration) ToNumber() *MyDecimal { _ = "STUB: not implemented"; return nil }

// We skip checking error here because time formatted string can be parsed certainly.

// ConvertToTime converts duration to Time.
// Tp is TypeDatetime, TypeTimestamp and TypeDate.
func (d Duration) ConvertToTime(tp uint8) (Time, error) {
	_ = "STUB: not implemented"
	return *new(Time), nil
}

// just use current year, month and day.

// RoundFrac rounds fractional seconds precision with new fsp and returns a new one.
// We will use the “round half up” rule, e.g, >= 0.5 -> 1, < 0.5 -> 0,
// so 10:10:10.999999 round 0 -> 10:10:11
// and 10:10:10.000000 round 0 -> 10:10:10
func (d Duration) RoundFrac(fsp int) (Duration, error) {
	_ = "STUB: not implemented"
	return *new(Duration), nil
}

// Compare returns an integer comparing the Duration instant t to o.
// If d is after o, return 1, equal o, return 0, before o, return -1.
func (d Duration) Compare(o Duration) int { _ = "STUB: not implemented"; return 0 }

// CompareString is like Compare,
// but parses str to Duration then compares.
func (d Duration) CompareString(str string) (int, error) {
	_ = "STUB: not implemented"
	// use MaxFsp to parse the string
	return 0, nil
}

// Hour returns current hour.
// e.g, hour("11:11:11") -> 11
func (d Duration) Hour() int { _ = "STUB: not implemented"; return 0 }

// Minute returns current minute.
// e.g, hour("11:11:11") -> 11
func (d Duration) Minute() int { _ = "STUB: not implemented"; return 0 }

// Second returns current second.
// e.g, hour("11:11:11") -> 11
func (d Duration) Second() int { _ = "STUB: not implemented"; return 0 }

// MicroSecond returns current microsecond.
// e.g, hour("11:11:11.11") -> 110000
func (d Duration) MicroSecond() int { _ = "STUB: not implemented"; return 0 }

// ParseDuration parses the time form a formatted string with a fractional seconds part,
// returns the duration type Time value.
// See http://dev.mysql.com/doc/refman/5.7/en/fractional-seconds.html
func ParseDuration(str string, fsp int) (Duration, error) {
	_ = "STUB: not implemented"
	return *new(Duration), nil
}

// Time format may has day.

// It has fractional precesion parts.

// It tries to split str with delimiter, time delimiter must be :

// No delimiter.

// HHMMSS

// MMSS

// SS

// Maybe only contains date.

// HH:MM

// Time format maybe HH:MM:SS or HHH:MM:SS.
// See https://dev.mysql.com/doc/refman/5.7/en/time.html

func splitDuration(t gotime.Duration) (int, int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0
}

var maxDaysInMonth = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func getTime(num int64, tp byte) (Time, error) { _ = "STUB: not implemented"; return *new(Time), nil }

// See number_to_datetime function.
// https://github.com/mysql/mysql-server/blob/5.7/sql-common/my_time.c
func parseDateTimeFromNum(num int64) (Time, error) {
	_ = "STUB: not implemented"

	// Check zero.
	return *new(Time), nil
}

// Check datetime type.

// Check MMDD.

// Adjust year
// YYMMDD, year: 2000-2069

// Check YYMMDD.

// Adjust year
// YYMMDD, year: 1970-1999

// Check YYYYMMDD.

// Adjust hour/min/second.

// Check MMDDHHMMSS.

// Set TypeDatetime type.

// Adjust year
// YYMMDDHHMMSS, 2000-2069

// Check YYYYMMDDHHMMSS.

// Adjust year
// YYMMDDHHMMSS, 1970-1999

// ParseTime parses a formatted string with type tp and specific fsp.
// Type is TypeDatetime, TypeTimestamp and TypeDate.
// Fsp is in range [0, 6].
// MySQL supports many valid datatime format, but still has some limitation.
// If delimiter exists, the date part and time part is separated by a space or T,
// other punctuation character can be used as the delimiter between date parts or time parts.
// If no delimiter, the format must be YYYYMMDDHHMMSS or YYMMDDHHMMSS
// If we have fractional seconds part, we must use decimal points as the delimiter.
// The valid datetime range is from '1000-01-01 00:00:00.000000' to '9999-12-31 23:59:59.999999'.
// The valid timestamp range is from '1970-01-01 00:00:01.000000' to '2038-01-19 03:14:07.999999'.
// The valid date range is from '1000-01-01' to '9999-12-31'
func ParseTime(str string, tp byte, fsp int) (Time, error) {
	_ = "STUB: not implemented"
	return *new(Time), nil
}

// ParseDatetime is a helper function wrapping ParseTime with datetime type and default fsp.
func ParseDatetime(str string) (Time, error) { _ = "STUB: not implemented"; return *new(Time), nil }

// ParseTimestamp is a helper function wrapping ParseTime with timestamp type and default fsp.
func ParseTimestamp(str string) (Time, error) { _ = "STUB: not implemented"; return *new(Time), nil }

// ParseDate is a helper function wrapping ParseTime with date type.
func ParseDate(str string) (Time, error) {
	_ = "STUB: not implemented"
	// date has no fractional seconds precision
	return *new(Time), nil
}

// ParseTimeFromNum parses a formatted int64,
// returns the value which type is tp.
func ParseTimeFromNum(num int64, tp byte, fsp int) (Time, error) {
	_ = "STUB: not implemented"
	return *new(Time), nil
}

// ParseDatetimeFromNum is a helper function wrapping ParseTimeFromNum with datetime type and default fsp.
func ParseDatetimeFromNum(num int64) (Time, error) {
	_ = "STUB: not implemented"
	return *new(Time), nil
}

// ParseTimestampFromNum is a helper function wrapping ParseTimeFromNum with timestamp type and default fsp.
func ParseTimestampFromNum(num int64) (Time, error) {
	_ = "STUB: not implemented"
	return *new(Time), nil
}

// ParseDateFromNum is a helper function wrapping ParseTimeFromNum with date type.
func ParseDateFromNum(num int64) (Time, error) {
	_ = "STUB: not implemented"
	// date has no fractional seconds precision
	return *new(Time), nil
}

// TimeFromDays Converts a day number to a date.
func TimeFromDays(num int64) Time { _ = "STUB: not implemented"; return *new(Time) }

func checkDateType(t TimeInternal) error { _ = "STUB: not implemented"; return nil }

func checkDateRange(t TimeInternal) error {
	_ = "STUB: not implemented"
	// Oddly enough, MySQL document says date range should larger than '1000-01-01',
	// but we can insert '0001-01-01' actually.
	return nil
}

func checkMonthDay(year, month, day int) error { _ = "STUB: not implemented"; return nil }

func checkTimestampType(t TimeInternal) error { _ = "STUB: not implemented"; return nil }

// TODO: Consider time_zone variable.

func checkDatetimeType(t TimeInternal) error { _ = "STUB: not implemented"; return nil }

// ExtractTimeNum extracts time value number from time unit and format.
func ExtractTimeNum(unit string, t Time) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO: Consider time_zone variable.

// 1 - 3 -> 1
// 4 - 6 -> 2
// 7 - 9 -> 3
// 10 - 12 -> 4

func extractSingleTimeValue(unit string, format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `SS.FFFFFF`.
func extractSecondMicrosecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `MM:SS.FFFFFF`.
func extractMinuteMicrosecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `MM:SS`.
func extractMinuteSecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `HH:MM:SS.FFFFFF`.
func extractHourMicrosecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `HH:MM:SS`.
func extractHourSecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `HH:MM`.
func extractHourMinute(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `DD HH:MM:SS.FFFFFF`.
func extractDayMicrosecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `DD HH:MM:SS`.
func extractDaySecond(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `DD HH:MM`.
func extractDayMinute(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `DD HH`.
func extractDayHour(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// Format is `YYYY-MM`.
func extractYearMonth(format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// ExtractTimeValue extracts time value from time unit and format.
func ExtractTimeValue(unit string, format string) (int64, int64, int64, gotime.Duration, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, *new(gotime.Duration), nil
}

// IsClockUnit returns true when unit is interval unit with hour, minute or second.
func IsClockUnit(unit string) bool { _ = "STUB: not implemented"; return false }

// IsDateFormat returns true when the specified time format could contain only date.
func IsDateFormat(format string) bool { _ = "STUB: not implemented"; return false }

// ParseTimeFromInt64 parses mysql time value from int64.
func ParseTimeFromInt64(num int64) (Time, error) { _ = "STUB: not implemented"; return *new(Time), nil }

// DateFormat returns a textual representation of the time value formatted
// according to layout.
// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_date-format
func (t Time) DateFormat(layout string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// It's not in pattern match now.

var abbrevWeekdayName = []string{
	"Sun", "Mon", "Tue",
	"Wed", "Thu", "Fri", "Sat",
}

func (t Time) convertDateFormat(b rune, buf *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func abbrDayOfMonth(day int) string { _ = "STUB: not implemented"; return "" }

// StrToDate converts date string according to format.
// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_date-format
func (t *Time) StrToDate(date, format string) bool { _ = "STUB: not implemented"; return false }

// mysqlTimeFix fixes the mysqlTime use the values in the context.
func mysqlTimeFix(t *mysqlTime, ctx map[string]int) error {
	_ = "STUB: not implemented"
	// Key of the ctx is the format char, such as `%j` `%p` and so on.
	return nil
}

// TODO: Implement the function that converts day of year to yy:mm:dd.

// 12 is a special hour.

// strToDate converts date string according to format, returns true on success,
// the value will be stored in argument t or ctx.
func strToDate(t *mysqlTime, date string, format string, ctx map[string]int) bool {
	_ = "STUB: not implemented"
	return false
}

// Extra characters at the end of date are ignored.

// getFormatToken takes one format control token from the string.
// format "%d %H %m" will get token "%d" and the remain is " %H %m".
func getFormatToken(format string) (token string, remain string, succ bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// Just one character.

// More than one character.

func skipEmptySpace(input string) string { _ = "STUB: not implemented"; return "" }

var weekdayAbbrev = map[string]gotime.Weekday{
	"Sun": gotime.Sunday,
	"Mon": gotime.Monday,
	"Tue": gotime.Tuesday,
	"Wed": gotime.Wednesday,
	"Thu": gotime.Tuesday,
	"Fri": gotime.Friday,
	"Sat": gotime.Saturday,
}

var monthAbbrev = map[string]gotime.Month{
	"Jan": gotime.January,
	"Feb": gotime.February,
	"Mar": gotime.March,
	"Apr": gotime.April,
	"May": gotime.May,
	"Jun": gotime.June,
	"Jul": gotime.July,
	"Aug": gotime.August,
	"Sep": gotime.September,
	"Oct": gotime.October,
	"Nov": gotime.November,
	"Dec": gotime.December,
}

type dateFormatParser func(t *mysqlTime, date string, ctx map[string]int) (remain string, succ bool)

var dateFormatParserTable = map[string]dateFormatParser{
	"%b": abbreviatedMonth,           // Abbreviated month name (Jan..Dec)
	"%c": monthNumeric,               // Month, numeric (0..12)
	"%d": dayOfMonthNumericTwoDigits, // Day of the month, numeric (00..31)
	"%e": dayOfMonthNumeric,          // Day of the month, numeric (0..31)
	"%f": microSeconds,               // Microseconds (000000..999999)
	"%h": hour24TwoDigits,            // Hour (01..12)
	"%H": hour24TwoDigits,            // Hour (01..12)
	"%I": hour24TwoDigits,            // Hour (01..12)
	"%i": minutesNumeric,             // Minutes, numeric (00..59)
	"%j": dayOfYearThreeDigits,       // Day of year (001..366)
	"%k": hour24Numeric,              // Hour (0..23)
	"%l": hour12Numeric,              // Hour (1..12)
	"%M": fullNameMonth,              // Month name (January..December)
	"%m": monthNumericTwoDigits,      // Month, numeric (00..12)
	"%p": isAMOrPM,                   // AM or PM
	"%r": time12Hour,                 // Time, 12-hour (hh:mm:ss followed by AM or PM)
	"%s": secondsNumeric,             // Seconds (00..59)
	"%S": secondsNumeric,             // Seconds (00..59)
	"%T": time24Hour,                 // Time, 24-hour (hh:mm:ss)
	"%Y": yearNumericFourDigits,      // Year, numeric, four digits
	// TODO: Add the following...
	// "%a": abbreviatedWeekday,         // Abbreviated weekday name (Sun..Sat)
	// "%D": dayOfMonthWithSuffix,       // Day of the month with English suffix (0th, 1st, 2nd, 3rd)
	// "%U": weekMode0,                  // Week (00..53), where Sunday is the first day of the week; WEEK() mode 0
	// "%u": weekMode1,                  // Week (00..53), where Monday is the first day of the week; WEEK() mode 1
	// "%V": weekMode2,                  // Week (01..53), where Sunday is the first day of the week; WEEK() mode 2; used with %X
	// "%v": weekMode3,                  // Week (01..53), where Monday is the first day of the week; WEEK() mode 3; used with %x
	// "%W": weekdayName,                // Weekday name (Sunday..Saturday)
	// "%w": dayOfWeek,                  // Day of the week (0=Sunday..6=Saturday)
	// "%X": yearOfWeek,                 // Year for the week where Sunday is the first day of the week, numeric, four digits; used with %V
	// "%x": yearOfWeek,                 // Year for the week, where Monday is the first day of the week, numeric, four digits; used with %v
	// Deprecated since MySQL 5.7.5
	// "%y": yearTwoDigits,         // Year, numeric (two digits)
}

func matchDateWithToken(t *mysqlTime, date string, token string, ctx map[string]int) (remain string, succ bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseDigits(input string, count int) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func hour24TwoDigits(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func secondsNumeric(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func minutesNumeric(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

const time12HourLen = len("hh:mm:ssAM")

func time12Hour(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	// hh:mm:ss AM
	return "", false
}

const time24HourLen = len("hh:mm:ss")

func time24Hour(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	// hh:mm:ss
	return "", false
}

const (
	constForAM = 1 + iota
	constForPM
)

func isAMOrPM(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func dayOfMonthNumericTwoDigits(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

var twoDigitRegex = regexp.MustCompile("^[1-9][0-9]?")

// parseTwoNumeric is used for pattens 0..31 0..24 0..60 and so on.
// It returns the parsed int, and remain data after parse.
func parseTwoNumeric(input string) (int, string) { _ = "STUB: not implemented"; return 0, "" }

func dayOfMonthNumeric(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// 0..31

func hour24Numeric(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	// 0..23
	return "", false
}

func hour12Numeric(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	// 1..12
	return "", false
}

func microSeconds(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func yearNumericFourDigits(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func dayOfYearThreeDigits(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func monthNumericTwoDigits(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func abbreviatedWeekday(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// TODO: We need refact mysql time to support this.

func abbreviatedMonth(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func fullNameMonth(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func monthNumeric(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// 0th 1st 2nd 3rd ...
func dayOfMonthWithSuffix(t *mysqlTime, input string, ctx map[string]int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseOrdinalNumbers(input string) (value int, remain string) {
	_ = "STUB: not implemented"
	return 0, ""
}
