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

// EncodeDecimal encodes a decimal d into a byte slice which can be sorted lexicographically later.
func EncodeDecimal(b []byte, d types.Datum) []byte { _ = "STUB: not implemented"; return nil }

// DecodeDecimal decodes bytes to decimal.
func DecodeDecimal(b []byte) ([]byte, types.Datum, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Datum), nil
}
