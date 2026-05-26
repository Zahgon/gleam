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

package segmentmap

import (
	"hash/crc32"
)

// SegmentMap is used for handle a big map slice by slice.
// It's not thread safe.
type SegmentMap struct {
	size int64
	maps []map[string]interface{}

	crcTable *crc32.Table
}

// NewSegmentMap creates a new SegmentMap.
func NewSegmentMap(size int64) (*SegmentMap, error) { _ = "STUB: not implemented"; return nil, nil }

// Get is the same as map[k].
func (sm *SegmentMap) Get(key []byte) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetSegment gets the map specific by index.
func (sm *SegmentMap) GetSegment(index int64) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set if key not exists, returns whether already exists.
func (sm *SegmentMap) Set(key []byte, value interface{}, force bool) bool {
	_ = "STUB: not implemented"
	return false
}

// SegmentCount returns how many inner segments.
func (sm *SegmentMap) SegmentCount() int64 { _ = "STUB: not implemented"; return 0 }
