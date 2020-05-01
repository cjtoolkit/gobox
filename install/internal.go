package install

import (
	"path/filepath"
	"strings"

	"github.com/cjtoolkit/gobox/model"
	"github.com/cjtoolkit/gobox/tool"
)

func execInternal(locals []model.TomlLocal, goCmd, binPath string) error {
	for _, local := range locals {
		path := binPath + filepath.FromSlash("/"+strings.Trim(local.BinPath, "/"))
		err := tool.RunCommand("", path, goCmd, "install", local.Install)
		if err != nil {
			return err
		}
	}
	return nil
}
