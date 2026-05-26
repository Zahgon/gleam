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
	"github.com/juju/errors"
)

const signMask uint64 = 0x8000000000000000

func encodeIntToCmpUint(v int64) uint64 { _ = "STUB: not implemented"; return 0 }

func decodeCmpUintToInt(u uint64) int64 { _ = "STUB: not implemented"; return 0 }

// EncodeInt appends the encoded value to slice b and returns the appended slice.
// EncodeInt guarantees that the encoded value is in ascending order for comparison.
func EncodeInt(b []byte, v int64) []byte { _ = "STUB: not implemented"; return nil }

// EncodeIntDesc appends the encoded value to slice b and returns the appended slice.
// EncodeIntDesc guarantees that the encoded value is in descending order for comparison.
func EncodeIntDesc(b []byte, v int64) []byte { _ = "STUB: not implemented"; return nil }

// DecodeInt decodes value encoded by EncodeInt before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeInt(b []byte) ([]byte, int64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// DecodeIntDesc decodes value encoded by EncodeInt before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeIntDesc(b []byte) ([]byte, int64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// EncodeUint appends the encoded value to slice b and returns the appended slice.
// EncodeUint guarantees that the encoded value is in ascending order for comparison.
func EncodeUint(b []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }

// EncodeUintDesc appends the encoded value to slice b and returns the appended slice.
// EncodeUintDesc guarantees that the encoded value is in descending order for comparison.
func EncodeUintDesc(b []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }

// DecodeUint decodes value encoded by EncodeUint before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeUint(b []byte) ([]byte, uint64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// DecodeUintDesc decodes value encoded by EncodeInt before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeUintDesc(b []byte) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// EncodeVarint appends the encoded value to slice b and returns the appended slice.
// Note that the encoded result is not memcomparable.
func EncodeVarint(b []byte, v int64) []byte { _ = "STUB: not implemented"; return nil }

// DecodeVarint decodes value encoded by EncodeVarint before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeVarint(b []byte) ([]byte, int64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// EncodeUvarint appends the encoded value to slice b and returns the appended slice.
// Note that the encoded result is not memcomparable.
func EncodeUvarint(b []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }

// DecodeUvarint decodes value encoded by EncodeUvarint before.
// It returns the leftover un-decoded slice, decoded value if no error.
func DecodeUvarint(b []byte) ([]byte, uint64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

const (
	negativeTagEnd   = 8        // Negative tag is (negativeTagEnd - length).
	positiveTagStart = 0xff - 8 // Positive tag is (positiveTagStart + length).
)

// EncodeComparableVarint encodes an int64 to a mem-comparable bytes.
func EncodeComparableVarint(b []byte, v int64) []byte {
	_ = "STUB: not implemented"

	// All negative value has a tag byte prefix (negativeTagEnd - length).
	// Smaller negative value encodes to more bytes, has smaller tag.
	return nil
}

// EncodeComparableUvarint encodes uint64 into mem-comparable bytes.
func EncodeComparableUvarint(b []byte, v uint64) []byte {
	_ = "STUB: not implemented"
	// The first byte has 256 values, [0, 7] is reserved for negative tags,
	// [248, 255] is reserved for larger positive tags,
	// So we can store value [0, 239] in a single byte.
	// Values cannot be stored in single byte has a tag byte prefix (positiveTagStart+length).
	// Larger value encodes to more bytes, has larger tag.
	return nil
}

var (
	errDecodeInsufficient = errors.New("insufficient bytes to decode value")
	errDecodeInvalid      = errors.New("invalid bytes to decode value")
)

// DecodeComparableUvarint decodes mem-comparable uvarint.
func DecodeComparableUvarint(b []byte) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// DecodeComparableVarint decodes mem-comparable varint.
func DecodeComparableVarint(b []byte) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Negative value has all bits on by default.
