// Copyright 2013 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

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

package expression

import (
	"regexp"
	"time"

	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
	"github.com/chrislusf/gleam/sql/util/types"
)

var (
	_ functionClass = &dateFunctionClass{}
	_ functionClass = &dateDiffFunctionClass{}
	_ functionClass = &timeDiffFunctionClass{}
	_ functionClass = &dateFormatFunctionClass{}
	_ functionClass = &dayFunctionClass{}
	_ functionClass = &hourFunctionClass{}
	_ functionClass = &minuteFunctionClass{}
	_ functionClass = &secondFunctionClass{}
	_ functionClass = &microSecondFunctionClass{}
	_ functionClass = &monthFunctionClass{}
	_ functionClass = &monthNameFunctionClass{}
	_ functionClass = &nowFunctionClass{}
	_ functionClass = &dayNameFunctionClass{}
	_ functionClass = &dayOfMonthFunctionClass{}
	_ functionClass = &dayOfWeekFunctionClass{}
	_ functionClass = &dayOfYearFunctionClass{}
	_ functionClass = &weekFunctionClass{}
	_ functionClass = &weekDayFunctionClass{}
	_ functionClass = &weekOfYearFunctionClass{}
	_ functionClass = &yearFunctionClass{}
	_ functionClass = &yearWeekFunctionClass{}
	_ functionClass = &fromUnixTimeFunctionClass{}
	_ functionClass = &strToDateFunctionClass{}
	_ functionClass = &sysDateFunctionClass{}
	_ functionClass = &currentDateFunctionClass{}
	_ functionClass = &currentTimeFunctionClass{}
	_ functionClass = &timeFunctionClass{}
	_ functionClass = &utcDateFunctionClass{}
	_ functionClass = &extractFunctionClass{}
	_ functionClass = &arithmeticFunctionClass{}
	_ functionClass = &unixTimestampFunctionClass{}
)

var (
	_ builtinFunc = &builtinDateSig{}
	_ builtinFunc = &builtinDateDiffSig{}
	_ builtinFunc = &builtinTimeDiffSig{}
	_ builtinFunc = &builtinDateFormatSig{}
	_ builtinFunc = &builtinDaySig{}
	_ builtinFunc = &builtinHourSig{}
	_ builtinFunc = &builtinMinuteSig{}
	_ builtinFunc = &builtinSecondSig{}
	_ builtinFunc = &builtinMicroSecondSig{}
	_ builtinFunc = &builtinMonthSig{}
	_ builtinFunc = &builtinMonthNameSig{}
	_ builtinFunc = &builtinNowSig{}
	_ builtinFunc = &builtinDayNameSig{}
	_ builtinFunc = &builtinDayOfMonthSig{}
	_ builtinFunc = &builtinDayOfWeekSig{}
	_ builtinFunc = &builtinDayOfYearSig{}
	_ builtinFunc = &builtinWeekSig{}
	_ builtinFunc = &builtinWeekDaySig{}
	_ builtinFunc = &builtinWeekOfYearSig{}
	_ builtinFunc = &builtinYearSig{}
	_ builtinFunc = &builtinYearWeekSig{}
	_ builtinFunc = &builtinFromUnixTimeSig{}
	_ builtinFunc = &builtinStrToDateSig{}
	_ builtinFunc = &builtinSysDateSig{}
	_ builtinFunc = &builtinCurrentDateSig{}
	_ builtinFunc = &builtinCurrentTimeSig{}
	_ builtinFunc = &builtinTimeSig{}
	_ builtinFunc = &builtinUTCDateSig{}
	_ builtinFunc = &builtinExtractSig{}
	_ builtinFunc = &builtinArithmeticSig{}
	_ builtinFunc = &builtinUnixTimestampSig{}
)

func convertToTime(sc *variable.StatementContext, arg types.Datum, tp byte) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func convertToDuration(sc *variable.StatementContext, arg types.Datum, fsp int) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type dateFunctionClass struct {
	baseFunctionClass
}

func (c *dateFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDateSig struct {
	baseBuiltinFunc
}

func (b *builtinDateSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_date
func builtinDate(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func convertDatumToTime(sc *variable.StatementContext, d types.Datum) (t types.Time, err error) {
	_ = "STUB: not implemented"
	return *new(types.Time), nil
}

type dateDiffFunctionClass struct {
	baseFunctionClass
}

func (c *dateDiffFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDateDiffSig struct {
	baseBuiltinFunc
}

func (b *builtinDateDiffSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_datediff
func builtinDateDiff(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type timeDiffFunctionClass struct {
	baseFunctionClass
}

func (c *timeDiffFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinTimeDiffSig struct {
	baseBuiltinFunc
}

func (b *builtinTimeDiffSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_timediff
func builtinTimeDiff(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type dateFormatFunctionClass struct {
	baseFunctionClass
}

func (c *dateFormatFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDateFormatSig struct {
	baseBuiltinFunc
}

func (b *builtinDateFormatSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_date-format
func builtinDateFormat(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type fromDaysFunctionClass struct {
	baseFunctionClass
}

func (c *fromDaysFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinFromDaysSig struct {
	baseBuiltinFunc
}

func (b *builtinFromDaysSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_from-days
func builtinFromDays(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type dayFunctionClass struct {
	baseFunctionClass
}

func (c *dayFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDaySig struct {
	baseBuiltinFunc
}

func (b *builtinDaySig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_day
// Day is a synonym for DayOfMonth.
func builtinDay(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type hourFunctionClass struct {
	baseFunctionClass
}

func (c *hourFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinHourSig struct {
	baseBuiltinFunc
}

func (b *builtinHourSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_hour
func builtinHour(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

type minuteFunctionClass struct {
	baseFunctionClass
}

func (c *minuteFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinMinuteSig struct {
	baseBuiltinFunc
}

func (b *builtinMinuteSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_minute
func builtinMinute(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

type secondFunctionClass struct {
	baseFunctionClass
}

func (c *secondFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSecondSig struct {
	baseBuiltinFunc
}

func (b *builtinSecondSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_second
func builtinSecond(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

type microSecondFunctionClass struct {
	baseFunctionClass
}

func (c *microSecondFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinMicroSecondSig struct {
	baseBuiltinFunc
}

func (b *builtinMicroSecondSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_microsecond
func builtinMicroSecond(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

type monthFunctionClass struct {
	baseFunctionClass
}

func (c *monthFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinMonthSig struct {
	baseBuiltinFunc
}

func (b *builtinMonthSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_month
func builtinMonth(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

type monthNameFunctionClass struct {
	baseFunctionClass
}

func (c *monthNameFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinMonthNameSig struct {
	baseBuiltinFunc
}

func (b *builtinMonthNameSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_monthname
func builtinMonthName(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type nowFunctionClass struct {
	baseFunctionClass
}

func (c *nowFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinNowSig struct {
	baseBuiltinFunc
}

func (b *builtinNowSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_now
func builtinNow(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// TODO: if NOW is used in stored function or trigger, NOW will return the beginning time
	// of the execution.
	return *new(types.Datum), nil
}

// set unspecified for later round

type dayNameFunctionClass struct {
	baseFunctionClass
}

func (c *dayNameFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDayNameSig struct {
	baseBuiltinFunc
}

func (b *builtinDayNameSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_dayname
func builtinDayName(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type dayOfMonthFunctionClass struct {
	baseFunctionClass
}

func (c *dayOfMonthFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDayOfMonthSig struct {
	baseBuiltinFunc
}

func (b *builtinDayOfMonthSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_dayofmonth
func builtinDayOfMonth(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// TODO: some invalid format like 2000-00-00 will return 0 too.
	return *new(types.Datum), nil
}

// No need to check type here.

type dayOfWeekFunctionClass struct {
	baseFunctionClass
}

func (c *dayOfWeekFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDayOfWeekSig struct {
	baseBuiltinFunc
}

func (b *builtinDayOfWeekSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_dayofweek
func builtinDayOfWeek(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

// TODO: log warning or return error?

// 1 is Sunday, 2 is Monday, .... 7 is Saturday

type dayOfYearFunctionClass struct {
	baseFunctionClass
}

func (c *dayOfYearFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDayOfYearSig struct {
	baseBuiltinFunc
}

func (b *builtinDayOfYearSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_dayofyear
func builtinDayOfYear(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// TODO: log warning or return error?

type weekFunctionClass struct {
	baseFunctionClass
}

func (c *weekFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinWeekSig struct {
	baseBuiltinFunc
}

func (b *builtinWeekSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_week
func builtinWeek(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

// TODO: log warning or return error?

type weekDayFunctionClass struct {
	baseFunctionClass
}

func (c *weekDayFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinWeekDaySig struct {
	baseBuiltinFunc
}

func (b *builtinWeekDaySig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_weekday
func builtinWeekDay(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

// TODO: log warning or return error?

// Monday is 0, ... Sunday = 6 in MySQL
// but in go, Sunday is 0, ... Saturday is 6
// w will do a conversion.

type weekOfYearFunctionClass struct {
	baseFunctionClass
}

func (c *weekOfYearFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinWeekOfYearSig struct {
	baseBuiltinFunc
}

func (b *builtinWeekOfYearSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_weekofyear
func builtinWeekOfYear(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	// WeekOfYear is equivalent to to Week(date, 3)
	return *new(types.Datum), nil
}

type yearFunctionClass struct {
	baseFunctionClass
}

func (c *yearFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinYearSig struct {
	baseBuiltinFunc
}

func (b *builtinYearSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_year
func builtinYear(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

type yearWeekFunctionClass struct {
	baseFunctionClass
}

func (c *yearWeekFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinYearWeekSig struct {
	baseBuiltinFunc
}

func (b *builtinYearWeekSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_yearweek
func builtinYearWeek(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// No need to check type here.

// TODO: log warning or return error?

type fromUnixTimeFunctionClass struct {
	baseFunctionClass
}

func (c *fromUnixTimeFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinFromUnixTimeSig struct {
	baseBuiltinFunc
}

func (b *builtinFromUnixTimeSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_from-unixtime
func builtinFromUnixTime(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// 0 <= unixTimeStamp <= INT32_MAX

// Split the integral part and fractional part of a decimal timestamp.
// e.g. for timestamp 12345.678,
// first get the integral part 12345,
// then (12345.678 - 12345) * (10^9) to get the decimal part and convert it to nanosecond precision.

// here fractionalPart is result multiplying the original fractional part by 10^9.

// Keep consistent with MySQL.

type strToDateFunctionClass struct {
	baseFunctionClass
}

func (c *strToDateFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinStrToDateSig struct {
	baseBuiltinFunc
}

func (b *builtinStrToDateSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.5/en/date-and-time-functions.html#function_str-to-date
func builtinStrToDate(args []types.Datum, _ context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type sysDateFunctionClass struct {
	baseFunctionClass
}

func (c *sysDateFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSysDateSig struct {
	baseBuiltinFunc
}

func (b *builtinSysDateSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_sysdate
func builtinSysDate(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	// SYSDATE is not the same as NOW if NOW is used in a stored function or trigger.
	// But here we can just think they are the same because we don't support stored function
	// and trigger now.
	return *new(types.Datum), nil
}

type currentDateFunctionClass struct {
	baseFunctionClass
}

func (c *currentDateFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCurrentDateSig struct {
	baseBuiltinFunc
}

func (b *builtinCurrentDateSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_curdate
func builtinCurrentDate(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type currentTimeFunctionClass struct {
	baseFunctionClass
}

func (c *currentTimeFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCurrentTimeSig struct {
	baseBuiltinFunc
}

func (b *builtinCurrentTimeSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_curtime
func builtinCurrentTime(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type timeFunctionClass struct {
	baseFunctionClass
}

func (c *timeFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinTimeSig struct {
	baseBuiltinFunc
}

func (b *builtinTimeSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_time
func builtinTime(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type utcDateFunctionClass struct {
	baseFunctionClass
}

func (c *utcDateFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinUTCDateSig struct {
	baseBuiltinFunc
}

func (b *builtinUTCDateSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_utc-date
func builtinUTCDate(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type extractFunctionClass struct {
	baseFunctionClass
}

func (c *extractFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinExtractSig struct {
	baseBuiltinFunc
}

func (b *builtinExtractSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_extract
func builtinExtract(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func checkFsp(sc *variable.StatementContext, arg types.Datum) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type dateArithFunctionClass struct {
	baseFunctionClass

	op ast.DateArithType
}

func (c *dateArithFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinDateArithSig struct {
	baseBuiltinFunc

	op ast.DateArithType
}

func (b *builtinDateArithSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func dateArithFuncFactory(op ast.DateArithType) BuiltinFunc {
	_ = "STUB: not implemented"
	return *new(BuiltinFunc)
}

// args[0] -> Date
// args[1] -> Interval Value
// args[2] -> Interval Unit
// health check for date and interval

// parse date

// parse interval

// TODO: Consider time_zone variable.

var reg = regexp.MustCompile(`[\d]+`)

func parseDayInterval(sc *variable.StatementContext, value types.Datum) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type timestampDiffFunctionClass struct {
	baseFunctionClass
}

func (c *timestampDiffFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinTimestampDiffSig struct {
	baseBuiltinFunc
}

func (b *builtinTimestampDiffSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_timestampdiff
func builtinTimestampDiff(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// errorOrWarning reports error or warning depend on the context.
func errorOrWarning(err error, ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO: Use better name, such as sc.IsInsert instead of sc.IgnoreTruncate.

type unixTimestampFunctionClass struct {
	baseFunctionClass
}

func (c *unixTimestampFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinUnixTimestampSig struct {
	baseBuiltinFunc
}

func (b *builtinUnixTimestampSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_unix-timestamp
func builtinUnixTimestamp(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func getTimeZone(ctx context.Context) *time.Location { _ = "STUB: not implemented"; return nil }
