// Disk-backed queue
package store

import (
	"io"
	"time"
)

type DataStore interface {
	io.Writer
	io.ReaderAt
	Destroy()
	LastWriteAt() time.Time
	LastReadAt() time.Time
}

type LocalFileDataStore struct {
	dir         string
	name        string
	store       *SingleFileStore
	lastWriteAt time.Time
	lastReadAt  time.Time
}

func NewLocalFileDataStore(dir, name string) (ds *LocalFileDataStore) {
	_ = "STUB: not implemented"
	return nil
}

func (ds *LocalFileDataStore) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ds *LocalFileDataStore) ReadAt(data []byte, offset int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ds *LocalFileDataStore) Destroy() { _ = "STUB: not implemented"; return }

func (ds *LocalFileDataStore) LastWriteAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (ds *LocalFileDataStore) LastReadAt() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
