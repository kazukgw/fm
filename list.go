package main

import (
	"encoding/json"
	"fmt"
	"github.com/kazukgw/coa"
	"os"
	"os/exec"
	"path/filepath"
)

type showList struct {
	GetCurrentDir
	GetFmPath
	GetListKey
	coa.DoSelf
	DefaultErrorHandler
}

func (ag *showList) Do(c coa.Context) error {
	var fpath string
	if ag.GetListKey.Key == "" {
		cmd := exec.Command("find", ag.GetCurrentDir.Dir, "-name", ".git", "-prune", "-o", "-print")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return err
		}
		return nil
	}

	fpath = filepath.Join(ag.FmPath, ag.GetListKey.Key+".json")

	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	var dat map[string]interface{}
	if err := json.NewDecoder(f).Decode(&dat); err != nil {
		return err
	}

	for cpath, val := range dat {
		fmt.Printf("%s\t%s\n", cpath, val)
	}

	return nil
}
