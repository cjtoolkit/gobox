package tool

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCommand(path, gobin, name string, args ...string) error {
	if path != "" {
		name = path + filepath.FromSlash("/"+name)
	}
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()
	if path != "" {
		cmd.Env = append(cmd.Env, "PATH="+path+fmt.Sprintf("%c", os.PathListSeparator)+os.Getenv("PATH"))
	}
	if gobin != "" {
		cmd.Env = append(cmd.Env, "GOBIN="+gobin)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
