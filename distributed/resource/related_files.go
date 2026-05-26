package resource

type FileResource struct {
	FullPath     string `json:"path,omitempty"`
	TargetFolder string `json:"targetFolder,omitempty"`
}

type FileHash struct {
	FullPath     string `json:"path,omitempty"`
	TargetFolder string `json:"targetFolder,omitempty"`
	File         string `json:"file,omitempty"`
	Hash         uint32 `json:"hash,omitempty"`
}

func GenerateFileHash(fullpath string) (*FileHash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
