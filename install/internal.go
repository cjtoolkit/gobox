package install

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cjtoolkit/gobox/model"
	"github.com/cjtoolkit/gobox/tool"
)

func execInternal(locals []model.TomlLocal, goCmd, binPath, projectPath string) error {
	curWd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Chdir(projectPath)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Chdir(curWd)
	for _, local := range locals {
		path := binPath + filepath.FromSlash("/"+strings.Trim(local.BinPath, "/"))
		err := tool.RunCommand("", path, goCmd, "install", local.Install)
		if err != nil {
			return err
		}
	}
	return nil
}
