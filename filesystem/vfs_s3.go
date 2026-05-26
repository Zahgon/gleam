package filesystem

import (
	"io"
	"os"
)

const (
	AWS_ACCESS_KEY = OptionName("aws_access_key")
	AWS_SECRET_KEY = OptionName("aws_secret_key")
)

type S3FileSystem struct {
}

func (fs *S3FileSystem) Accept(fl *FileLocation) bool { _ = "STUB: not implemented"; return false }

func (fs *S3FileSystem) Open(fl *FileLocation) (VirtualFile, error) {
	_ = "STUB: not implemented"
	return *new(VirtualFile), nil
}

// Required
// Required

func (fs *S3FileSystem) List(fl *FileLocation) (fileLocations []*FileLocation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *S3FileSystem) IsDir(fl *FileLocation) bool { _ = "STUB: not implemented"; return false }

func splitS3LocationToParts(location string) (bucketName, objectKey string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type VirtualFileS3 struct {
	*os.File
	filename string
	size     int64
}

func newVirtualFileS3(readerCloser io.ReadCloser) (*VirtualFileS3, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vf *VirtualFileS3) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (vf *VirtualFileS3) Close() error { _ = "STUB: not implemented"; return nil }
