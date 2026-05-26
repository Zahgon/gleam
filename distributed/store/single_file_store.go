package store

import (
	"os"
	"sync"
)

type SingleFileStore struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.
	Filename string

	size           int64
	file           *os.File
	mu             sync.Mutex
	waitForReading *sync.Cond

	Offset   int64 // offset at the head of the file
	Position int64 // offset for current tail, write position
}

func (l *SingleFileStore) init() {
	l.waitForReading = sync.NewCond(&l.mu)
}

func (l *SingleFileStore) ReadAt(data []byte, offset int64) (int, error) {
	_ = "STUB: not implemented"

	// fmt.Printf("Read: l.Offset=%d, offset=%d\n", l.Offset, offset)
	return 0, nil
}

// create the file does not exist

// fmt.Printf("Read: creating new file...\n")

// read written data

// wait for data not written yet

// fmt.Printf("Read: wait for reading...\n")

// fmt.Printf("Read: file reading...\n")

// Write implements io.Writer.  If a write would cause the log file to be larger
// than MaxMegaByte, the file is closed, renamed to include a timestamp of the
// current time, and a new log file is created using the original log file name.
// If the length of the write is greater than MaxMegaByte, an error is returned.
func (l *SingleFileStore) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements io.Closer, and closes the current logfile.
func (l *SingleFileStore) Close() error { _ = "STUB: not implemented"; return nil }

// close closes the file if it is open.
func (l *SingleFileStore) close() error { _ = "STUB: not implemented"; return nil }

// openNew opens a new log file for writing, moving any old log file out of the
// way.  This methods assumes the file has already been closed.
func (l *SingleFileStore) openNew() error { _ = "STUB: not implemented"; return nil }

// we use truncate here because this should only get called when we've moved
// the file ourselves. if someone else creates the file in the meantime,
// just wipe out the contents.

func (l *SingleFileStore) filename() string { _ = "STUB: not implemented"; return "" }

func (l *SingleFileStore) Destroy() {
	_ = "STUB: not implemented"
	// println("removing file", l.filename())
	return
}

// dir returns the directory for the current filename.
func (l *SingleFileStore) dir() string { _ = "STUB: not implemented"; return "" }
