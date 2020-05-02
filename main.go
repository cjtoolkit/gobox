package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cjtoolkit/gobox/install"
	"github.com/cjtoolkit/gobox/model"
	"github.com/cjtoolkit/gobox/tool"

	"github.com/pelletier/go-toml"
)

func main() {
	goboxPath := tool.SeekGoboxFile()

	cachePath, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(goboxPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	projectPath := filepath.Dir(goboxPath)

	goboxData := model.TomlSupplement{}

	err = toml.NewDecoder(file).Decode(&goboxData)
	if err != nil {
		log.Fatal(err)
	}

	cachePath += filepath.FromSlash("/gobox/" + goboxData.Hash())
	binPath := cachePath + filepath.FromSlash("/bin")

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		install.Exec(goboxData, cachePath, binPath, projectPath)
	}

	if len(os.Args) < 2 {
		return
	}

	cmdName := os.Args[1]
	if strings.Index(cmdName, "/") >= 1 || strings.Index(cmdName, "\\") >= 1 {
		binPath += filepath.FromSlash("/" + filepath.Dir(cmdName))
		cmdName = filepath.Base(cmdName)
	}

	err = tool.RunCommand(binPath, "", cmdName, os.Args[2:]...)
	if err != nil {
		log.Fatal(err)
	}
}
