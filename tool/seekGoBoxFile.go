package tool

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func SeekGoboxFile() string {
	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gnode := filepath.FromSlash("/gobox.toml")
	lastPath := false
	for {
		if _, err := os.Stat(curdir + gnode); os.IsNotExist(err) {
			curdir = filepath.Dir(curdir)
			if lastPath || curdir == "." {
				log.Fatal("Could not find 'gobox.toml'")
			} else if strings.Trim(curdir, "/") == "" || (runtime.GOOS == "windows" && len(strings.Trim(curdir, "\\")) == 2) {
				lastPath = true
			}
			continue
		}
		break
	}
	return curdir + gnode
}
