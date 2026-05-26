package filesystem

import (
	"os"
)

type LocalFileSystem struct {
}

func (fs *LocalFileSystem) Accept(fl *FileLocation) bool { _ = "STUB: not implemented"; return false }

func (fs *LocalFileSystem) Open(fl *FileLocation) (VirtualFile, error) {
	_ = "STUB: not implemented"
	return *new(VirtualFile), nil
}

func (fs *LocalFileSystem) List(fl *FileLocation) (fileLocations []*FileLocation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *LocalFileSystem) IsDir(fl *FileLocation) bool { _ = "STUB: not implemented"; return false }

type VirtualFileLocal struct {
	*os.File
}

func (vf *VirtualFileLocal) Size() int64 { _ = "STUB: not implemented"; return 0 }
