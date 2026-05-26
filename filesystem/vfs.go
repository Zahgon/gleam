package filesystem

// this file defines the virtual file system to provide consistent file access APIs

import (
	"io"
)

type OptionName string

var (
	Option = make(map[OptionName]string)
)

type FileLocation struct {
	Location string
}

type VirtualFile interface {
	io.ReaderAt
	io.ReadCloser
	io.Seeker
	Size() int64
}

type VirtualFileSystem interface {
	Accept(*FileLocation) bool
	Open(*FileLocation) (VirtualFile, error)
	List(*FileLocation) ([]*FileLocation, error)
	IsDir(*FileLocation) bool
}

var (
	fileSystems = []VirtualFileSystem{
		&LocalFileSystem{},
		&HdfsFileSystem{},
		&S3FileSystem{},
		&GoogleStorageFileSystem{},
	}
)

func Set(name OptionName, value string) { _ = "STUB: not implemented"; return }

func Open(filepath string) (VirtualFile, error) {
	_ = "STUB: not implemented"
	return *new(VirtualFile), nil
}

func List(filepath string) ([]*FileLocation, error) { _ = "STUB: not implemented"; return nil, nil }

func IsDir(filepath string) bool { _ = "STUB: not implemented"; return false }
