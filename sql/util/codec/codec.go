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

package codec

import (
	"github.com/chrislusf/gleam/sql/util/types"
)

// First byte in the encoded value which specifies the encoding type.
const (
	NilFlag          byte = 0
	bytesFlag        byte = 1
	compactBytesFlag byte = 2
	intFlag          byte = 3
	uintFlag         byte = 4
	floatFlag        byte = 5
	decimalFlag      byte = 6
	durationFlag     byte = 7
	varintFlag       byte = 8
	uvarintFlag      byte = 9
	maxFlag          byte = 250
)

func encode(b []byte, vals []types.Datum, comparable bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// duration may have negative value, so we cannot use String to encode directly.

func encodeBytes(b []byte, v []byte, comparable bool) []byte { _ = "STUB: not implemented"; return nil }

func encodeSignedInt(b []byte, v int64, comaprable bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func encodeUnsignedInt(b []byte, v uint64, comparable bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// EncodeKey appends the encoded values to byte slice b, returns the appended
// slice. It guarantees the encoded value is in ascending order for comparison.
// For Decimal type, datum must set datum's length and frac.
func EncodeKey(b []byte, v ...types.Datum) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// EncodeValue appends the encoded values to byte slice b, returning the appended
		// slice. It does not guarantee the order for comparison.
		nil
}

func EncodeValue(b []byte, v ...types.Datum) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Decode decodes values from a byte slice generated with EncodeKey or EncodeValue
		// before.
		// size is the size of decoded datum slice.
		nil
}

func Decode(b []byte, size int) ([]types.Datum, error) { _ = "STUB: not implemented"; return nil, nil }

// DecodeOne decodes on datum from a byte slice generated with EncodeKey or EncodeValue.
func DecodeOne(b []byte) (remain []byte, d types.Datum, err error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Datum), nil
}

// use max fsp, let outer to do round manually.

// CutOne cuts the first encoded value from b.
// It will return the first encoded item and the remains as byte slice.
func CutOne(b []byte) (data []byte, remain []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// peeks the first encoded value from b and returns its length.
func peek(b []byte) (length int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Those types are stored in 8 bytes.

func peekBytes(b []byte, reverse bool) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// The byte slice is encoded into many groups.
// For each group, there are 8 bytes for data and 1 byte for marker.

// When padCount is not zero, it means we get the end of the byte slice.

func peekCompactBytes(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// Get length.
	return 0, nil
}

func peekVarint(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func peekUvarint(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
