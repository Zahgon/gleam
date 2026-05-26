// Copyright 2016 PingCAP, Inc.
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
	gotime "time"
)

type mysqlTime struct {
	year        uint16 // year <= 9999
	month       uint8  // month <= 12
	day         uint8  // day <= 31
	hour        uint8  // hour <= 23
	minute      uint8  // minute <= 59
	second      uint8  // second <= 59
	microsecond uint32
}

func (t mysqlTime) Year() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Month() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Day() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Hour() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Minute() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Second() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Microsecond() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) Weekday() gotime.Weekday {
	_ = "STUB: not implemented"
	// TODO: Consider time_zone variable.
	return *new(gotime.Weekday)
}

func (t mysqlTime) YearDay() int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) YearWeek(mode int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (t mysqlTime) Week(mode int) int { _ = "STUB: not implemented"; return 0 }

func (t mysqlTime) GoTime(loc *gotime.Location) (gotime.Time, error) {
	_ = "STUB: not implemented"
	// gotime.Time can't represent month 0 or day 0, date contains 0 would be converted to a nearest date,
	// For example, 2006-12-00 00:00:00 would become 2015-11-30 23:59:59.
	return *new(gotime.Time), nil
}

// This function will check the result, and return an error if it's not the same with the origin input.

func newMysqlTime(year, month, day, hour, minute, second, microsecond int) mysqlTime {
	_ = "STUB: not implemented"
	return *new(mysqlTime)
}

func calcTimeFromSec(to *mysqlTime, seconds, microseconds int) { _ = "STUB: not implemented"; return }

const secondsIn24Hour = 86400

// calcTimeDiff calculates difference between two datetime values as seconds + microseconds.
// t1 and t2 should be TIME/DATE/DATETIME value.
// sign can be +1 or -1, and t2 is preprocessed with sign first.
func calcTimeDiff(t1, t2 TimeInternal, sign int) (seconds, microseconds int, neg bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// datetimeToUint64 converts time value to integer in YYYYMMDDHHMMSS format.
func datetimeToUint64(t TimeInternal) uint64 { _ = "STUB: not implemented"; return 0 }

// dateToUint64 converts time value to integer in YYYYMMDD format.
func dateToUint64(t TimeInternal) uint64 { _ = "STUB: not implemented"; return 0 }

// timeToUint64 converts time value to integer in HHMMSS format.
func timeToUint64(t TimeInternal) uint64 { _ = "STUB: not implemented"; return 0 }

// calcDaynr calculates days since 0000-00-00.
func calcDaynr(year, month, day int) int { _ = "STUB: not implemented"; return 0 }

// DateDiff calculates number of days between two days.
func DateDiff(startTime, endTime TimeInternal) int { _ = "STUB: not implemented"; return 0 }

// calcDaysInYear calculates days in one year, it works with 0 <= year <= 99.
func calcDaysInYear(year int) int { _ = "STUB: not implemented"; return 0 }

// calcWeekday calculates weekday from daynr, returns 0 for Monday, 1 for Tuesday ...
func calcWeekday(daynr int, sundayFirstDayOfWeek bool) int { _ = "STUB: not implemented"; return 0 }

type weekBehaviour uint

const (
	// If set, Sunday is first day of week, otherwise Monday is first day of week.
	weekBehaviourMondayFirst weekBehaviour = 1 << iota
	// If set, Week is in range 1-53, otherwise Week is in range 0-53.
	// Note that this flag is only releveant if WEEK_JANUARY is not set.
	weekBehaviourYear
	// If not set, Weeks are numbered according to ISO 8601:1988.
	// If set, the week that contains the first 'first-day-of-week' is week 1.
	weekBehaviourFirstWeekday
)

func (v weekBehaviour) test(flag weekBehaviour) bool { _ = "STUB: not implemented"; return false }

func weekMode(mode int) weekBehaviour { _ = "STUB: not implemented"; return *new(weekBehaviour) }

// calcWeek calculates week and year for the time.
func calcWeek(t *mysqlTime, wb weekBehaviour) (year int, week int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// mixDateAndTime mixes a date value and a time value.
func mixDateAndTime(date, time *mysqlTime, neg bool) { _ = "STUB: not implemented"; return }

// Time is negative or outside of 24 hours internal.

// If we want to use this function with arbitrary dates, this code will need
// to cover cases when time is negative and "date < -time".

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// getDateFromDaynr changes a daynr to year, month and day,
// daynr 0 is returned as date 00.00.00
func getDateFromDaynr(daynr uint) (year uint, month uint, day uint) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// Handle leapyears leapday.

const (
	intervalYEAR        = "YEAR"
	intervalQUARTER     = "QUARTER"
	intervalMONTH       = "MONTH"
	intervalWEEK        = "WEEK"
	intervalDAY         = "DAY"
	intervalHOUR        = "HOUR"
	intervalMINUTE      = "MINUTE"
	intervalSECOND      = "SECOND"
	intervalMICROSECOND = "MICROSECOND"
)

func timestampDiff(intervalType string, t1 TimeInternal, t2 TimeInternal) int64 {
	_ = "STUB: not implemented"
	return 0
}

// calc years

// calc months

// In MySQL difference between any two valid datetime values
// in microseconds fits into longlong.
