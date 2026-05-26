package zipfile

import (
	"archive/zip"

	"github.com/chrislusf/gleam/util"
)

type FileReader struct {
	reader   *zip.ReadCloser
	Cursor   int
	NumFiles int
}

func New(filename string) *FileReader { _ = "STUB: not implemented"; return nil }

func (r *FileReader) ReadHeader() (fieldNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read will iterate through the zip file, it will treat each file
// in the zipfile as a row and return back to the caller, where the
// key is file or directory name and the value is the bytes of the
// file from the input zip file
func (r *FileReader) Read() (row *util.Row, err error) { _ = "STUB: not implemented"; return nil, nil }
