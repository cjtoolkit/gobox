package model

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"
)

type TomlSupplement struct {
	GoCmd   string       `toml:"goCmd"`
	Locals  []TomlLocal  `toml:"local"`
	Modules []TomlModule `toml:"module"`
}

func (t TomlSupplement) Cmd() string {
	if t.GoCmd == "" {
		return "go"
	}
	return strings.Split(getGoVersion(t.GoCmd), " ")[2]
}

func (t TomlSupplement) Hash() string {
	var i []string
	for _, local := range t.Locals {
		i = append(i, local.hashData())
	}
	sort.Strings(i)
	hash := sha256.New()
	fmt.Fprint(hash, getGoVersion(t.Cmd()), i)
	var v []string
	for _, module := range t.Modules {
		v = append(v, module.hashData())
	}
	sort.Strings(v)
	fmt.Fprint(hash, v)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

type TomlLocal struct {
	BinPath string `toml:"binPath"`
	Install string `toml:"install"`
}

func (t TomlLocal) hashData() string {
	return fmt.Sprint(t.BinPath, t.Install)
}

type TomlModule struct {
	Repo     string   `toml:"repo"`
	Tag      string   `toml:"tag"`
	BinPath  string   `toml:"binPath"`
	Installs []string `toml:"installs"`
}

func (t TomlModule) RepoAndTag() string {
	if t.Tag == "" {
		return t.Repo
	}

	return t.Repo + "@" + t.Tag
}

func (t TomlModule) hashData() string {
	i := append([]string{}, t.Installs...)
	sort.Strings(i)
	return fmt.Sprint(t.Repo, t.Tag, t.BinPath, i)
}

func getGoVersion(goCmd string) string {
	b := &bytes.Buffer{}
	cmd := exec.Command(goCmd, "version")
	cmd.Stdout = b
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return b.String()
}
