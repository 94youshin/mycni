package hanlder

import (
	"github.com/usmhk/mycni/pkg/args"
	"github.com/usmhk/mycni/pkg/version"
)

type FileHandler struct {
	*version.VersionInfo
	IpStore string
}

func NewFileHandler(filename string) Handler {
	return &FileHandler{
		VersionInfo: &version.VersionInfo{},
		IpStore:     filename,
	}
}

func (fh *FileHandler) HandleAdd(cmdArgs *args.CmdArgs) error {
	return nil
}

func (fh *FileHandler) HandleDel(cmdArgs *args.CmdArgs) error {
	return nil
}

func (fh *FileHandler) HandleCheck(cmdArgs *args.CmdArgs) error {
	return nil
}

func (fh *FileHandler) HandleVersion(cmdArgs *args.CmdArgs) error {
	return nil
}
