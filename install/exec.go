package install

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cjtoolkit/gobox/model"
)

func Exec(goboxData model.TomlSupplement, cachePath, binPath, projectPath string) {
	curWd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	defer os.Chdir(curWd)
	err = os.Chdir(projectPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(binPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	goCmd := goboxData.Cmd()

	err = createFile(cachePath+filepath.FromSlash("/"+"main.go"), model.MainGo)
	if err != nil {
		cleanPath(cachePath)
		log.Fatal(err)
	}
	err = createFile(cachePath+filepath.FromSlash("/"+"go.mod"), model.GoMod)
	if err != nil {
		cleanPath(cachePath)
		log.Fatal(err)
	}

	if goboxData.Locals != nil {
		err := execInternal(goboxData.Locals, goCmd, binPath)
		if err != nil {
			cleanPath(cachePath)
			log.Fatal(err)
		}
	}
	if goboxData.Modules != nil {
		err := execExternal(goboxData.Modules, goCmd, cachePath, binPath)
		if err != nil {
			cleanPath(cachePath)
			log.Fatal(err)
		}
	}
}

func cleanPath(cachDir string) {
	os.RemoveAll(cachDir)
}

func createFile(name, content string) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	return err
}
