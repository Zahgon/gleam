package file

import (
	"github.com/chrislusf/gleam/flow"
)

type FileSource struct {
	folder         string
	fileBaseName   string
	hasWildcard    bool
	Path           string
	HasHeader      bool
	PartitionCount int
	FileType       string
	Fields         []string

	prefix string
}

// Generate generates data shard info,
// partitions them via round robin,
// and reads each shard on each executor
func (s *FileSource) Generate(f *flow.Flow) *flow.Dataset { _ = "STUB: not implemented"; return nil }

// SetHasHeader sets whether the data contains header
func (q *FileSource) SetHasHeader(hasHeader bool) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

// TODO adjust FileSource api to denote which data source can support columnar reads
// Select selects fields that can be pushed down to data sources supporting columnar reads
func (q *FileSource) Select(fields ...string) *FileSource { _ = "STUB: not implemented"; return nil }

// New creates a FileSource based on a file name.
// The base file name can have "*", "?" pattern denoting a list of file names.
func newFileSource(fileType, fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("file source: %+v\n", s)

func (s *FileSource) genShardInfos(f *flow.Flow) *flow.Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (s *FileSource) match(fullPath string) bool { _ = "STUB: not implemented"; return false }
