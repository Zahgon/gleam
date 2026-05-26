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

func encodeFloatToCmpUint64(f float64) uint64 { _ = "STUB: not implemented"; return 0 }

func decodeCmpUintToFloat(u uint64) float64 { _ = "STUB: not implemented"; return 0 }

// EncodeFloat encodes a float v into a byte slice which can be sorted lexicographically later.
// EncodeFloat guarantees that the encoded value is in ascending order for comparison.
func EncodeFloat(b []byte, v float64) []byte { _ = "STUB: not implemented"; return nil }

// DecodeFloat decodes a float from a byte slice generated with EncodeFloat before.
func DecodeFloat(b []byte) ([]byte, float64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

// EncodeFloatDesc encodes a float v into a byte slice which can be sorted lexicographically later.
// EncodeFloatDesc guarantees that the encoded value is in descending order for comparison.
func EncodeFloatDesc(b []byte, v float64) []byte { _ = "STUB: not implemented"; return nil }

// DecodeFloatDesc decodes a float from a byte slice generated with EncodeFloatDesc before.
func DecodeFloatDesc(b []byte) ([]byte, float64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
