package install

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cjtoolkit/gobox/model"
	"github.com/cjtoolkit/gobox/tool"
)

func execExternal(modules []model.TomlModule, goCmd, cachdPath, binPath string) error {
	err := os.Chdir(cachdPath)
	if err != nil {
		return err
	}
	for _, module := range modules {
		err := tool.RunCommand("", "", goCmd, "get", "-d", module.RepoAndTag())
		if err != nil {
			return err
		}
		moduleBinPath := binPath
		if module.BinPath != "" {
			moduleBinPath += filepath.FromSlash("/" + strings.Trim(module.BinPath, "/"))
		}
		for _, install := range module.Installs {
			err := installExternal(module, install, goCmd, cachdPath, moduleBinPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func installExternal(module model.TomlModule, install, goCmd, cachePath, binPath string) error {
	output := binPath
	if install == "" || install == "." {
		output += filepath.FromSlash("/" + filepath.Base(module.Repo))
		install = module.Repo
	} else {
		output += filepath.FromSlash("/" + filepath.Base(install))
		install = module.Repo + filepath.FromSlash("/"+strings.Trim(install, "/"))
	}
	if runtime.GOOS == "windows" {
		output += ".exe"
	}
	tool.RunCommand("", "", goCmd, "build", "-o", output, "-i", install)
	return nil
}
