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
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
)

// Kind constants.
const (
	KindNull          byte = 0
	KindInt64         byte = 1
	KindUint64        byte = 2
	KindFloat32       byte = 3
	KindFloat64       byte = 4
	KindString        byte = 5
	KindBytes         byte = 6
	KindMysqlBit      byte = 7
	KindMysqlDecimal  byte = 8
	KindMysqlDuration byte = 9
	KindMysqlEnum     byte = 10
	KindMysqlHex      byte = 11
	KindMysqlSet      byte = 12
	KindMysqlTime     byte = 13
	KindRow           byte = 14
	KindInterface     byte = 15
	KindMinNotNull    byte = 16
	KindMaxValue      byte = 17
)

// Datum is a data box holds different kind of data.
// It has better performance and is easier to use than `interface{}`.
type Datum struct {
	k         byte        // datum kind.
	collation uint8       // collation can hold uint8 values.
	decimal   uint16      // decimal can hold uint16 values.
	length    uint32      // length can hold uint32 values.
	i         int64       // i can hold int64 uint64 float64 values.
	b         []byte      // b can hold string or []byte values.
	x         interface{} // f hold all other types.
}

// Kind gets the kind of the datum.
func (d *Datum) Kind() byte {
	_ = "STUB: not implemented"

	// Collation gets the collation of the datum.
	return 0
}

func (d *Datum) Collation() byte {
	_ = "STUB: not implemented"

	// SetCollation sets the collation of the datum.
	return 0
}

func (d *Datum) SetCollation(collation byte) { _ = "STUB: not implemented"; return }

// Frac gets the frac of the datum.
func (d *Datum) Frac() int { _ = "STUB: not implemented"; return 0 }

// SetFrac sets the frac of the datum.
func (d *Datum) SetFrac(frac int) { _ = "STUB: not implemented"; return }

// Length gets the length of the datum.
func (d *Datum) Length() int { _ = "STUB: not implemented"; return 0 }

// SetLength sets the length of the datum
func (d *Datum) SetLength(l int) { _ = "STUB: not implemented"; return }

// IsNull checks if datum is null.
func (d *Datum) IsNull() bool { _ = "STUB: not implemented"; return false }

// GetInt64 gets int64 value.
func (d *Datum) GetInt64() int64 {
	_ = "STUB: not implemented"

	// SetInt64 sets int64 value.
	return 0
}

func (d *Datum) SetInt64(i int64) { _ = "STUB: not implemented"; return }

// GetUint64 gets uint64 value.
func (d *Datum) GetUint64() uint64 {
	_ = "STUB: not implemented"

	// SetUint64 sets uint64 value.
	return 0
}

func (d *Datum) SetUint64(i uint64) { _ = "STUB: not implemented"; return }

// GetFloat64 gets float64 value.
func (d *Datum) GetFloat64() float64 { _ = "STUB: not implemented"; return 0 }

// SetFloat64 sets float64 value.
func (d *Datum) SetFloat64(f float64) { _ = "STUB: not implemented"; return }

// GetFloat32 gets float32 value.
func (d *Datum) GetFloat32() float32 { _ = "STUB: not implemented"; return 0 }

// SetFloat32 sets float32 value.
func (d *Datum) SetFloat32(f float32) { _ = "STUB: not implemented"; return }

// GetString gets string value.
func (d *Datum) GetString() string { _ = "STUB: not implemented"; return "" }

// SetString sets string value.
func (d *Datum) SetString(s string) { _ = "STUB: not implemented"; return }

// sink prevents s from being allocated on the stack.
var sink = func(s string) {
}

// GetBytes gets bytes value.
func (d *Datum) GetBytes() []byte {
	_ = "STUB: not implemented"

	// SetBytes sets bytes value to datum.
	return nil
}

func (d *Datum) SetBytes(b []byte) { _ = "STUB: not implemented"; return }

// SetBytesAsString sets bytes value to datum as string type.
func (d *Datum) SetBytesAsString(b []byte) { _ = "STUB: not implemented"; return }

// GetInterface gets interface value.
func (d *Datum) GetInterface() interface{} {
	_ = "STUB: not implemented"

	// SetInterface sets interface to datum.
	return nil
}

func (d *Datum) SetInterface(x interface{}) { _ = "STUB: not implemented"; return }

// GetRow gets row value.
func (d *Datum) GetRow() []Datum { _ = "STUB: not implemented"; return nil }

// SetRow sets row value.
func (d *Datum) SetRow(ds []Datum) { _ = "STUB: not implemented"; return }

// SetNull sets datum to nil.
func (d *Datum) SetNull() { _ = "STUB: not implemented"; return }

// GetMysqlBit gets Bit value
func (d *Datum) GetMysqlBit() Bit { _ = "STUB: not implemented"; return *new(Bit) }

// SetMysqlBit sets Bit value
func (d *Datum) SetMysqlBit(b Bit) { _ = "STUB: not implemented"; return }

// GetMysqlDecimal gets Decimal value
func (d *Datum) GetMysqlDecimal() *MyDecimal { _ = "STUB: not implemented"; return nil }

// SetMysqlDecimal sets Decimal value
func (d *Datum) SetMysqlDecimal(b *MyDecimal) { _ = "STUB: not implemented"; return }

// GetMysqlDuration gets Duration value
func (d *Datum) GetMysqlDuration() Duration { _ = "STUB: not implemented"; return *new(Duration) }

// SetMysqlDuration sets Duration value
func (d *Datum) SetMysqlDuration(b Duration) { _ = "STUB: not implemented"; return }

// GetMysqlEnum gets Enum value
func (d *Datum) GetMysqlEnum() Enum { _ = "STUB: not implemented"; return *new(Enum) }

// SetMysqlEnum sets Enum value
func (d *Datum) SetMysqlEnum(b Enum) { _ = "STUB: not implemented"; return }

// GetMysqlHex gets Hex value
func (d *Datum) GetMysqlHex() Hex {
	_ = "STUB: not implemented"
	return *

	// SetMysqlHex sets Hex value
	new(Hex)
}

func (d *Datum) SetMysqlHex(b Hex) { _ = "STUB: not implemented"; return }

// GetMysqlSet gets Set value
func (d *Datum) GetMysqlSet() Set { _ = "STUB: not implemented"; return *new(Set) }

// SetMysqlSet sets Set value
func (d *Datum) SetMysqlSet(b Set) { _ = "STUB: not implemented"; return }

// GetMysqlTime gets types.Time value
func (d *Datum) GetMysqlTime() Time {
	_ = "STUB: not implemented"

	// SetMysqlTime sets types.Time value
	return *new(Time)
}

func (d *Datum) SetMysqlTime(b Time) { _ = "STUB: not implemented"; return }

// GetValue gets the value of the datum of any kind.
func (d *Datum) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

// SetValue sets any kind of value.
func (d *Datum) SetValue(val interface{}) { _ = "STUB: not implemented"; return }

// CompareDatum compares datum to another datum.
// TODO: return error properly.
func (d *Datum) CompareDatum(sc *variable.StatementContext, ad Datum) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareInt64(sc *variable.StatementContext, i int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareUint64(sc *variable.StatementContext, u uint64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareFloat64(sc *variable.StatementContext, f float64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareString(sc *variable.StatementContext, s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareBytes(sc *variable.StatementContext, b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlBit(sc *variable.StatementContext, bit Bit) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlDecimal(sc *variable.StatementContext, dec *MyDecimal) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlDuration(sc *variable.StatementContext, dur Duration) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlEnum(sc *variable.StatementContext, enum Enum) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlHex(sc *variable.StatementContext, e Hex) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlSet(sc *variable.StatementContext, set Set) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareMysqlTime(sc *variable.StatementContext, time Time) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) compareRow(sc *variable.StatementContext, row []Datum) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Cast casts datum to certain types.
func (d *Datum) Cast(sc *variable.StatementContext, target *FieldType) (ad Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// ConvertTo converts a datum to the target field type.
func (d *Datum) ConvertTo(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// TODO: implement mysql types convert when "CAST() AS" syntax are supported.

func (d *Datum) convertToFloat(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// For float and following double type, we will only truncate it for float(M, D) format.
// If no D is set, we will handle it like origin float whether M is set or not.

func (d *Datum) convertToString(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// Flen is the rune length, not binary length, for UTF8 charset, we need to calculate the
// rune count and truncate to Flen runes if it is too long.

// We do break here because we need to iterate to the end to get runeCount.

func (d *Datum) convertToInt(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToUint(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToMysqlTime(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToMysqlDuration(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToMysqlDecimal(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToMysqlYear(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToMysqlBit(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// check bit boundary, if bit has n width, the boundary is
// in [0, (1 << n) - 1]

func (d *Datum) convertToMysqlEnum(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convertToMysqlSet(sc *variable.StatementContext, target *FieldType) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// ToBool converts to a bool.
// We will use 1 for true, and 0 for false.
func (d *Datum) ToBool(sc *variable.StatementContext) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ConvertDatumToDecimal converts datum to decimal.
func ConvertDatumToDecimal(sc *variable.StatementContext, d Datum) (*MyDecimal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToDecimal converts to a decimal.
func (d *Datum) ToDecimal(sc *variable.StatementContext) (*MyDecimal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToInt64 converts to a int64.
func (d *Datum) ToInt64(sc *variable.StatementContext) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Datum) toSignedInteger(sc *variable.StatementContext, tp byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// 2011-11-10 11:11:11.999999 -> 20111110111112

// 11:11:11.999999 -> 111112

// ToFloat64 converts to a float64
func (d *Datum) ToFloat64(sc *variable.StatementContext) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToString gets the string representation of the datum.
func (d *Datum) ToString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func invalidConv(d *Datum, tp byte) (Datum, error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

func (d *Datum) convergeType(hasUint, hasDecimal, hasFloat *bool) (x Datum) {
	_ = "STUB: not implemented"
	return *new(Datum)
}

// CoerceDatum changes type.
// If a or b is Float, changes the both to Float.
// Else if a or b is Decimal, changes the both to Decimal.
// Else if a or b is Uint and op is not div, mod, or intDiv changes the both to Uint.
func CoerceDatum(sc *variable.StatementContext, a, b Datum) (x, y Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), *new(Datum), nil
}

// NewDatum creates a new Datum from an interface{}.
func NewDatum(in interface{}) (d Datum) { _ = "STUB: not implemented"; return *new(Datum) }

// NewIntDatum creates a new Datum from an int64 value.
func NewIntDatum(i int64) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewUintDatum creates a new Datum from an uint64 value.
	new(Datum)
}

func NewUintDatum(i uint64) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewBytesDatum creates a new Datum from a byte slice.
	new(Datum)
}

func NewBytesDatum(b []byte) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewStringDatum creates a new Datum from a string.
	new(Datum)
}

func NewStringDatum(s string) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewFloat64Datum creates a new Datum from a float64 value.
	new(Datum)
}

func NewFloat64Datum(f float64) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewFloat32Datum creates a new Datum from a float32 value.
	new(Datum)
}

func NewFloat32Datum(f float32) (d Datum) {
	_ = "STUB: not implemented"
	return *

	// NewDurationDatum creates a new Datum from a Duration value.
	new(Datum)
}

func NewDurationDatum(dur Duration) (d Datum) { _ = "STUB: not implemented"; return *new(Datum) }

// NewDecimalDatum creates a new Datum form a MyDecimal value.
func NewDecimalDatum(dec *MyDecimal) (d Datum) { _ = "STUB: not implemented"; return *new(Datum) }

// MakeDatums creates datum slice from interfaces.
func MakeDatums(args ...interface{}) []Datum { _ = "STUB: not implemented"; return nil }

// DatumsToInterfaces converts a datum slice to interface slice.
func DatumsToInterfaces(datums []Datum) []interface{} { _ = "STUB: not implemented"; return nil }

// MinNotNullDatum returns a datum represents minimum not null value.
func MinNotNullDatum() Datum { _ = "STUB: not implemented"; return *new(Datum) }

// MaxValueDatum returns a datum represents max value.
func MaxValueDatum() Datum { _ = "STUB: not implemented"; return *new(Datum) }

// EqualDatums compare if a and b contains the same datum values.
func EqualDatums(sc *variable.StatementContext, a []Datum, b []Datum) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SortDatums sorts a slice of datum.
func SortDatums(sc *variable.StatementContext, datums []Datum) error {
	_ = "STUB: not implemented"
	return nil
}

type datumsSorter struct {
	datums []Datum
	sc     *variable.StatementContext
	err    error
}

func (ds *datumsSorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (ds *datumsSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ds *datumsSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func handleTruncateError(sc *variable.StatementContext) error {
	_ = "STUB: not implemented"
	return nil
}
