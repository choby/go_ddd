package files

import "github.com/choby/go_ddd/interfaces"

type FileInfos struct {
	interfaces.FileInfo
	FileBody []byte
}
