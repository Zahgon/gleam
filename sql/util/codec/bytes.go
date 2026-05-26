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
	"runtime"
	"unsafe"
)

const (
	encGroupSize = 8
	encMarker    = byte(0xFF)
	encPad       = byte(0x0)
)

var (
	pads    = make([]byte, encGroupSize)
	encPads = []byte{encPad}
)

// EncodeBytes guarantees the encoded value is in ascending order for comparison,
// encoding with the following rule:
//
//	[group1][marker1]...[groupN][markerN]
//	group is 8 bytes slice which is padding with 0.
//	marker is `0xFF - padding 0 count`
//
// For example:
//
//	[] -> [0, 0, 0, 0, 0, 0, 0, 0, 247]
//	[1, 2, 3] -> [1, 2, 3, 0, 0, 0, 0, 0, 250]
//	[1, 2, 3, 0] -> [1, 2, 3, 0, 0, 0, 0, 0, 251]
//	[1, 2, 3, 4, 5, 6, 7, 8] -> [1, 2, 3, 4, 5, 6, 7, 8, 255, 0, 0, 0, 0, 0, 0, 0, 0, 247]
//
// Refer: https://github.com/facebook/mysql-5.6/wiki/MyRocks-record-format#memcomparable-format
func EncodeBytes(b []byte, data []byte) []byte {
	_ = "STUB: not implemented"
	// Allocate more space to avoid unnecessary slice growing.
	// Assume that the byte slice size is about `(len(data) / encGroupSize + 1) * (encGroupSize + 1)` bytes,
	// that is `(len(data) / 8 + 1) * 9` in our implement.
	return nil
}

func decodeBytes(b []byte, reverse bool) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Check validity of padding bytes.

// DecodeBytes decodes bytes which is encoded by EncodeBytes before,
// returns the leftover bytes and decoded value if no error.
func DecodeBytes(b []byte) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// EncodeBytesDesc first encodes bytes using EncodeBytes, then bitwise reverses
		// encoded value to guarantee the encoded value is in descending order for comparison.
		nil, nil
}

func EncodeBytesDesc(b []byte, data []byte) []byte { _ = "STUB: not implemented"; return nil }

// DecodeBytesDesc decodes bytes which is encoded by EncodeBytesDesc before,
// returns the leftover bytes and decoded value if no error.
func DecodeBytesDesc(b []byte) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// EncodeCompactBytes joins bytes with its length into a byte slice. It is more
		// efficient in both space and time compare to EncodeBytes. Note that the encoded
		// result is not memcomparable.
		nil, nil
}

func EncodeCompactBytes(b []byte, data []byte) []byte { _ = "STUB: not implemented"; return nil }

// DecodeCompactBytes decodes bytes which is encoded by EncodeCompactBytes before.
func DecodeCompactBytes(b []byte) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// See https://golang.org/src/crypto/cipher/xor.go
const wordSize = int(unsafe.Sizeof(uintptr(0)))
const supportsUnaligned = runtime.GOARCH == "386" || runtime.GOARCH == "amd64"

func fastReverseBytes(b []byte) { _ = "STUB: not implemented"; return }

func safeReverseBytes(b []byte) { _ = "STUB: not implemented"; return }

func reverseBytes(b []byte) { _ = "STUB: not implemented"; return }

// like realloc.
func reallocBytes(b []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

// slice b has capability to store n bytes
