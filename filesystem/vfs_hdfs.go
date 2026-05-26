package filesystem

import (
	"github.com/colinmarc/hdfs"
)

/*
Get the namenode from
1) the hdfs://namenode/... string
2) Or from env HADOOP_NAMENODE.
3) Or based on env HADOOP_CONF_DIR or HADOOP_HOME
to locate hdfs-site.xml and core-site.xml
*/
type HdfsFileSystem struct {
}

func (fs *HdfsFileSystem) Accept(fl *FileLocation) bool { _ = "STUB: not implemented"; return false }

func (fs *HdfsFileSystem) Open(fl *FileLocation) (VirtualFile, error) {
	_ = "STUB: not implemented"
	return *new(VirtualFile), nil
}

// List generates a full list of file locations under the given
// location, which should have a prefix of hdfs://
func (fs *HdfsFileSystem) List(fl *FileLocation) (fileLocations []*FileLocation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *HdfsFileSystem) IsDir(fl *FileLocation) bool { _ = "STUB: not implemented"; return false }

func splitLocationToParts(location string) (namenode, path string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type VirtualFileHdfs struct {
	*hdfs.FileReader
}

func (vf *VirtualFileHdfs) Size() int64 { _ = "STUB: not implemented"; return 0 }
