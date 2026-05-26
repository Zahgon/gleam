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

// RoundFloat rounds float val to the nearest integer value with float64 format, like MySQL Round function.
// RoundFloat uses default rounding mode, see https://dev.mysql.com/doc/refman/5.7/en/precision-math-rounding.html
// so rounding use "round half away from zero".
// e.g, 1.5 -> 2, -1.5 -> -2.
func RoundFloat(f float64) float64 { _ = "STUB: not implemented"; return 0 }

// Round rounds the argument f to dec decimal places.
// dec defaults to 0 if not specified. dec can be negative
// to cause dec digits left of the decimal point of the
// value f to become zero.
func Round(f float64, dec int) float64 { _ = "STUB: not implemented"; return 0 }

func getMaxFloat(flen int, decimal int) float64 { _ = "STUB: not implemented"; return 0 }

// TruncateFloat tries to truncate f.
// If the result exceeds the max/min float that flen/decimal allowed, returns the max/min float allowed.
func TruncateFloat(f float64, flen int, decimal int) (float64, error) {
	_ = "STUB: not implemented"

	// nan returns 0
	return 0, nil
}

func isSpace(c byte) bool { _ = "STUB: not implemented"; return false }

func isDigit(c byte) bool { _ = "STUB: not implemented"; return false }

func myMax(a, b int) int { _ = "STUB: not implemented"; return 0 }

func myMaxInt8(a, b int8) int8 { _ = "STUB: not implemented"; return 0 }

func myMin(a, b int) int { _ = "STUB: not implemented"; return 0 }

func myMinInt8(a, b int8) int8 { _ = "STUB: not implemented"; return 0 }

// strToInt converts a string to an integer in best effort.
// TODO: handle overflow and add unittest.
func strToInt(str string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO: if i < len(str), we should return an error.
