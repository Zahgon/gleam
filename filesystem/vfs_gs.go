package filesystem

import (
	"io"
	"os"
)

// Contains logic for Google Storage virtual filesystem
// Note that Google Storage authentication expects GOOGLE_APPLICATION_CREDENTIALS
// see: https://cloud.google.com/docs/authentication/getting-started
// Note that this is handled internally for clusters running within Google Cloud

// GoogleStorageFileSystem type
type GoogleStorageFileSystem struct{}

// Accept - criteria for accepting path as google storage location
func (fs *GoogleStorageFileSystem) Accept(fl *FileLocation) bool {
	_ = "STUB: not implemented"
	return false
}

// Open - open given file location and return virtual file
func (fs *GoogleStorageFileSystem) Open(fl *FileLocation) (v VirtualFile, err error) {
	_ = "STUB: not implemented"
	return *new(VirtualFile), nil
}

// List - list items in google storage directory
func (fs *GoogleStorageFileSystem) List(fl *FileLocation) (fileLocations []*FileLocation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsDir - returns true if directory detected
// This function assumes that if the prefix ends with "/"
// then the intent of the user is to represent a Dir
func (fs *GoogleStorageFileSystem) IsDir(fl *FileLocation) bool {
	_ = "STUB: not implemented"
	return false
}

// VirtualFileGS - Virtual File implementation for Google Storage
type VirtualFileGS struct {
	*os.File
	filename string
	size     int64
}

// Size - returns virtual file  size
func (vf *VirtualFileGS) Size() int64 {
	_ = "STUB: not implemented"

	// Close - virtual file Closeer function
	return 0
}

func (vf *VirtualFileGS) Close() error { _ = "STUB: not implemented"; return nil }

func newVirtualFileGS(readerCloser io.ReadCloser) (*VirtualFileGS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
